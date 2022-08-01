package helper

import(
	"os"
)

func HandleError(err error) {
	if err != nil {
		os.Exit(1)
		print(err.Error())
	}
}