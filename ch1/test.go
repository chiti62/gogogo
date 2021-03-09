package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flagSet := flag.NewFlagSet("name", flag.ExitOnError)
	// flag.StringVar(&name, "n", "Go Hello World", "info")
	flagSet.Parse(flag.Args())

	log.Printf("name: %v", name)
}
