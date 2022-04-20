package masutils

import (
	"fmt"
	"strings"
)

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsNotEmpty(s string) bool {
	return len(s) > 0
}

func IsNil(s interface{}) bool {
	return s == nil
}

func IsNotNil(s interface{}) bool {
	return s != nil
}

func RemoveFromSlice(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func RemoveExtraSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func init() {
	fmt.Println("masutils.init()")
}
