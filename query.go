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
	str = strings.Replace(str, " ", "%20", -1)
	return "=" + str
}

func ISNOT(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
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
	str = strings.Replace(str, " ", "%20", -1)
	return "<" + str
}

func ISLESSTHANOREQUALS(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "<=" + str
}

func ISGREATERTHAN(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return ">" + str
}

func ISGREATERTHANOREQUALS(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return ">=" + str
}

func ISBETWEEN(str1, str2 string) string {
	return "BETWEEN" + str1 + "@" + str2
}

func ISANYTHING() string {
	return "ANYTHING"
}

func ISSAMEAS(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
	return "SAMEAS" + str
}

func ISDIFFERENTFROM(str string) string {
	str = strings.Replace(str, " ", "%20", -1)
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

func ON(str string) string {
	return "ON" + str
}

func NOTON(str string) string {
	return "NOTON" + str
}

func BEFORE(str string) string {
	return "<" + str
}

func ATORBEFORE(str string) string {
	return "<=" + str
}

func AFTER(str string) string {
	return ">" + str
}

func ATORAFTER(str string) string {
	return ">=" + str
}

func BETWEEN(str string) string {
	return "BETWEEN" + str
}

//MORETHAN is for date-type fields only!
func MORETHAN(str string) string {
	return "MORETHAN" + str
}

//LESSTHAN is for date-type fields only!
func LESSTHAN(str string) string {
	return "LESSTHAN" + str
}
