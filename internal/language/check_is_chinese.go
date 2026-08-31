package language

import (
	"regexp"
)

// IsChinese checks if a single character is chinese
func IsChinese(character string) bool {
	reg := regexp.MustCompile("^[\u4e00-\u9fa5\u3007]+$")
	return reg.MatchString(character)
}
