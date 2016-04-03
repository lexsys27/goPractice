package main

import (
  "fmt"
  "os"
)

func main() {
  n, err := fmt.Printf("Hello, world!\n")

  switch {
    case err != nil:
      os.Exit(1)
    case n == 0:
      fmt.Printf("Nothing was printed\n")
    case n != 14:
      fmt.Printf("Wrong number of bytes printed out\n")
    default:
      fmt.Printf("OK!\n")
  }
}
