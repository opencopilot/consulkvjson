package consulkvjson

import (
	"encoding/json"
	"strconv"
	"strings"
)

// KV represents a KV pair
type KV struct {
	key   string
	value string
}

func traverse(path string, j interface{}) ([]*KV, error) {
	kvs := make([]*KV, 0)

	switch j.(type) {
	case []interface{}:
		for sk, sv := range j.([]interface{}) {
			skvs, err := traverse(path+"/"+strconv.Itoa(sk), sv)
			if err != nil {
				return nil, err
			}
			kvs = append(kvs, skvs...)
		}
	case map[string]interface{}:
		for sk, sv := range j.(map[string]interface{}) {
			skvs, err := traverse(path+"/"+sk, sv)
			if err != nil {
				return nil, err
			}
			kvs = append(kvs, skvs...)
		}
	case float64:
		kvs = append(kvs, &KV{key: path, value: strconv.FormatFloat(j.(float64), 'f', -1, 64)})
	case bool:
		kvs = append(kvs, &KV{key: path, value: strconv.FormatBool(j.(bool))})
	default:
		kvs = append(kvs, &KV{key: path, value: j.(string)})
	}

	return kvs, nil
}

// ToKVs takes a json byte array and returns a list of KV pairs where each key is a path in the Consul KV store
func ToKVs(jsonData []byte) ([]*KV, error) {
	var i interface{}
	err := json.Unmarshal(jsonData, &i)
	if err != nil {
		return nil, err
	}
	m := i.(map[string]interface{})

	return traverse("", m)
}

// ToJSON converts a list of KVs to a JSON tree
func ToJSON(kvs []*KV) ([]byte, error) {
	m := make(map[string]interface{})
	for _, kv := range kvs {
		path := strings.Split(kv.key[1:], "/")
		var parent = m
		for s, segment := range path {
			if s == len(path)-1 {
				parent[segment] = string(kv.value)
			} else {
				if parent[segment] == nil {
					parent[segment] = make(map[string]interface{})
				}
				parent = parent[segment].(map[string]interface{})
			}
		}
	}
	return json.Marshal(m)
}
