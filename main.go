package main

import (
	"github.com/solarkennedy/gocards/cardlib"
)

func main() {
	deck := cardlib.GetDeckOfCards()
	deck.Shuffle()
	deck.Print()
}
