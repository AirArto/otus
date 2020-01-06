package main

import (
	"fmt"
	"log"

	hw "github.com/AirArto/hw-2"
)

func init() {
	//
}

func main() {
	data, err := hw.Do()
	if err == nil {
		fmt.Println("Requested data:", data)
	} else {
		log.Println("Data displaying error: \n\t", err)
	}
}
