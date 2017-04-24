package main

import (
	"fmt"
	"github.com/solarkennedy/gocards/cardlib"
)

type CrazyEightsGame struct {
	deck    *cardlib.Deck
	winner  int
	players int
	hands       [][]string
	current_player int
}

func (g *CrazyEightsGame) Initialize() {
	g.winner = -1
	g.players = 4
	deck := cardlib.GetDeckOfCards()
	deck.Shuffle()
	g.deck = deck
	g.current_player = 0
}

func (g *CrazyEightsGame) Deal() {
	if g.players == 2 {
		g.hands = g.deck.DealIntoHands(g.players, 7)
	} else {
		g.hands = g.deck.DealIntoHands(g.players, 8)
	}
}

func (g *CrazyEightsGame) DoRound() {
	top_card := cardlib.ColorizeCard(g.deck.Cards[0])
	fmt.Printf("Top card: %s\n", top_card)
	// If they can play a valid card
	if HasValidCard(g.hands[g.current_player], top_card) != -1 {
		// Play all the cards you can onto the deck
		for HasValidCard(g.hands[g.current_player], top_card) != -1  {
			// Discard ...
			discard_index := HasValidCard(g.hands[g.current_player], top_card)
			discard := g.hands[g.current_player][discard_index]
			fmt.Printf("Player %d discards a %s\n", g.current_player, cardlib.ColorizeCard(discard))
			g.hands[g.current_player] = cardlib.Remove(g.hands[g.current_player], discard_index)
		}
	} else {
		// Otherwise you need to draw until you can
		for HasValidCard(g.hands[g.current_player], top_card) != -1 {
			// Draw a card...
			g.hands[g.current_player] = append(g.hands[g.current_player], g.deck.Pop())
		}
		// Now play that card right away
	}
	if len(g.hands[g.current_player]) == 0 {
		g.winner = g.current_player
	}
}

func HasValidCard(hand []string, showing_card string) int {
	for index,card := range hand {
		if cardlib.CardRank(card) == cardlib.CardRank(showing_card) || cardlib.CardSuit(card) == cardlib.CardSuit(showing_card) {
			return index
		}
	}
	return -1
}

func main() {
	var game = new(CrazyEightsGame)
	game.Initialize()
	game.Deal()
	for game.winner != -1 {
		game.DoRound()
	}
	fmt.Printf("The winner is player %d!\n", game.winner)
}
