package clog

import (
	"fmt"
)

func formatTime(value int) string {
	var result string

	if value < 10 {
		result = fmt.Sprintf("0%d", value)
	} else {
		result = fmt.Sprint(value)
	}

	return result
}
