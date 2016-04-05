package main

import (
	"fmt"
	"os"
)

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 30

	days, ok := dayMonths["January"]
	if !ok {
		fmt.Printf("Can't get number of days\n")
		os.Exit(1)
	}
	fmt.Printf("Days in February: %d\n", days)

}
