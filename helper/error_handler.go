package helper

import(
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}