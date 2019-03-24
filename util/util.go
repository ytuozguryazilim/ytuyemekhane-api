package util

import (
	"strconv"
)

// StringToInt : "string" turundeki degiskeni "int" turune cevirir.
func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}
