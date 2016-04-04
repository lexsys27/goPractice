package main

import (
	"fmt"
)

func printer(words [4]string) {
  for _, word := range words {
    fmt.Printf("%s", word)
  }

  fmt.Printf("\n")
}

func main() {
  words := [4]string{"the", "quick", "brown", "fox"}
  printer(words)
}
