package arrayutils

//Contains checks wether an element is in an array.
func Contains(arr []interface{}, element interface{}) bool {
    for _, elem := range arr {
        if elem == element {
            return true
        }
    }
    return false
}
