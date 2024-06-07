package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No file provided")
		return
	}

	file, err := os.ReadFile(args[0])

	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	fmt.Println(string(file))
	fmt.Println("Build page")
}
