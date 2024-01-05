package main

import (
	"fmt"
	"log"

	"github.com/varadekd/card-game/helper"
)

func init() {
	err := helper.GenerateDefaultDeck()

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println("Starting application.")
}
