package util

import(
	"os"
)

func HandleError(err error) {
	if err != nil {
		os.Exit(1)
		Print(err.Error())
	}
}