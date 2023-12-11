package utils

import (
	"fmt"
)

func HandleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}
