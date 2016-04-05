package main

import (
	"fmt"
  "os"
)

func main() {
  f, err := os.Open("text.txt")
  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }
  defer f.Close()

  b := make([]byte, 100)

  n, err := f.Read(b)
  stringVersion := string(b)
  fmt.Printf("%d bytes read from file\n%s\n", n, stringVersion)
}
