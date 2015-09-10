package arrayutils

//ReverseString reverses an array of strings.
func ReverseString(arr []string) []string {
	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

//ReverseInt reverses an array of integers
func ReverseInt(arr []int) []int {
	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
