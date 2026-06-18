package main

import (
	"fmt"
	"os"
)

func greeting() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go DSR")
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Printf("Hello, %s, Welcome to backend engineering.\n", name)
}
