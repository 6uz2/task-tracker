package main

import (
	"flag"
	"os"
	"task-tracker/task-tracker/command"
)

func main() {
	flag.Parse()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
