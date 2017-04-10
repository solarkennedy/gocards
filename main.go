package main

import (
	"github.com/solarkennedy/gocards/cardlib"
)

func main() {
	cards := cardlib.GetDeckOfCards()
	cards.Print()
}
