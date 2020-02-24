package main

import (
	"flag"
	"fmt"

	hw "github.com/AirArto/hw-6"
)

var (
	from   string
	to     string
	limit  int
	offset int
)

func init() {
	flag.StringVar(&from, "from", "", "file to read")
	flag.StringVar(&to, "to", "", "file to write")
	flag.IntVar(&limit, "limit", 0, "limit to copy in bytes. By default limit = fileSize - offset")
	flag.IntVar(&offset, "offset", 0, "offset to copy in bytes.")
}

func main() {
	flag.Parse()
	_, err := hw.Copy(from, to, limit, offset)
	if err != nil {
		fmt.Println(err)
	}
}
