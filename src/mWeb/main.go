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

func main() {
  url := "http://ya.ru"

  pageLength, err := getPage(url)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s length is %d\n", url, pageLength)
}
