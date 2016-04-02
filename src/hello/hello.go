package main

import (
	"fmt"
	"os"
)

func main() {
  _, err := fmt.Printf("Hello, world!\n")

  if err != nil {
		os.Exit(1)
	}
}
