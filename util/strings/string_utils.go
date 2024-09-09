package strings

import "strings"

func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func Join(delimiter string, items []string) string {
	return strings.Join(items, delimiter)
}

func Split(s string, delimiter string) []string {
	return strings.Split(s, delimiter)
}

func IndexOf(s string, substr string) int {
	return strings.Index(s, substr)
}

func LastIndexOf(s string, substr string) int {
	return strings.LastIndex(s, substr)
}

func Contains(s string, substr string) bool {
	return strings.Contains(s, substr)
}

func StartsWith(s string, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func EndsWith(s string, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
