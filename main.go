package main

import (
	"fmt"

	hw "github.com/AirArto/hw-2"
)

func init() {
	//
}

func main() {
	PackedArr := []string{"a4bc2d5e", "abcd", "45", "qwe\\4\\5", "qwe\\45", "qwe\\\\5", "\\\\54", ""}

	fmt.Println("Requested data:")
	for _, str := range PackedArr {
		UnpackedStr, err := hw.Do(str)
		fmt.Println("\tPacked string: ", str)
		if err == nil {
			fmt.Println("\tUnpacked string: ", UnpackedStr)
		} else {
			fmt.Println("\tError: ", err)
		}
	}
}
