package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	dice := flag.String("d", "d6", "The type of dice to roll. Format: dX where X is an integer Default: d6")
	numRoll := flag.Int("n", 1, "The number of die to roll. Default: 1")
	flag.Parse()

	matched, _ := regexp.Match("d\\d+", []byte(*dice))

	if matched == true {
		diceSides := (*dice)[1:]
		d, err := strconv.Atoi(diceSides)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < *numRoll; i++ {
			roll := rand.Intn(d) + 1
			fmt.Printf("You rolled a %d\n", roll)
		}

	} else {
		log.Fatalf("Improper formt for dice. Format should be dX where X is integer")
	}
}
