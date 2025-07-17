package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Name       string
	Port       int
	Passphrase string
}

func ParseArgs() Config {
	var name string
	var port int
	var passphrase string

	flag.StringVar(&name, "name", "", "Your display name in the chat")
	flag.IntVar(&port, "port", 9090, "Port to use for TCP chat server")
	flag.StringVar(&passphrase, "passphrase", "", "Passphrase")

	flag.Parse()

	if name == "" || passphrase == "" {
		fmt.Println("Error: --name and --passphrase are required")
		flag.Usage()
		os.Exit(1)
	}

	return Config{
		Name:       name,
		Port:       port,
		Passphrase: passphrase,
	}
}
