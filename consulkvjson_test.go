package consulkvjson

import (
	"reflect"
	"testing"
)

var kvsPass = []*KV{
	&KV{key: "root", value: "I am a string"},
	&KV{key: "how/about/some/depth", value: "1"},
	&KV{key: "how/about/some/more_depth", value: "2"},
}

var jsPass = []byte(`{
	"hello": "world",
	"count": 10,
	"list" : [{
		"nested_in_list" : 1
	}, "2", "3"],
	"nested" : {
		"boolean" : false
	},
	"map" : {
		"hello": "world"
	}
}`)

var jsPassKVs = []*KV{
	&KV{key: "hello", value: "world"},
	&KV{key: "count", value: "10"},
	&KV{key: "list/0/nested_in_list", value: "1"},
	&KV{key: "list/1", value: "2"},
	&KV{key: "list/2", value: "3"},
	&KV{key: "nested/boolean", value: "false"},
	&KV{key: "map/hello", value: "world"},
}

func TestToJSONPass(t *testing.T) {
	want := string(`{"how":{"about":{"some":{"depth":"1","more_depth":"2"}}},"root":"I am a string"}`)
	json, err := ToJSON(kvsPass)
	if err != nil {
		t.Errorf("Failure in ToJSON")
	}
	if string(json) != want {
		t.Errorf("Could not translate KVs to JSON")
	}
	// log.Printf("%s", json)
}

func TestToKVs(t *testing.T) {
	kvs, err := ToKVs(jsPass)
	if err != nil {
		t.Errorf("There was an error: %v", err)
	}
	kvsCorrectMap := make(map[KV]interface{})
	kvsMap := make(map[KV]interface{})

	for _, kv := range jsPassKVs {
		kvsCorrectMap[*kv] = true
	}

	for _, kv := range kvs {
		kvsMap[*kv] = true
	}

	if !reflect.DeepEqual(kvsCorrectMap, kvsMap) {
		t.Errorf("Could not translate JSON to KVs")
	}

	// for _, kv := range kvs {
	// 	log.Printf("%s, %s", kv.key, kv.value)
	// }
}
