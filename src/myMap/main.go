package main

import (
	"fmt"
)

func main() {
	dayMonths := map[string]int{
	  "Jan": 31,
	  "Feb": 28,
	  "Mar": 30,
  }

  has31 := 0

  delete(dayMonths, "Feb")

  for _, days := range dayMonths {
    if days == 31 {
      has31++
    }
  }

  fmt.Printf("%d months have 31 days\n", has31)

}
