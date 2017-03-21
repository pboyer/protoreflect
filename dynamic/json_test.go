package dynamic

import (
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestUnaryFieldsJSON(t *testing.T) {
	jsonTranslationParty(t, unaryFieldsMsg)
}

func TestRepeatedFieldsJSON(t *testing.T) {
	jsonTranslationParty(t, repeatedFieldsMsg)
}

func TestMapKeyFieldsJSON(t *testing.T) {
	// translation party wants deterministic marshalling to bytes
	sort_map_keys = true
	defer func() {
		sort_map_keys = false
	}()

	jsonTranslationParty(t, mapKeyFieldsMsg)
}

func TestMapValueFieldsJSON(t *testing.T) {
	// translation party wants deterministic marshalling to bytes
	sort_map_keys = true
	defer func() {
		sort_map_keys = false
	}()

	jsonTranslationParty(t, mapValueFieldsMsg)
}

func TestUnknownFieldsJSON(t *testing.T) {
	// TODO
}

func TestExtensionFieldsJSON(t *testing.T) {
	// TODO
}

func TestLenientParsingJSON(t *testing.T) {
	// TODO
	// include optional commas, different ways to indicate extension names, and array notation for repeated fields
}

func jsonTranslationParty(t *testing.T, msg proto.Message) {
	doTranslationParty(t, msg,
		func(pm proto.Message) ([]byte, error) {
			return []byte(proto.MarshalJSONString(pm)), nil
		},
		func(b []byte, pm proto.Message) error {
			return proto.UnmarshalJSON(string(b), pm)
		},
		(*Message).MarshalJSON, (*Message).UnmarshalJSON)
}
