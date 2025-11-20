package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomFromRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := RandomFromRange(1, 10)
		assert.GreaterOrEqual(t, result, 1)
		assert.Less(t, result, 10)
	}
}

func TestRandomFromMap(t *testing.T) {
	input := map[string]string{"a": "1", "b": "2", "c": "3"}
	result, err := RandomFromMap(input)
	assert.NoError(t, err)
	assert.Contains(t, input, result)
}

func TestRandomFromList(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := RandomFromList(input)
	assert.Contains(t, input, result)
}

func TestDeepCopy(t *testing.T) {
	input := map[string]interface{}{"key": "value"}
	copy, err := DeepCopy(input)
	assert.NoError(t, err)
	assert.Equal(t, input, copy)

	// Ensure it's a deep copy
	copy.(map[string]interface{})["key"] = "new value"
	assert.NotEqual(t, input, copy)
}
func TestDeepCopyWithNestedObjects(t *testing.T) {
	input := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"nestedKey1": "nestedValue1",
			"nestedKey2": []interface{}{"item1", "item2"},
		},
	}

	copy, err := DeepCopy(input)
	assert.NoError(t, err)
	assert.Equal(t, input, copy)

	// Ensure it's a deep copy
	copy.(map[string]interface{})["key2"].(map[string]interface{})["nestedKey1"] = "modifiedValue"
	assert.NotEqual(t, input, copy)

	// Ensure nested value in the original map hasnt changed
	assert.Equal(t, "nestedValue1", input["key2"].(map[string]interface{})["nestedKey1"])
}

func TestBytesToTiB(t *testing.T) {
	// 1 TiB in bytes
	const oneTiB = uint64(1 << 40)

	// Test cases
	tests := []struct {
		bytes    uint64
		expected float64
	}{
		{bytes: 0, expected: 0},
		{bytes: oneTiB, expected: 1},
		{bytes: oneTiB * 2, expected: 2},
		{bytes: oneTiB / 2, expected: 0.5},
		{bytes: oneTiB + (oneTiB / 2), expected: 1.5},
	}

	for _, test := range tests {
		result := BytesToTiB(test.bytes)
		assert.InDelta(t, test.expected, result, 0.0001, "BytesToTiB(%d) should be approximately %f", test.bytes, test.expected)
	}
}
