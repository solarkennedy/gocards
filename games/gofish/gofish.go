package main

import (
	"fmt"
	"github.com/solarkennedy/gocards/cardlib"
)

type GoFishGame struct {
	deck *cardlib.Deck
	players int
	total_books int
	books []int
	hands [][]string
	whos_turn  int
}

func GoFish() {
	deck := cardlib.GetDeckOfCards()
	deck.Shuffle()
	deck.Print()

	var game = new(GoFishGame)
	game.total_books = 0
	game.players = 3
	game.deck = deck
	game.whos_turn = 0

	game.Deal()
	for game.total_books <= 12 {
		game.DoRound()
	}
	fmt.Println(game.CalculateWinner())
}

func (g *GoFishGame) DoRound() {
	fmt.Printf("Player %d's turn.\n", g.whos_turn)
	// Pick a card to ask about
	// TOOD: pick something other than the first card
	current_player = g.whos_turn
	card_to_ask := g.hands[current_player][0]
	// Pick a person to ask
	person_to_ask := (current_player + 1) % g.players

	fmt.Printf("Player %d asks player %d: 'Do you have any %s?'\n", current_player, person_to_ask, CardToPluralString(card_to_ask))
	hand_locations := hand_has_a(g.hands[person_to_ask], cardlib.CardValue(card_to_ask))
	if len(hand_locations) == 0 {
		// If they don't have it, GO FISH!
		fmt.Printf("\tPlayer %d fishes for %s ...\n", current_player, CardToPluralString(card_to_ask))
		if len(g.deck.Cards) == 0 {
			fmt.Println("\tThere are no fish left in the sea, moving play to the next player")
			g.whos_turn = (current_player + 1) % g.players
		} else {
			fish := g.deck.Pop()
			// If we got what we wanted, then we get to go again
			if cardlib.CardValue(fish) == cardlib.CardValue(card_to_ask) {
				fmt.Printf("\tPlayer %d found it! - %s\n", current_player, cardlib.ColorizeCard(fish))
				g.hands[current_player] = append(g.hands[current_player], fish)
			} else {
				// Otherwise we keep the card, and play moves onto the left
				fmt.Println("\tDidn't get it. Passing to the left.")
				g.hands[current_player] = append(g.hands[current_player], fish)
				g.whos_turn = (current_player + 1) % g.players
			}
		}
	} else {
		// If they have it, put it in our hand
		fmt.Printf("\tPlayer %d does have %d %s! Give them to player %d'\n",person_to_ask, len(hand_locations), CardToPluralString(card_to_ask), current_player)
		for _, location := range hand_locations {
			g.hands[current_player] = append(g.hands[curent_player], g.hands[person_to_ask][location])
			g.hands[person_to_ask] = remove(g.hands[person_to_ask], location)
		}
	}
	g.CalculateBooks(current_player)
	g.total_books++
}

func (g *GoFishGame) CalculateBooks(player int) {
	for index, card := range g.hands[player] {
		
	}
}

func (g *GoFishGame) Deal() {
	if g.players <= 4 {
		g.hands = g.deck.DealIntoHands(g.players, 7)
	} else {
		g.hands = g.deck.DealIntoHands(g.players, 4)
	}
}

func hand_has_a(hand []string, card_inquery int) []int {
	locations := []int{}
	for index, card := range hand {
		if cardlib.CardValue(card) == card_inquery {
			locations = append(locations, index)
		}
	}
	return locations
}

func (g *GoFishGame) CalculateWinner() string {
	return "I dont know who the winner is yet"
}

func CardToPluralString(card string) string {
	value := cardlib.CardValue(card)
	if value == 14 {
		return "Aces"
	} else if value == 2 {
		return "Twos"
	} else if value == 3 {
		return "Threes"
	} else if value == 4 {
		return "Fours"
	} else if value == 5 {
		return "Fives"
	} else if value == 6 {
		return "Sixes"
	} else if value == 7 {
		return "Sevens"
	} else if value == 8 {
		return "Eights"
	} else if value == 9 {
		return "Nines"
	} else if value == 10 {
		return "Tens"
	} else if value == 11 {
		return "Jacks"
	} else if value == 12 {
		return "Queens"
	} else if value == 13 {
		return "Kings"
	} else {
		panic(value)
	}
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	GoFish()
}
