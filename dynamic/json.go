package dynamic

// JSON marshalling and unmarshalling for dynamic messages

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"reflect"
	"strconv"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/jhump/protoreflect/desc"
)

func (m *Message) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	if err := m.marshalJSON(&b, false); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (m *Message) MarshalJSONIndent() ([]byte, error) {
	var b bytes.Buffer
	if err := m.marshalJSON(&b, true); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (m *Message) marshalJSON(b *bytes.Buffer, indent bool) error {
	// TODO
	return nil
}

func (m *Message) UnmarshalJSON(js []byte) error {
	m.Reset()
	if err := m.UnmarshalMergeJSON(js); err != nil {
		return err
	}
	return m.Validate()
}

func (m *Message) UnmarshalMergeJSON(js []byte) error {
	r := &jsReader{ dec: json.NewDecoder(bytes.NewReader(js)) }
	r.dec.UseNumber()
	err := m.unmarshalJson(r)
	if err != nil {
		return err
	}
	if t, err := r.poll(); err != io.EOF {
		b, _ := ioutil.ReadAll(r.dec.Buffered())
		s := fmt.Sprintf("%v%s", t, string(b))
		return fmt.Errorf("Superfluous data found after JSON object: %q", s)
	}
	return nil
}

func (m *Message) unmarshalJson(r *jsReader) error {
	t, err := r.peek()
	if err != nil {
		return err
	}
	if t == nil {
		// if json is simply "null" we do nothing
		r.poll()
		return nil
	}

	if err := r.beginObject(); err != nil {
		return err
	}

	for r.hasNext() {
		f, err := r.nextObjectKey()
		if err != nil {
			return err
		}
		fd := m.FindFieldDescriptorByName(f)
		if fd == nil {
			r.skip()
			continue
		}
		v, err := unmarshalJsField(fd, r, m.er)
		if err != nil {
			return err
		}
		if v != nil {
			m.internalSetField(fd, v)
		} else if m.values != nil {
			delete(m.values, fd.GetNumber())
		}
	}

	if err := r.endObject(); err != nil {
		return err
	}

	return nil
}

func unmarshalJsField(fd *desc.FieldDescriptor, r *jsReader, er *ExtensionRegistry) (interface{}, error) {
	t, err := r.peek()
	if err != nil {
		return nil, err
	}
	if t == nil {
		// if value is null, just return nil
		r.poll()
		return nil, nil
	}

	if t == json.Delim('{') && fd.IsMap() {
		entryType := fd.GetMessageType()
		keyType := entryType.FindFieldByNumber(1)
		valueType := entryType.FindFieldByNumber(2)
		mp := map[interface{}]interface{}{}

		// TODO: if there are just two map keys "key" and "value" and they have the right type of values,
		// treat this JSON object as a single map entry message. (In keeping with support of map fields as
		// if they were normal repeated field of entry messages as well as supporting a transition from
		// optional to repeated...)

		if err := r.beginObject(); err != nil {
			return nil, err
		}
		for r.hasNext() {
			kk, err := unmarshalJsFieldElement(keyType, r, er)
			if err != nil {
				return nil, err
			}
			vv, err := unmarshalJsFieldElement(valueType, r, er)
			if err != nil {
				return nil, err
			}
			mp[kk] = vv
		}
		if err := r.endObject(); err != nil {
			return nil, err
		}

		return mp, nil
	} else if t == json.Delim('[') {
		// We support parsing an array, even if field is not repeated, to mimic support in proto
		// binary wire format that supports changing an optional field to repeated and vice versa.
		// If the field is not repeated, we only keep the last value in the array.

		if err := r.beginArray(); err != nil {
			return nil, err
		}
		var sl []interface{}
		var v interface{}
		for r.hasNext() {
			var err error
			v, err = unmarshalJsFieldElement(fd, r, er)
			if err != nil {
				return nil, err
			}
			if fd.IsRepeated() && v != nil {
				sl = append(sl, v)
			}
		}
		if err := r.endArray(); err != nil {
			return nil, err
		}
		if fd.IsMap() {
			mp := map[interface{}]interface{}{}
			for _, m := range sl {
				msg := m.(*Message)
				kk, err := msg.TryGetFieldByNumber(1)
				if err != nil {
					return nil, err
				}
				vv, err := msg.TryGetFieldByNumber(2)
				if err != nil {
					return nil, err
				}
				mp[kk] = vv
			}
			return mp, nil
		} else if fd.IsRepeated() {
			return sl, nil
		} else {
			return v, nil
		}
	} else {
		// We support parsing a singular value, even if field is repeated, to mimic support in proto
		// binary wire format that supports changing an optional field to repeated and vice versa.
		// If the field is repeated, we store value as singleton slice of that one value.

		v, err := unmarshalJsFieldElement(fd, r, er)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, nil
		}
		if fd.IsRepeated() {
			return []interface{} { v }, nil
		} else {
			return v, nil
		}
	}
}

func unmarshalJsFieldElement(fd *desc.FieldDescriptor, r *jsReader, er *ExtensionRegistry) (interface{}, error) {
	t, err := r.peek()
	if err != nil {
		return nil, err
	}
	if t == nil {
		// if value is null, just return nil
		r.poll()
		return nil, nil
	}

	switch fd.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE,
		descriptor.FieldDescriptorProto_TYPE_GROUP:
		m := NewMessageWithExtensionRegistry(fd.GetMessageType(), er)
		if err := m.unmarshalJson(r); err != nil {
			return nil, err
		} else {
			return m, nil
		}

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		if e, err := r.nextNumber(); err != nil {
			return nil, err
		} else {
			// value could be string or number
			if i, err := e.Int64(); err != nil {
				// number cannot be parsed, so see if it's an enum value name
				fqn := fmt.Sprintf("%s.%s", fd.GetEnumType(), string(e))
				if ev, ok := fd.GetFile().FindSymbol(fqn).(*desc.EnumValueDescriptor); ok {
					return ev.GetNumber(), nil
				} else {
					// could not find it!
					return nil, err
				}
			} else if i > math.MaxInt32 || i < math.MinInt32 {
				return nil, NumericOverflowError
			} else {
				return int32(i), err
			}
		}

	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		if i, err := r.nextInt(); err != nil {
			return nil, err
		} else if i > math.MaxInt32 || i < math.MinInt32 {
			return nil, NumericOverflowError
		} else {
			return int32(i), err
		}

	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return r.nextInt()

	case descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32:
		if i, err := r.nextUint(); err != nil {
			return nil, err
		} else if i > math.MaxUint32 {
			return nil, NumericOverflowError
		} else {
			return uint32(i), err
		}

	case descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return r.nextUint()

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return r.nextBool()

	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if f, err := r.nextFloat(); err != nil {
			return nil, err
		} else {
			return float32(f), nil
		}

	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return r.nextFloat()

	default:
		return nil, fmt.Errorf("Unknown field type: %v", fd.GetType())
	}
}

type jsReader struct {
	dec     *json.Decoder
	current json.Token
	peeked  bool
}

func (r *jsReader) hasNext() bool {
	return r.dec.More()
}

func (r *jsReader) peek() (json.Token, error) {
	if r.peeked {
		return r.current, nil
	}
	t, err := r.dec.Token()
	if err != nil {
		return nil, err
	}
	r.peeked = true
	r.current = t
	return t, nil
}

func (r *jsReader) poll() (json.Token, error) {
	if r.peeked {
		ret := r.current
		r.current = nil
		r.peeked = false
		return ret, nil
	}
	return r.dec.Token()
}

func (r *jsReader) beginObject() error {
	_, err := r.expect(func(t json.Token) bool { return t == json.Delim('{') }, nil, "start of JSON object: '{'")
	return err
}

func (r *jsReader) endObject() error {
	_, err := r.expect(func(t json.Token) bool { return t == json.Delim('}') }, nil, "end of JSON object: '}'")
	return err
}

func (r *jsReader) beginArray() error {
	_, err := r.expect(func(t json.Token) bool { return t == json.Delim('[') }, nil, "start of array: '['")
	return err
}

func (r *jsReader) endArray() error {
	_, err := r.expect(func(t json.Token) bool { return t == json.Delim(']') }, nil, "end of array: ']'")
	return err
}

func (r *jsReader) nextObjectKey() (string, error) {
	return r.nextString()
}

func (r *jsReader) nextString() (string, error) {
	t, err := r.expect(func(t json.Token) bool { _, ok := t.(string); return ok }, "", "string")
	if err != nil {
		return "", err
	}
	return t.(string), nil
}

func (r *jsReader) nextBytes() ([]byte, error) {
	str, err := r.nextString()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(str)
}

func (r *jsReader) nextBool() (bool, error) {
	t, err := r.expect(func(t json.Token) bool { _, ok := t.(bool); return ok }, false, "boolean")
	if err != nil {
		return false, err
	}
	return t.(bool), nil
}

func (r *jsReader) nextInt() (int64, error) {
	n, err := r.nextNumber()
	if err != nil {
		return 0, err
	}
	return n.Int64()
}

func (r *jsReader) nextUint() (uint64, error) {
	n, err := r.nextNumber()
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(string(n), 10, 64)
}

func (r *jsReader) nextFloat() (float64, error) {
	n, err := r.nextNumber()
	if err != nil {
		return 0, err
	}
	return n.Float64()
}

func (r *jsReader) nextNumber() (json.Number, error) {
	t, err := r.expect(func(t json.Token) bool { return reflect.TypeOf(t).Kind() == reflect.String }, "0", "number")
	if err != nil {
		return "", err
	}
	switch t := t.(type) {
	case json.Number:
		return t, nil
	case string:
		return json.Number(t), nil
	}
	return "", fmt.Errorf("Expecting a number but got %v", t)
}

func (r *jsReader) skip() error {
	t, err := r.poll()
	if err != nil {
		return err
	}
	if t == json.Delim('[') {
		if err := r.skipArray(); err != nil {
			return err
		}
	} else if t == json.Delim('{') {
		if err := r.skipObject(); err != nil {
			return err
		}
	}
	return nil
}

func (r *jsReader) skipArray() error {
	for r.hasNext() {
		if err := r.skip(); err != nil {
			return err
		}
	}
	if err := r.endArray(); err != nil {
		return err
	}
	return nil
}

func (r *jsReader) skipObject() error {
	for r.hasNext() {
		// skip object key
		if err := r.skip(); err != nil {
			return err
		}
		// and value
		if err := r.skip(); err != nil {
			return err
		}
	}
	if err := r.endObject(); err != nil {
		return err
	}
	return nil
}

func (r *jsReader) expect(predicate func(json.Token) bool, ifNil interface{}, expected string) (interface{}, error) {
	t, err := r.poll()
	if err != nil {
		return nil, err
	}
	if t == nil && ifNil != nil {
		return ifNil, nil
	}
	if !predicate(t) {
		return t, fmt.Errorf("Bad input. Expecting %s. Instead got: %v.", expected, t)
	}
	return t, nil
}
