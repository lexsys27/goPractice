package main

import (
	"fmt"
)

func printer(msgs ...string) {
  defer fmt.Printf("\n")

  for _, msg := range msgs {
    fmt.Printf("%s", msg)
  }
}

func main() {
  printer("Hello", " , ", "world!")
}
