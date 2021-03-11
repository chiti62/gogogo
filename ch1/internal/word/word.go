package word

import (
	"strconv"
	"strings"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func WordCountStr(s string) string {
	return strconv.Itoa(len(s))
}
