package util

import (
	"encoding/json"
	"fmt"
)

// Marshal and unmarshal to create a deep copy of a deployment
func DeepCopy(obj interface{}) (interface{}, error) {

	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var deploy interface{}
	if err := json.Unmarshal(data, &deploy); err != nil {
		return nil, err
	}
	return deploy, nil
}

// Change value of specified json field
func SetStringField(jsonObj interface{}, field string, val string) (interface{}, error) {
	if field == "" {
		return nil, fmt.Errorf("field not specified. What field to set?")
	}

	newMap, ok := jsonObj.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid json %v", jsonObj), nil
	}

	newMap[field] = val

	return newMap, nil
}
