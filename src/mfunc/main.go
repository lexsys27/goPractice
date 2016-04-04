package main

import (
	"fmt"
  "os"
)

func printer(msg string) error {
  f, err := os.Create("hello.txt")

  if err != nil {
    return err
  }

  defer f.Close()
  f.Write(msg)

  return err
}

func main() {
  printer("Hello, world!")
}
