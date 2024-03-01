package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	// var name type
	var vName string = "Rangga"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4

	pl(vName, v1, v2, v3, v4)

	cV1 := 1.5
	cV2 := int(cV1)
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)

	pl(cV4, err, reflect.TypeOf(cV4))

	pl(cV2)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	} else {
		log.Fatal(err)
	}

	cV9 := fmt.Sprintf("%f", 3.14) // convert float to string
	pl(cV9)

}
