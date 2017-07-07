package internal

// IsAlpha checks if a string is comprised of only alphabetical characters.
func IsAlpha(str string) bool {
	for i := range str {
		if str[i] < 'A' || str[i] > 'z' {
			return false
		} else if str[i] > 'Z' && str[i] < 'a' {
			return false
		}
	}
	return true
}
