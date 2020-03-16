package main

import (
	"log"
	"os"
	"path/filepath"

	hw "github.com/AirArto/hw-7"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatalf("Error, need more args")
	}

	dir := args[1]
	cmd := args[2:]
	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("Env directory error: %v", err)
	}

	env, err := hw.ReadDir(absDir)
	if err != nil {
		log.Fatalf("Env directory error: %v", err)
	}
	statusCode := hw.RunCmd(cmd, env)

	os.Exit(statusCode)
}
