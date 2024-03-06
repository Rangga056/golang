package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	seedSecs := time.Now().Unix()
	rand.New(rand.NewSource(seedSecs)) // replace the deprecated rand.Seed
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a numebr between 0 and 50 :")
		pl("Random Number is :", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Pick a lower number")
		} else if iGuess < randNum {
			pl("Pick a higher number")
		} else {
			pl("You guessed it right!!!!")
			break
		}
	}
}
