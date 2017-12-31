package stringutil

// Reverse returns a reversed version of the input string.
func Reverse(str string) string {
	length := len(str) - 1
	reverse := ""
	for i := range str {
		reverse += string(str[length-i])
	}
	return reverse
}

// IsPalindrome returns a boolean indicating wether a string is a palindrome or not.
func IsPalindrome(str string) bool {
	length := len(str) - 1
	for i := 0; i < length/2; i++ {
		if string(str[i]) != string(str[length-i]) {
			return false
		}
	}
	return true
}
