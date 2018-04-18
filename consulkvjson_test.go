package consulkvjson

import (
	"reflect"
	"testing"
)

var kvsPass = []*KV{
	&KV{Key: "root", Value: "I am a string"},
	&KV{Key: "how/about/some/depth", Value: "1"},
	&KV{Key: "how/about/some/more_depth", Value: "2"},
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
	},
	"null_val": null
}`)

var jsPassKVs = []*KV{
	&KV{Key: "hello", Value: "world"},
	&KV{Key: "count", Value: "10"},
	&KV{Key: "list/0/nested_in_list", Value: "1"},
	&KV{Key: "list/1", Value: "2"},
	&KV{Key: "list/2", Value: "3"},
	&KV{Key: "nested/boolean", Value: "false"},
	&KV{Key: "map/hello", Value: "world"},
	&KV{Key: "null_val", Value: ""},
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
