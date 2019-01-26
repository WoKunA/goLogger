package main

import (
	"WoKunA/goLogger/logger"
	"flag"
	"fmt"
	"os"
)

type Person struct {
	Age int
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "c", fmt.Sprintf("%s/logger.yml", os.Getenv("PWD")), "logger config file path")
	flag.Parse()

	person := &Person{
		Age: 10,
	}
	logger.DefaultLog.Log(person)
}
