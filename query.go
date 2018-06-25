package goserve

import (
	"strings"
)

func Filter(str string) string {
	return strings.ToLower(str)
}

func Item(str ...string) (total string) {
	for _, num := range str {
		total += strings.ToLower(num) + ","
	}
	total = total[:len(total)-1]
	return
}

func AND() string {
	return "^"
}

func OR() string {
	return "^OR"
}

func IS(str string) string {
	return "=" + str
}

func ISNOT(str string) string {
	return "!=" + str
}

func ISONEOF(str ...string) string {
	return "IN" + Item(str...)
}

func ISNOTONEOF(str ...string) string {
	return "NOT%2520IN" + Item(str...)
}

func ISEMPTY() string {
	return "ISEMPTY"
}

func ISNOTEMPTY() string {
	return "ISNOTEMPTY"
}

func ISLESSTHAN(str string) string {
	return "<" + str
}

func ISLESSTHANOREQUALS(str string) string {
	return "<=" + str
}

func ISGREATERTHAN(str string) string {
	return ">" + str
}

func ISGREATERTHANOREQUALS(str string) string {
	return ">=" + str
}

func ISBETWEEN(str1, str2 string) string {
	return "BETWEEN" + str1 + "@" + str2
}

func ISANYTHING() string {
	return "ANYTHING"
}

func ISSAMEAS(str string) string {
	return "SAMEAS" + str
}

func ISDIFFERENTFROM(str string) string {
	return "NSAMEAS" + str
}

func ORDERBY(str string) string {
	return "^ORDERBY" + str
}

func ISLIKE(str string) string {
	return "LIKE" + str
}

func ISNOTLIKE(str string) string {
	return "NOT%20LIKE" + str
}
