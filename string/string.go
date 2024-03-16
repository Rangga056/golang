package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array: %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)

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

	// %d : Interger
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n",
		"Stuff", 1, 'A', 3.14, true, 1, 1,
	)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.15123123)
	fmt.Printf("%9.f\n", 3.14512)

	sp1 := fmt.Sprintf("%9.f\n", 3.123123123)
	pl(sp1)
}
