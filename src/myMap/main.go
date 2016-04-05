package main

import (
	"fmt"
)

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 30

  has31 := 0

  delete(dayMonths, "Feb")

  for _, days := range dayMonths {
    if days == 31 {
      has31++
    }
  }

  fmt.Printf("%d months have 31 days\n", has31)

}
