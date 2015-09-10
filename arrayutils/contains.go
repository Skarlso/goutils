package arrayutils

//ContainsInt checks wether an element is in an array.
func ContainsInt(arr []int, element interface{}) bool {
	for _, elem := range arr {
		if elem == element {
			return true
		}
	}
	return false
}

//ContainsString checks wether an element is in an array.
func ContainsString(arr []string, element interface{}) bool {
	for _, elem := range arr {
		if elem == element {
			return true
		}
	}
	return false
}
