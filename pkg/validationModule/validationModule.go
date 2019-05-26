package validationModule

import (
	"fmt"
	"strings"
)

func IsSAlphabetAndDot(s string) bool {
	for _, r := range s {
		//if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		if (r < 'a' || r > 'z') && r != '.' {
			return false
		}
	}
	return true
}
func CheckLocaleValue(locale string) bool {
	//ko, en, ja
	result := false
	fmt.Println("CheckLocaleValue ", locale)
	if strings.EqualFold(locale, "ko") || strings.EqualFold(locale, "en") || strings.EqualFold(locale, "ja") {
		result = true
	}
	fmt.Println("CheckLocaleValue ", result)
	return result
}
