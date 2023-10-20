package util

import (
	"strings"
	"unicode"
)

// PrintableString removes non-graphic characters from the given string.
// Removable characters are the ones for which unicode.IsGraphic() returns false.
func PrintableString(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
}

// Title converts a string into camel case starting with a upper case letter and keep uppercase letter.
func Title(s string) string {
	return camelCase(s, true, true)
}

// UpperCamelCase converts a string into camel case starting with a upper case letter.
func UpperCamelCase(s string) string {
	return camelCase(s, true, false)
}

// LowerCamelCase converts a string into camel case starting with a lower case letter.
func LowerCamelCase(s string) string {
	return camelCase(s, false, false)
}

func camelCase(s string, upper bool, keepUpper bool) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s))

	stringIter(s, func(prev, curr, next rune) {
		if !isDelimiter(curr) {
			if isDelimiter(prev) || (upper && prev == 0) {
				buffer = append(buffer, toUpper(curr))
			} else if isLower(prev) || keepUpper {
				buffer = append(buffer, curr)
			} else {
				buffer = append(buffer, toLower(curr))
			}
		}
	})

	return string(buffer)
}

// KebabCase converts a string into kebab case.
func KebabCase(s string) string {
	return delimiterCase(s, '-', false)
}

// UpperKebabCase converts a string into kebab case with capital letters.
func UpperKebabCase(s string) string {
	return delimiterCase(s, '-', true)
}

// SnakeCase converts a string into snake case.
func SnakeCase(s string) string {
	return delimiterCase(s, '_', false)
}

// UpperSnakeCase converts a string into snake case with capital letters.
func UpperSnakeCase(s string) string {
	return delimiterCase(s, '_', true)
}

func delimiterCase(s string, delimiter rune, upperCase bool) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s)+3)
	adjustCase := toLower
	if upperCase {
		adjustCase = toUpper
	}
	var prev rune
	var curr rune
	for _, next := range s {
		if isDelimiter(curr) {
			if !isDelimiter(prev) {
				buffer = append(buffer, delimiter)
			}
		} else if isUpper(curr) {
			if isLower(prev) || (isUpper(prev) && isLower(next)) {
				buffer = append(buffer, delimiter)
			}
			buffer = append(buffer, adjustCase(curr))
		} else if curr != 0 {
			buffer = append(buffer, adjustCase(curr))
		}
		prev = curr
		curr = next
	}
	if len(s) > 0 {
		if isUpper(curr) && isLower(prev) && prev != 0 {
			buffer = append(buffer, delimiter)
		}
		buffer = append(buffer, toLower(curr))
	}

	return string(buffer)
}

// isLower checks if a character is lower case. More precisely it evaluates if it is
// in the range of ASCII character 'a' to 'z'.
func isLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

// toLower converts a character in the range of ASCII characters 'A' to 'Z' to its lower
// case counterpart. Other characters remain the same.
func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

// isLower checks if a character is upper case. More precisely it evaluates if it is
// in the range of ASCII characters 'A' to 'Z'.
func isUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// toLower converts a character in the range of ASCII characters 'a' to 'z' to its lower
// case counterpart. Other characters remain the same.
func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

// isSpace checks if a character is some kind of whitespace.
func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isDelimiter checks if a character is some kind of whitespace or '_' or '-'.
func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}

// iterFunc is a callback that is called fro a specific position in a string. Its arguments are the
// rune at the respective string position as well as the previous and the next rune. If curr is at the
// first position of the string prev is zero. If curr is at the end of the string next is zero.
type iterFunc func(prev, curr, next rune)

// stringIter iterates over a string, invoking the callback for every single rune in the string.
func stringIter(s string, callback iterFunc) {
	var prev rune
	var curr rune
	for _, next := range s {
		if curr == 0 {
			prev = curr
			curr = next
			continue
		}

		callback(prev, curr, next)

		prev = curr
		curr = next
	}

	if len(s) > 0 {
		callback(prev, curr, 0)
	}
}

func StringsJoin(strs ...string) string {
	var n int
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}
	if n <= 0 {
		return ""
	}
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	builder.Grow(n)
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}
