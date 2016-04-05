package main

import (
	"fmt"
)

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 30

  for month, days := range dayMonths {
    fmt.Printf("%s has %d days\n", month, days)
  }

}
