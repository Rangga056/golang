package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main()  {
  rStr := "abcdefg"
  pl("Rune Count :", utf8.RuneCountInString(rStr))
  for i, runeVal := range rStr {
    fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
    //%d = int,
    //%#U = represent a Unicode character in the format U+xxxx, 
    //%c = represent the Unicode character itself. 
  }
}
