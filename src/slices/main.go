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
  words := make([]string, 0, 4)
  fmt.Printf("%d %d\n", len(words), cap(words))
  words = append(words, "the")
  words = append(words, "quick")
  words = append(words, "brown")
  words = append(words, "fox")

  newWords := make([]string, 4)
  copy(newWords, words)
  printer(newWords)

  newWords[2] = "blue"
  printer(newWords)
  printer(words)
}
