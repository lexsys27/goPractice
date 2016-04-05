package main

import (
	"fmt"
  "os"
)

func printer(msg string) error {
  if msg == "" {
    return fmt.Errorf("Won't print blanc message")
  }

  _, err := fmt.Printf("%s\n", msg)

  return err
}

func main() {
  if err := printer(""); err != nil {
    fmt.Printf("printer failed: %s\n", err)
    os.Exit(1)
  }
}
