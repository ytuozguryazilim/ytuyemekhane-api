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

// CreateVirtualPATH :
func CreateVirtualPATH(values ...int) string {
	var result string
	for _, value := range values {
		result = result + "/" + IntToString(value)
	}
	return result
}
