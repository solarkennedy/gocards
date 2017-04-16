package main

import (
	"fmt"
	"github.com/solarkennedy/gocards/cardlib"
)

type GoFishGame struct {
	deck *cardlib.Deck
	players int
	books int
}

func GoFish() {
	deck := cardlib.GetDeckOfCards()
	deck.Shuffle()
	deck.Print()

	var game = new(GoFishGame)
	game.books = 0
	game.players = 2
	game.deck = deck

	for game.books <= 12 {
		game.DoRound()
	}
	fmt.Println(game.CalculateWinner())
}

func (g *GoFishGame) DoRound() {
	g.books = 13
}

func (g *GoFishGame) CalculateWinner() string {
	return "I dont know who the winner is yet"
}

func main() {
	GoFish()
}
