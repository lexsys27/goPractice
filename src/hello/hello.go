package main

import (
	"fmt"
	"os"
)

func main() {
	if n, err := fmt.Printf("Hello, world!\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed number %d of bytes\n", n)
	}
}
