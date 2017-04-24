package main

import (
	"fmt"
	"github.com/solarkennedy/gocards/cardlib"
	"math/rand"
)

type GoFishGame struct {
	deck        *cardlib.Deck
	players     int
	total_books int
	books       []int
	hands       [][]string
	whos_turn   int
}

func GoFish() {
	var game = new(GoFishGame)
	game.Initialize()
	game.Deal()
	for game.total_books <= 12 {
		game.DoRound()
	}
	fmt.Println(game.CalculateWinner())
}

func (g *GoFishGame) Initialize() {
	deck := cardlib.GetDeckOfCards()
	deck.Shuffle()
	deck.Print()
	g.total_books = 0
	g.players = 3
	g.deck = deck
	g.whos_turn = 0
	g.books = make([]int, g.players)
	for player := 0; player < g.players; player++ {
		fmt.Println(player)
		g.books[player] = 0
	}

}

func (g *GoFishGame) DoRound() {
	fmt.Printf("Player %d's turn (they have %d cards).\n", g.whos_turn, len(g.hands[g.whos_turn]))
	current_player := g.whos_turn

	if len(g.hands[g.whos_turn]) == 0 {
		if len(g.deck.Cards) > 0 {
			fmt.Println("  No cards. Drawing from the sea...")
			fish := g.deck.Pop()
			g.hands[current_player] = append(g.hands[current_player], fish)
		} else {
			fmt.Println("  No cards and the sea is empty. Turn is forfeit.\n")
			g.whos_turn = (current_player + 1) % g.players
			return
		}
	}

	// Pick a card to ask about
	r := rand.Intn(len(g.hands[current_player]))
	card_to_ask := g.hands[current_player][r]
	// Pick a person to ask
	person_to_ask := g.PickAPerson(current_player)

	fmt.Printf("Player %d asks player %d: 'Do you have any %s?'\n", current_player, person_to_ask, CardToPluralString(card_to_ask))
	hand_locations := hand_has_a(g.hands[person_to_ask], cardlib.CardRank(card_to_ask))
	if len(hand_locations) == 0 {
		// If they don't have it, GO FISH!
		fmt.Printf("\tNope! GO FISH! Player %d fishes for %s ...\n", current_player, CardToPluralString(card_to_ask))
		if len(g.deck.Cards) == 0 {
			fmt.Println("\tThere are no fish left in the sea, moving play to the next player")
			g.whos_turn = (current_player + 1) % g.players
		} else {
			fish := g.deck.Pop()
			// If we got what we wanted, then we get to go again
			if cardlib.CardRank(fish) == cardlib.CardRank(card_to_ask) {
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
		fmt.Printf("\tPlayer %d does have %d %s! Give them to player %d'\n", person_to_ask, len(hand_locations), CardToPluralString(card_to_ask), current_player)
		for i := 0; i < len(hand_locations); i++ {
			ls := hand_has_a(g.hands[person_to_ask], cardlib.CardRank(card_to_ask))
			l := ls[0]
			g.hands[current_player] = append(g.hands[current_player], g.hands[person_to_ask][l])
			g.hands[person_to_ask] = remove(g.hands[person_to_ask], l)
		}

	}
	g.CalculateBooks(current_player)
}

func (g *GoFishGame) PickAPerson(current_player int) int {
	// TODO: try picking the player with the most cards
	// TODO: AI pick previously seen cards?
	r := rand.Intn(g.players)
	for r == current_player && len(g.hands[r]) != 0 {
		r = rand.Intn(g.players)
	}

	return r
}
func (g *GoFishGame) CalculateBooks(player int) {
	for rank := 2; rank <= 14; rank++ {
		locations := hand_has_a(g.hands[player], rank)
		if len(locations) == 4 {
			fmt.Printf("\tPlayer %d has four %s: ", player, CardRanktoPluralString(rank))
			// We have to remove the cards one at a time because the index changes
			for i := 0; i < 4; i++ {
				locations := hand_has_a(g.hands[player], rank)
				location := locations[0]
				fmt.Printf("%s", cardlib.ColorizeCard(g.hands[player][location]))
				g.hands[player] = remove(g.hands[player], location)
			}
			fmt.Println("")
			g.books[player]++
			g.total_books++
		}
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
		if cardlib.CardRank(card) == card_inquery {
			locations = append(locations, index)
		}
	}
	return locations
}

func (g *GoFishGame) CalculateWinner() string {
	if len(g.deck.Cards) != 0 {
		g.deck.Print()
		panic("The game is done but there is still fish in the sea?")
	}
	var winners = make([]int, 0)
	high_score := 0
	fmt.Printf("\n\nCalculating winner...\n")
	for player := 0; player < g.players; player++ {
		fmt.Printf("\tPlayer %d has %d books\n", player, g.books[player])
		if g.books[player] > high_score {
			high_score = g.books[player]
			winners = []int{player}
		} else if g.books[player] == high_score {
			winners = append(winners, player)
		}
	}
	fmt.Println("")
	if len(winners) == 1 {
		return fmt.Sprintf("Player %d is the winner with %d books!", winners[0], g.books[winners[0]])
	} else if len(winners) == 2 {
		return fmt.Sprintf("We have a tie between player %d and %d with %d books!", winners[0], winners[1], g.books[winners[0]])
	} else if len(winners) >= 3 {
		return fmt.Sprintf("We have a %d-way tie with players %v with %d books!", len(winners), winners, g.books[winners[0]])
	} else {
		panic("I don't know who won")
	}
}

func CardToPluralString(card string) string {
	value := cardlib.CardRank(card)
	return CardRanktoPluralString(value)
}

func CardRanktoPluralString(rank int) string {
	if rank == 14 {
		return "Aces"
	} else if rank == 2 {
		return "Twos"
	} else if rank == 3 {
		return "Threes"
	} else if rank == 4 {
		return "Fours"
	} else if rank == 5 {
		return "Fives"
	} else if rank == 6 {
		return "Sixes"
	} else if rank == 7 {
		return "Sevens"
	} else if rank == 8 {
		return "Eights"
	} else if rank == 9 {
		return "Nines"
	} else if rank == 10 {
		return "Tens"
	} else if rank == 11 {
		return "Jacks"
	} else if rank == 12 {
		return "Queens"
	} else if rank == 13 {
		return "Kings"
	} else {
		panic(rank)
	}

}

func main() {
	GoFish()
}
