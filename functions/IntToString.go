package functions

import (
	"strconv"
)

func IntToString(i int) string {
	result := strconv.Itoa(i)
	return result
}
