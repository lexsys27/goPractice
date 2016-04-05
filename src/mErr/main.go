package main

import (
	"fmt"
  "os"
  "errors"
)

var (
  emptyError = errors.New("Won't print blanc message")
)

func printer(msg string) error {
  if msg == "" {
    return emptyError
  }

  _, err := fmt.Printf("%s\n", msg)

  return err
}

func main() {
  if err := printer(""); err == emptyError {
    fmt.Printf("You are trying to print empty string\n")
    os.Exit(1)
  }
}
