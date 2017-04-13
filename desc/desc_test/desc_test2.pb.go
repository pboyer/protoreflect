// Code generated by protoc-gen-go.
// source: desc_test2.proto
// DO NOT EDIT!

package desc_test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import jhump_protoreflect_desc "github.com/pboyer/protoreflect/desc/desc_test/pkg"
import desc_test_nopkg "github.com/pboyer/protoreflect/desc/desc_test/nopkg"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Frobnitz struct {
	A *TestMessage        `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B *AnotherTestMessage `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	// Types that are valid to be assigned to Abc:
	//	*Frobnitz_C1
	//	*Frobnitz_C2
	Abc isFrobnitz_Abc             `protobuf_oneof:"abc"`
	D   *TestMessage_NestedMessage `protobuf:"bytes,5,opt,name=d" json:"d,omitempty"`
	E   *TestMessage_NestedEnum    `protobuf:"varint,6,opt,name=e,enum=desc_test.TestMessage_NestedEnum,def=2" json:"e,omitempty"`
	F   []string                   `protobuf:"bytes,7,rep,name=f" json:"f,omitempty"`
	// Types that are valid to be assigned to Def:
	//	*Frobnitz_G1
	//	*Frobnitz_G2
	//	*Frobnitz_G3
	Def              isFrobnitz_Def `protobuf_oneof:"def"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Frobnitz) Reset()                    { *m = Frobnitz{} }
func (m *Frobnitz) String() string            { return proto.CompactTextString(m) }
func (*Frobnitz) ProtoMessage()               {}
func (*Frobnitz) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

const Default_Frobnitz_E TestMessage_NestedEnum = TestMessage_VALUE2

type isFrobnitz_Abc interface {
	isFrobnitz_Abc()
}
type isFrobnitz_Def interface {
	isFrobnitz_Def()
}

type Frobnitz_C1 struct {
	C1 *TestMessage_NestedMessage `protobuf:"bytes,3,opt,name=c1,oneof"`
}
type Frobnitz_C2 struct {
	C2 TestMessage_NestedEnum `protobuf:"varint,4,opt,name=c2,enum=desc_test.TestMessage_NestedEnum,oneof"`
}
type Frobnitz_G1 struct {
	G1 int32 `protobuf:"varint,8,opt,name=g1,oneof"`
}
type Frobnitz_G2 struct {
	G2 int32 `protobuf:"zigzag32,9,opt,name=g2,oneof"`
}
type Frobnitz_G3 struct {
	G3 uint32 `protobuf:"varint,10,opt,name=g3,oneof"`
}

func (*Frobnitz_C1) isFrobnitz_Abc() {}
func (*Frobnitz_C2) isFrobnitz_Abc() {}
func (*Frobnitz_G1) isFrobnitz_Def() {}
func (*Frobnitz_G2) isFrobnitz_Def() {}
func (*Frobnitz_G3) isFrobnitz_Def() {}

func (m *Frobnitz) GetAbc() isFrobnitz_Abc {
	if m != nil {
		return m.Abc
	}
	return nil
}
func (m *Frobnitz) GetDef() isFrobnitz_Def {
	if m != nil {
		return m.Def
	}
	return nil
}

func (m *Frobnitz) GetA() *TestMessage {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Frobnitz) GetB() *AnotherTestMessage {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *Frobnitz) GetC1() *TestMessage_NestedMessage {
	if x, ok := m.GetAbc().(*Frobnitz_C1); ok {
		return x.C1
	}
	return nil
}

func (m *Frobnitz) GetC2() TestMessage_NestedEnum {
	if x, ok := m.GetAbc().(*Frobnitz_C2); ok {
		return x.C2
	}
	return TestMessage_VALUE1
}

func (m *Frobnitz) GetD() *TestMessage_NestedMessage {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *Frobnitz) GetE() TestMessage_NestedEnum {
	if m != nil && m.E != nil {
		return *m.E
	}
	return Default_Frobnitz_E
}

func (m *Frobnitz) GetF() []string {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *Frobnitz) GetG1() int32 {
	if x, ok := m.GetDef().(*Frobnitz_G1); ok {
		return x.G1
	}
	return 0
}

func (m *Frobnitz) GetG2() int32 {
	if x, ok := m.GetDef().(*Frobnitz_G2); ok {
		return x.G2
	}
	return 0
}

func (m *Frobnitz) GetG3() uint32 {
	if x, ok := m.GetDef().(*Frobnitz_G3); ok {
		return x.G3
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Frobnitz) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Frobnitz_OneofMarshaler, _Frobnitz_OneofUnmarshaler, _Frobnitz_OneofSizer, []interface{}{
		(*Frobnitz_C1)(nil),
		(*Frobnitz_C2)(nil),
		(*Frobnitz_G1)(nil),
		(*Frobnitz_G2)(nil),
		(*Frobnitz_G3)(nil),
	}
}

func _Frobnitz_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Frobnitz)
	// abc
	switch x := m.Abc.(type) {
	case *Frobnitz_C1:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.C1); err != nil {
			return err
		}
	case *Frobnitz_C2:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.C2))
	case nil:
	default:
		return fmt.Errorf("Frobnitz.Abc has unexpected type %T", x)
	}
	// def
	switch x := m.Def.(type) {
	case *Frobnitz_G1:
		b.EncodeVarint(8<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.G1))
	case *Frobnitz_G2:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeZigzag32(uint64(x.G2))
	case *Frobnitz_G3:
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.G3))
	case nil:
	default:
		return fmt.Errorf("Frobnitz.Def has unexpected type %T", x)
	}
	return nil
}

func _Frobnitz_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Frobnitz)
	switch tag {
	case 3: // abc.c1
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TestMessage_NestedMessage)
		err := b.DecodeMessage(msg)
		m.Abc = &Frobnitz_C1{msg}
		return true, err
	case 4: // abc.c2
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Abc = &Frobnitz_C2{TestMessage_NestedEnum(x)}
		return true, err
	case 8: // def.g1
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Def = &Frobnitz_G1{int32(x)}
		return true, err
	case 9: // def.g2
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag32()
		m.Def = &Frobnitz_G2{int32(x)}
		return true, err
	case 10: // def.g3
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Def = &Frobnitz_G3{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Frobnitz_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Frobnitz)
	// abc
	switch x := m.Abc.(type) {
	case *Frobnitz_C1:
		s := proto.Size(x.C1)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Frobnitz_C2:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.C2))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// def
	switch x := m.Def.(type) {
	case *Frobnitz_G1:
		n += proto.SizeVarint(8<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.G1))
	case *Frobnitz_G2:
		n += proto.SizeVarint(9<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64((uint32(x.G2) << 1) ^ uint32((int32(x.G2) >> 31))))
	case *Frobnitz_G3:
		n += proto.SizeVarint(10<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.G3))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Whatchamacallit struct {
	Foos             *jhump_protoreflect_desc.Foo `protobuf:"varint,1,req,name=foos,enum=jhump.protoreflect.desc.Foo" json:"foos,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *Whatchamacallit) Reset()                    { *m = Whatchamacallit{} }
func (m *Whatchamacallit) String() string            { return proto.CompactTextString(m) }
func (*Whatchamacallit) ProtoMessage()               {}
func (*Whatchamacallit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Whatchamacallit) GetFoos() jhump_protoreflect_desc.Foo {
	if m != nil && m.Foos != nil {
		return *m.Foos
	}
	return jhump_protoreflect_desc.Foo_ABC
}

type Whatzit struct {
	Gyzmeau          []*jhump_protoreflect_desc.Bar `protobuf:"bytes,1,rep,name=gyzmeau" json:"gyzmeau,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *Whatzit) Reset()                    { *m = Whatzit{} }
func (m *Whatzit) String() string            { return proto.CompactTextString(m) }
func (*Whatzit) ProtoMessage()               {}
func (*Whatzit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Whatzit) GetGyzmeau() []*jhump_protoreflect_desc.Bar {
	if m != nil {
		return m.Gyzmeau
	}
	return nil
}

type GroupX struct {
	Groupxi          *int64  `protobuf:"varint,1041,opt,name=groupxi" json:"groupxi,omitempty"`
	Groupxs          *string `protobuf:"bytes,1042,opt,name=groupxs" json:"groupxs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GroupX) Reset()                    { *m = GroupX{} }
func (m *GroupX) String() string            { return proto.CompactTextString(m) }
func (*GroupX) ProtoMessage()               {}
func (*GroupX) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GroupX) GetGroupxi() int64 {
	if m != nil && m.Groupxi != nil {
		return *m.Groupxi
	}
	return 0
}

func (m *GroupX) GetGroupxs() string {
	if m != nil && m.Groupxs != nil {
		return *m.Groupxs
	}
	return ""
}

var E_Otl = &proto.ExtensionDesc{
	ExtendedType:  (*desc_test_nopkg.TopLevel)(nil),
	ExtensionType: (*desc_test_nopkg.TopLevel)(nil),
	Field:         100,
	Name:          "desc_test.otl",
	Tag:           "bytes,100,opt,name=otl",
	Filename:      "desc_test2.proto",
}

var E_Groupx = &proto.ExtensionDesc{
	ExtendedType:  (*desc_test_nopkg.TopLevel)(nil),
	ExtensionType: (*GroupX)(nil),
	Field:         104,
	Name:          "desc_test.groupx",
	Tag:           "group,104,opt,name=GroupX,json=groupx",
	Filename:      "desc_test2.proto",
}

func init() {
	proto.RegisterType((*Frobnitz)(nil), "desc_test.Frobnitz")
	proto.RegisterType((*Whatchamacallit)(nil), "desc_test.Whatchamacallit")
	proto.RegisterType((*Whatzit)(nil), "desc_test.Whatzit")
	proto.RegisterType((*GroupX)(nil), "desc_test.GroupX")
	proto.RegisterExtension(E_Otl)
	proto.RegisterExtension(E_Groupx)
}

func init() { proto.RegisterFile("desc_test2.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xdb, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x39, 0xc9, 0xd6, 0xae, 0x67, 0x62, 0x74, 0x7e, 0x00, 0x33, 0x40, 0x0a, 0xd5, 0x84,
	0x22, 0x21, 0xa5, 0xd4, 0x45, 0x05, 0x15, 0x09, 0xa9, 0x45, 0x1b, 0x7b, 0x18, 0x3c, 0x58, 0xe3,
	0x22, 0x5e, 0x26, 0x37, 0x71, 0x2e, 0x2c, 0xa9, 0xa3, 0xc4, 0x41, 0xd0, 0xff, 0x81, 0x07, 0xf8,
	0x8b, 0x51, 0x9c, 0xde, 0x26, 0xa1, 0x69, 0x4f, 0xf6, 0xf9, 0xbe, 0xdf, 0x67, 0xfb, 0xe8, 0x18,
	0xbb, 0x81, 0x2c, 0xfd, 0x4b, 0x2d, 0x4b, 0xcd, 0xbc, 0xbc, 0x50, 0x5a, 0x91, 0xce, 0x5a, 0x39,
	0xda, 0x98, 0x83, 0xc6, 0x3c, 0x7a, 0x90, 0x5f, 0x45, 0xfd, 0xb5, 0x7a, 0x99, 0x5f, 0x45, 0x4b,
	0xe3, 0xd1, 0x5c, 0x5d, 0xb7, 0x4c, 0xdd, 0x98, 0xbd, 0xdf, 0x36, 0xee, 0x9d, 0x16, 0x6a, 0x36,
	0x4f, 0xf4, 0x82, 0x1c, 0x23, 0x08, 0x0a, 0x0e, 0xb8, 0xfb, 0xec, 0xbe, 0xb7, 0xe6, 0xbd, 0x0b,
	0x59, 0xea, 0x0f, 0xb2, 0x2c, 0x45, 0x24, 0x39, 0x08, 0xf2, 0x1c, 0x61, 0x46, 0x2d, 0x43, 0x3d,
	0xd9, 0xa2, 0x26, 0x73, 0xa5, 0x63, 0x59, 0x5c, 0x83, 0x67, 0x64, 0x84, 0x96, 0x3f, 0xa0, 0xb6,
	0xa1, 0x8f, 0xff, 0x7f, 0xa6, 0xf7, 0x51, 0x96, 0x5a, 0x06, 0xcb, 0xea, 0xec, 0x0e, 0xb7, 0xfc,
	0x01, 0x19, 0xa2, 0xe5, 0x33, 0xba, 0xe3, 0x80, 0x7b, 0xc0, 0x9e, 0xde, 0x98, 0x3b, 0x99, 0x57,
	0x99, 0x09, 0x31, 0xc2, 0x10, 0x02, 0xba, 0x7b, 0xfb, 0xbb, 0x38, 0x04, 0xe4, 0x15, 0x82, 0xa4,
	0xad, 0x5b, 0xde, 0x33, 0x6e, 0x7d, 0x9e, 0x9c, 0x7f, 0x3a, 0x61, 0x1c, 0x24, 0xe9, 0x22, 0x84,
	0xb4, 0xed, 0xd8, 0x6e, 0x67, 0x6a, 0x51, 0xe0, 0x10, 0x92, 0x2e, 0x5a, 0xd1, 0x80, 0xee, 0x39,
	0xe0, 0xee, 0x9e, 0x01, 0xb7, 0xa2, 0x81, 0x51, 0x18, 0xed, 0x38, 0xe0, 0x1e, 0x1a, 0x85, 0x19,
	0x65, 0x48, 0xd1, 0x01, 0xf7, 0xae, 0x51, 0x86, 0xd3, 0x5d, 0xb4, 0xc5, 0xcc, 0xaf, 0x97, 0x40,
	0x86, 0xbd, 0x77, 0x78, 0xef, 0x4b, 0x2c, 0xb4, 0x1f, 0x8b, 0x4c, 0xf8, 0x22, 0x4d, 0x13, 0x4d,
	0x5e, 0xe0, 0x4e, 0xa8, 0x54, 0x49, 0xc1, 0xb1, 0xdc, 0x03, 0xf6, 0xd8, 0xfb, 0x1e, 0x57, 0x59,
	0xde, 0x8c, 0xaf, 0x90, 0x61, 0x2a, 0x7d, 0x6d, 0xde, 0xed, 0x9d, 0x2a, 0xc5, 0x0d, 0xd9, 0x9b,
	0x60, 0xbb, 0x3e, 0x64, 0x91, 0x68, 0x32, 0xc2, 0x76, 0xf4, 0x6b, 0x91, 0x49, 0x51, 0x51, 0x70,
	0x6c, 0x77, 0xff, 0x86, 0xfc, 0x54, 0x14, 0x7c, 0x05, 0xf7, 0xde, 0x62, 0xeb, 0x7d, 0xa1, 0xaa,
	0xfc, 0x2b, 0x79, 0x88, 0xed, 0xa8, 0xde, 0xfd, 0x4c, 0xe8, 0x9f, 0xba, 0x37, 0x9b, 0xaf, 0xea,
	0x8d, 0x55, 0xd2, 0xbf, 0xb5, 0xd5, 0x59, 0x59, 0xe5, 0xf8, 0x19, 0xda, 0x4a, 0xa7, 0xa4, 0xe3,
	0x5d, 0xa8, 0xfc, 0x5c, 0xfe, 0x90, 0x29, 0x0d, 0xcc, 0x5c, 0x36, 0x02, 0xaf, 0x81, 0xf1, 0x4b,
	0x6c, 0x35, 0x91, 0x6d, 0x34, 0x76, 0xc0, 0x45, 0x76, 0xb8, 0x35, 0x8e, 0xe6, 0x2d, 0x7c, 0xc9,
	0x4e, 0x5f, 0x7f, 0x1b, 0x45, 0x89, 0x8e, 0xab, 0x99, 0xe7, 0xab, 0xac, 0x6f, 0x1a, 0xea, 0x6f,
	0x37, 0x64, 0x3e, 0xfb, 0xe6, 0xc7, 0xbf, 0x59, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x31,
	0xb9, 0x7d, 0x0e, 0x55, 0x03, 0x00, 0x00,
}
