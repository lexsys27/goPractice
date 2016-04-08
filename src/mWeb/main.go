package main

import (
	"fmt"
  "net/http"
  "io/ioutil"
)

func getPage(url string) (int, error) {
  resp, err := http.Get(url)
  if err != nil {
    return 0, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return 0, err
  }

  return len(body), nil
}

func getter(url string, size chan int) {
  length, err := getPage(url)
  if err == nil {
    size <- length
  }
}

func main() {
  urls := []string{"http://ya.ru", "http://google.com", "http://rbc.ru", "http://vk.com"}

  size := make(chan int)

  for _, url := range urls {
    go getter(url, size)
  }
  for i := 0; i < len(urls); i++ {
    fmt.Printf("Length is %d\n", <-size)
  }
}
