package main

import (
	"fmt"
  "os"
)

func printer(msg string) error{
  _, err := fmt.Printf("%s\n", msg)

  return err
}

func main() {
  printerError := printer("Hello, world!")

  if printerError != nil {
    os.Exit(1)
  }
}
