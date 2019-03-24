package util

import (
	"strconv"
	"strings"
)

// ClearText : gelen datayi temizler.
func ClearText(orginal string) string {
	return strings.Trim(orginal, "\n\n ")
}

// IntToString : "int" turundeki degiskeni "string" turune cevirir.
func IntToString(value int) string {
	return strconv.Itoa(value)
}

// StringToInt : "string" turundeki degiskeni "int" turune cevirir.
func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

// CreateURL : gelen verilerden arda arda birlestirip "string" turunde url donuyoruz.
func CreateURL(values ...interface{}) string {
	var result string
	for _, value := range values {
		switch value.(type) {
		case int:
			result += IntToString(value.(int))
		case string:
			result += value.(string)
		}
		result += "/"
	}
	return result
}
