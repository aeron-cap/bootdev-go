package iterations

import "strings"

func Repeat(text string, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteString(text)
	}
	return repeated.String()
}