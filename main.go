package main

import (
	"fmt"

	hw "github.com/AirArto/hw-4"
)

func init() {
	//
}

func main() {
	newList := hw.List{}
	slice := newList.ToSlice()
	fmt.Println("Requested data:", slice)
	newList.PushFront("second")
	newList.PushFront("first")
	newList.PushBack("third")
	newList.PushBack("last")
	slice = newList.ToSlice()
	fmt.Println("Requested data:", slice)
}
