package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main () {
  sV1 := "A Word"
  // []bytes
  replacer := strings.NewReplacer("A", "Another")

  sV2 := replacer.Replace(sV1)
  pl(sV2)
}
