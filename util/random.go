package util

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

// returns random number from specified range, max excluded
func RandomFromRange(min int, maxExcluded int) int {
	return rand.Intn(maxExcluded-min) + min
}

// returns random number from specified range, max excluded
func RandomUuid() string {
	return uuid.NewString()
}

// returns random key from map
func RandomFromMap[Key comparable, Value any](m map[Key]Value) (Key, error) {

	// Pick a random index
	randomIdx := rand.Intn(len(m))
	var idx int = 0

	for key := range m {
		if idx == randomIdx {
			return key, nil
		}
		idx++
	}

	//error case, nothing found
	var zero Key
	return zero, fmt.Errorf("no key found in map")

}

func RandomFromList[T any](list []T) T {

	// Pick a random index
	randomIdx := rand.Intn(len(list))

	return list[randomIdx]

}
