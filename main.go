package main

import (
	"errors"

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
	_ = hw.Run(tasks, 3, 5)
}
