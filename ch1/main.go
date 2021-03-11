package main

import (
	"log"

	"github.com/chiti62/gogogo/ch1/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
