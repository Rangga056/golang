package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	sV1 := "A Word"
	// []bytes
	replacer := strings.NewReplacer("A", "Another") // replace a word/ letter

	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length :", len(sV2))                                   // see the length
	pl("Contains Another :", strings.Contains(sV2, "Another")) //check if it contains the value
	pl("o index :", strings.Index(sV2, "o"))                   //see the index of the letter o
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))        // -1 replace all occurance o with 0
	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3) // get rid of the whitespace
	pl("Split :", strings.Split("a-b-c-d-e", "-"))
	pl("Lowercase :", strings.ToLower(sV3))
	pl("Upperrcase :", strings.ToUpper(sV2))
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))

}
