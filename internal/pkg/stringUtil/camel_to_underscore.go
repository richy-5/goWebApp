package stringUtil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelToUnderscore case to CamelToUnderscore style.
func CamelToUnderscore(camelStr string) string {
	l := utf8.RuneCountInString(camelStr)
	ss := strings.Split(camelStr, "")

	// we just care about the key of idx map,
	// the value of map is meaningless
	idx := make(map[int]int, 1)

	var rs []rune
	for _, s := range camelStr {
		rs = append(rs, []rune(string(s))...)
	}

	for i := l - 1; i >= 0; {
		if unicode.IsUpper(rs[i]) {
			var start, end int
			end = i
			j := i - 1
			for ; j >= 0; j-- {
				if unicode.IsLower(rs[j]) {
					start = j + 1
					break
				}
			}
			// handle the case: "BBC" or "AaBBB" case
			if end == l-1 {
				idx[start] = 1
			} else {
				if start == end {
					// value=1 is meaningless
					idx[start] = 1
				} else {
					idx[start] = 1
					idx[end] = 1
				}
			}
			i = j - 1
		} else {
			i--
		}
	}

	for i := l - 1; i >= 0; i-- {
		ss[i] = strings.ToLower(ss[i])
		if _, ok := idx[i]; ok && i != 0 {
			ss = append(ss[0:i], append([]string{"_"}, ss[i:]...)...)
		}
	}

	return strings.Join(ss, "")
}
