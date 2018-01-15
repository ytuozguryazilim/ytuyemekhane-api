package util

import (
	"errors"
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

// IsYearSuitable :
func IsYearSuitable(value int) error {
	if value > 2010 && value < 2030 {
		return nil
	}
	return errors.New("2010 ile 2030 arasinda bir yil secin:)")
}

// IsMonthSuitable :
func IsMonthSuitable(value int) error {
	if value > 0 && value < 13 {
		return nil
	}
	return errors.New("1 ile 12 arasinda bir ay secin:)")
}

// IsDaySuitable :
func IsDaySuitable(value int) error {
	if value > 0 && value < 32 {
		return nil
	}
	return errors.New("1 ile 31 arasinda bir gun secin:)")
}
