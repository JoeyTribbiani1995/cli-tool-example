package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
)

func main() {
	dice := flag.String("d", "d6", "The type of dice to roll. Format: dX where X is an integer Default: d6")
	flag.Parse()

	matched, _ := regexp.Match("d\\d+", []byte(*dice))

	if matched == true {
		diceSides := (*dice)[1:]
		d, err := strconv.Atoi(diceSides)
		if err != nil {
			log.Fatal(err)
		}

		roll := rand.Intn(d) + 1

		fmt.Println(roll)
	} else {
		log.Fatalf("Improper formt for dice. Format should be dX where X is integer")
	}
}
