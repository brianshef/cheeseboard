package app

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const data = "app/cheese.txt"

// GenerateCheese is a placeholder function
func GenerateCheese() {
	cheeses, err := ReadLines(data)
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}

	rand.Seed(time.Now().Unix())
	cheese := fmt.Sprint("You've found a cheese: ", cheeses[rand.Intn(len(cheeses))])

	fmt.Println(cheese)
}
