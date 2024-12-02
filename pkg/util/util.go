package util

import "encoding/json"

func stringToMap(s string) []map[string]interface{} {
	var value []map[string]interface{}
	err := json.Unmarshal([]byte(s), &value)
	if err != nil {
		panic(err)
	}
	return value
}
