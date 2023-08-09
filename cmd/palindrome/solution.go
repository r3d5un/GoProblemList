package palindrome

func IsPalindrome[T comparable](list []T) bool {
	for i := 0; i < len(list); i++ {
		if list[i] != list[len(list)-i-1] {
			return false
		}
	}

	return true
}
