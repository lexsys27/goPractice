package main

import (
	"fmt"
)

func emit(c chan string) {
  words := []string{"Quick", "brown", "fox", "jumps"}

  for _, w := range words {
    c <- w
  }

  close(c)
}

func main() {
  wordsChan := make(chan string)

  go emit(wordsChan)

  for word := range wordsChan {
    fmt.Printf("%s\n", word)
  }

  fmt.Printf("\n")
}
