// Warmup-1 exercises from codingbat.com re-worked for Golang practice
// Simple warmup problems to get started, no loops
// package warmup1
package warmup1

import "strings"

// SleepIn takes two boolean arguments (for weekday and vacation) and returns a
// boolean. The parameter w is true if it is a weekday, and the parameter v is
// true if we are on vacation. We sleep in if it is not a weekday or we're on
// vacation. Return true if we sleep in.
func SleepIn(weekday, vacation bool) bool {
	// input: (weekday, vacation bool) -> output: boolean
	if !weekday || vacation {
		return true
	} else {
		return false
	}
}

// We have two monkeys, a and b, and the parameters a and b indicate
// if each is smiling. We are in trouble if they are both smiling or if neither
// of them is smiling. Return True if we are in trouble.
func MonkeyTrouble(a, b bool) bool {
	if a && b {
		return true
	} else if !a && !b {
		return true
	} else {
		return false
	}
}

// Given two int values, return their sum. Unless the two values are the same, then
// return double their sum.
func SumDouble(a, b int) int {
	if a == b {
		return 2 * (a + b)
	}
	return a + b
}

// Given an int n, return the absolute difference between n and 21, except return
// double the absolute difference if n is over 21.
func Diff21(n int) int {
	if n < 21 {
		return -(n - 21)
	} else {
		return 2 * (n - 21)
	}
}

// We have a loud talking parrot. The "hour" parameter is the current hour time in the range 0..23.
// We are in trouble if the parrot is talking and the hour is before 7 or after 20.
// Return True if we are in trouble.
func ParrotTrouble(t bool, h int) bool {
	if (h < 7 || h > 20) && t {
		return true
	}
	return false
}

// Given 2 ints, a and b, return true if one if them is 10 or if their sum is 10.
func Makes10(a, b int) bool {
	if a+b == 10 {
		return true
	} else if a == 10 || b == 10 {
		return true
	}
	return false
}

// Given an int n, return true if it is within 10 of 100 or 200.
func NearHundred(n int) bool {
	if n+10 >= 100 && n-10 <= 100 {
		return true
	}

	if n-10 <= 200 && n+10 >= 200 {
		return true
	}

	return false
}

// Given 2 int values, return true if one is negative and one is positive. Except if the parameter
// "negative" is true, then return true only if both are negative.
func PosNeg(a int, b int, neg bool) bool {
	if neg && (a < 0 && b < 0) {
		return true
	} else if a < 0 && b > 0 {
		return true
	} else if a > 0 && b < 0 {
		return true
	}
	return false
}

// Given a string, return a new string where "not " has been added to the front. However, if the
// string already begins with "not", return the string unchanged.
func NotString(s string) string {
	if strings.HasPrefix(s, "not") {
		return s
	}
	return "not " + s
}

// Given a non-empty string and an int n, return a new string where the char at index n has been removed.
// The value of n will be a valid index of a char in the original string (i.e. n will be in the range
// 0..len(str)-1 inclusive).
func MissingChar(s string, n int) string {
	l := strings.Split(s, "")
	r := ""
	i := 0
	for _, v := range l {
		if i != n {
			r += v
		}
		i++
	}
	return r
}

// Given a string, return a new string where the first and last chars have been exchanged.
func FrontBack(s string) string {

	l := strings.Split(s, "")
	if len(l) == 0 || s == "" {
		return s
	}

	tmp := l[0]
	l[0] = l[len(l)-1]
	l[len(l)-1] = tmp

	r := ""
	for _, v := range l {
		r += v
	}

	return r
}

// Given a string, we'll say that the front is the first 3 chars of the string. If the string
// length is less than 3, the front is whatever is there. Return a new string which is 3 copies of the front.
func Front3(s string) string {

	l := strings.Split(s, "")
	r := ""
	if len(l) >= 3 {
		t := l[:3]
		for _, v := range t {
			r += v
		}
	} else {
		for _, v := range l {
			r += v
		}
	}

	return r + r + r
}
