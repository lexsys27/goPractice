package main

import (
	"fmt"
)

func printer(msg string) (string, error) {
  msg += "\n"

  _, err := fmt.Printf("%s", msg)

  return msg, err
}

func main() {
  appendedString, printerError := printer("Hello, world!")

  if printerError == nil {
    fmt.Printf("% x\n", appendedString)
  }
}
