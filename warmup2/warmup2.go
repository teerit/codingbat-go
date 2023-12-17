// Exercises from codingbat.com ported to Golang.
package warmup2

// Given an slice of ints, return true if the sequence of numbers 1, 2, 3
// appears in the slice somewhere.
func Array123(s []int) bool {
	c := 0
	for i := 0; i < len(s); i++ {
		if c == 3 {
			return true
		} else if s[i] == 1 {
			c = 1
		} else if s[i] == c+1 {
			c++
		} else {
			return false
		}
	}

	if c == 3 {
		return true
	}

	return false
}

// Given a slice of ints, return the number of 9's in the slice.
func ArrayCount9(s []int) int {
	cnt := 0
	for _, v := range s {
		if v == 9 {
			cnt++
		}
	}
	return cnt
}

// Given a slice of ints, return true if one of the first 4 elements in the
// slice is a 9. The slice length may be less than 4.
func ArrayFront(s []int) bool {

	if len(s) < 4 {
		for i := 0; i < len(s); i++ {
			if s[i] == 9 {
				return true
			}
		}
	} else {
		for i := 0; i < 4; i++ {
			if s[i] == 9 {
				return true
			}
		}
	}

	return false
}

// Given a string and a non-negative int n, we'll say that the front of the
// string is the first 3 chars, or whatever is there if the string is less
// than length 3. Return n copies of the front
func FrontTimes(s string, n int) string {
	r := ""
	chars := []rune(s)
	if len(chars) < 3 {
		for i := 0; i < n; i++ {
			r += s
		}
		return r
	} else {
		c := ""
		for i := 0; i < 3; i++ {
			c += string(chars[i])
		}

		for i := 0; i < n; i++ {
			r += c
		}
		return r
	}
}

// Given a string, return the count of the number of times that a substring length
// 2 appears in the string and also as the last 2 chars of the string, so
// "hixxxhi" yields 1 (we won't count the end substring).
func Last2(s string) int {
	return 0
}

// Given a string, return a new string made of every other char starting with
// the first, so "Hello" yields "Hlo".
func StringBits(s string) string {
	return ""
}

// Given 2 strings, a and b, return the number of the positions where they contain
// the same length 2 substring. So "xxcaazz" and "xxbaaz" yields 3, since the
// "xx", "aa", and "az" substrings appear in the same place in both strings.
func StringMatch(a, b string) int {
	return 0
}

// Given a non-empty string like "Code" return a string like "CCoCodCode".
func StringSplosion(s string) string {
	return ""
}

// Given a string and a non-negative int n, return a larger string that is n
// copies of the original string.
func StringTimes(s string, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}
