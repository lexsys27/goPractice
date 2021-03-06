package main

import (
	"fmt"
  "time"
)

func emit(wordsChannel chan string, done chan bool) {
  defer close(wordsChannel)
  words := []string{"The", "quick", "brown", "fox"}

  i := 0

  t := time.NewTimer(3 * time.Second)

  for {
    select {
      case wordsChannel <- words[i]:
        i += 1
        if i == len(words) {
          i = 0
        }
      case <-done:
        done <- true
        return
      case <-t.C:
        return
    }
  }
}

func main() {
  wordCh := make(chan string)
  doneCh := make(chan bool)

  go emit(wordCh, doneCh)

  for word := range wordCh {
    fmt.Printf("%s ", word)
  }

  fmt.Printf("\n")
}
