package main

import (
	"fmt"
)

func printer(msg string, repeat int){
  for repeat > 0 {
    fmt.Printf("%s\n", msg)
    repeat -= 1
  }
}

func main() {
  printer("Hello, world!", 3)
}
