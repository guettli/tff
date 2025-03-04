package tff

import (
	"fmt"
	"os"
)

func PrintMain(path string) error {
	sourceDev, err := GetDeviceFromPath(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return printEvents(sourceDev)
}
