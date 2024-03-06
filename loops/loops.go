package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// for initialization; condition; postStatement {BODY}
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	// you can also make it like a while loop
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}

}
