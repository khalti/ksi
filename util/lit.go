package util

import (
	"fmt"
	"strings"
)

// Trims trailing characters from the passed string
func Trim(s string, a ...string)string{
	trimmedOutput := s
	for _, arg:= range a{
		trimmedOutput = strings.Trim(trimmedOutput, arg)
	}
	return trimmedOutput
}

// Prints formatted strings
func Print(format string, a ...any)  {
	fmt.Println(fmt.Sprintf(format, a...))
}