package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	seedSecs := time.Now().Unix()
	rand.New(rand.NewSource(seedSecs)) // replace the deprecated rand.Seed
	randNum := rand.Intn(50) + 1
	pl("Random values =", randNum)
}
