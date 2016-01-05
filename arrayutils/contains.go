package arrayutils

import "reflect"

//ContainsInt checks whether an element is in a slice.
func ContainsInt(arr []int, element interface{}) bool {
	for _, elem := range arr {
		if elem == element {
			return true
		}
	}
	return false
}

//ContainsString checks whether an element is in a slice.
func ContainsString(arr []string, element interface{}) bool {
	for _, elem := range arr {
		if elem == element {
			return true
		}
	}
	return false
}

//ContainsByteSlice check whether []byte is in a slice of []bytes
func ContainsByteSlice(arr [][]byte, element []byte) bool {
	for _, elem := range arr {
		if reflect.DeepEqual(elem, element) {
			return true
		}
	}
	return false
}
