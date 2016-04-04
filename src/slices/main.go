package main

import (
	"fmt"
)

func printer(words []string) {
  for _, word := range words {
    fmt.Printf("%s", word)
  }

  fmt.Printf("\n")
}

func main() {
  words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
  printer(words[:2])
}
