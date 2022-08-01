package helper

import(
	"fmt"
)

func Print(format string, a ...any)  {
	fmt.Println(fmt.Sprintf(format, a...))
}