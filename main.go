package main

import (
	"errors"
	"fmt"
	"log"

	hw "github.com/AirArto/hw-5"
)

func init() {
	//
}

func main() {
	taskList := [...]func() error{
		func() error { return nil },
		func() error { return errors.New("err") },
		func() error { return nil },
		func() error { return errors.New("err") },
		func() error { return nil },
		func() error { return errors.New("err") },
		func() error { return nil },
		func() error { return errors.New("err") },
		func() error { return nil },
		func() error { return errors.New("err") },
		func() error { return nil },
		func() error { return errors.New("err") },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
	}
	tasks := taskList[:]
	err := hw.Run(tasks, 4, 7)
	if err == nil {
		fmt.Println("It's OK")
	} else {
		log.Println("Data displaying error: \n\t", err)
	}
}
