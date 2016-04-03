package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumped ovew the lazy dog"

	vowels := 0
	consonants := 0
  zeds := 0

	for _, r := range atoz {
		switch r {
		case 'i', 'e', 'a', 'o', 'u':
			vowels += 1
    case 'z':
      zeds += 1
      fallthrough
		default:
			consonants += 1
		}
	}

	fmt.Printf("String contains %d vowels and %d consonants(including %d zeds)\n",
		vowels,
		consonants,
    zeds)
}
