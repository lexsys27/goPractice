package main

import (
	"fmt"
)

func makeID(idChan chan int) {
  var id int
  id = 1

  for {
    idChan <- id;
    id += 1
  }
}

func main() {
  ids := make(chan int)

  go makeID(ids)

  fmt.Printf("%d\n", <- ids)
}
