package comma

import "strings"

func Comma(s string) string{
	dotstring := ""

	if dot := strings.LastIndex(s, "."); dot >= 0{
		dotstring = s[dot:]
		s = s[:dot-1]
	}

	n := len(s)

	if n<= 3{
		return s
	}

	return Comma(s[:n-3]) + "," + s[n-3:] + dotstring
}