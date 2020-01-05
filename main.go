package main

import (
	"fmt"
	"log"

	print_time "github.com/AirArto/hw-1"
)

func init() {
	//
}

func main() {
	data, err := print_time.PrintDt()
	if err == nil {
		fmt.Printf("Requested data: %s\n", data)
	} else {
		log.Println("Data displaying error: \n\t", err)
	}
}
