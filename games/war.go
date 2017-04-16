package games

import (
	"fmt"
	"github.com/solarkennedy/gocards/cardlib"
	"strings"
)

func War() {
	fmt.Println("Starting War")
	deck := cardlib.GetDeckOfCards()
	deck.Shuffle()
	deck.Print()

	hand_array := deck.DealIntoHands(2, 26)
	hand1 := hand_array[0]
	fmt.Println(len(hand1))
	hand2 := hand_array[1]
	round_counter := 0
	winner := 0

	for len(hand1) > 0 && len(hand2) > 0 {
		round_counter++
		fmt.Printf("Round %d... (Player 1: %d cards   Player 2: %d cards)\n", round_counter, len(hand1), len(hand2))
		hand1, hand2, winner = doRound(hand1, hand2, 0)
		if len(hand1) + len(hand2) != 52 { panic("BUG: somehow a card got lost!")}
	}

	fmt.Printf("\nPlayer %d won after %d rounds!\n", winner, round_counter)
}

func doRound(hand1 []string, hand2 []string, war_depth int) ([]string, []string, int) {
	card1 := "X"
	card2 := "Y"
	winner := 0
	indent := strings.Repeat("\t", war_depth)

	fmt.Printf("%sHand1: %v\n", indent, hand1)
	fmt.Printf("%sHand2: %v\n", indent, hand2)
	if len(hand1) == 0 {

		fmt.Printf("Player1 has no cards.\n")
		return hand1, hand2, 2

	} else if len(hand2) == 0{
		fmt.Printf("Player2 has no cards.\n")
		return hand1, hand2, 1
	}

	card1, hand1 = cardlib.PopSlice(hand1)
	card2, hand2 = cardlib.PopSlice(hand2)
	cmp := cardCompare(card1, card2)
	cmp_string := ""
	if cmp == 0 {
	  cmp_string = "WAR!"
	} else if cmp > 0 {
		cmp_string = "Player 1 has the higher card"
	} else if cmp < 0 {
		cmp_string = "Player 2 has the higher card"
	}
	fmt.Printf("%sPlayer1: %s	Player 2: %s (%s)\n", indent, cardlib.ColorizeCard(card1), cardlib.ColorizeCard(card2), cmp_string)
	if cmp == 0 {
		card1a := "1A"
		card2a := "2A"
		card1b := "1B"
		card2b := "2B"
		card1c := "1C"
		card2c := "2C"

		if len(hand1) == 0 {
			fmt.Printf("Player1 doesn't have enough cards for war. Hand is %v\n", hand1)
			hand2 = append(hand2, card1)
			hand2 = append(hand2, card2)
			return hand1, hand2, 2
		} else if len(hand2) == 0 {
			fmt.Printf("Player2 doesn't have enough cards for war. Hand is %v\n", hand1)
			hand1 = append(hand1, card1)
			hand1 = append(hand1, card2)
			return hand1, hand2, 1
		}
		card1a, hand1 = cardlib.PopSlice(hand1)
		card2a, hand2 = cardlib.PopSlice(hand2)
		fmt.Printf("%s%s\t%s\n", indent, cardlib.ColorizeCard(card1a), cardlib.ColorizeCard(card2a))

		if len(hand1) == 0 {
			fmt.Printf("Player1 doesn't have enough cards for war. Hand is %v\n", hand1)
			hand2 = append(hand2, card1)
			hand2 = append(hand2, card2)
			hand2 = append(hand2, card1a)
			hand2 = append(hand2, card2a)
			return hand1, hand2, 2

		} else if len(hand2) == 0 {
			fmt.Printf("Player2 doesn't have enough cards for war. Hand is %v\n", hand2)
			hand1 = append(hand1, card1)
			hand1 = append(hand1, card2)
			hand1 = append(hand1, card1a)
			hand1 = append(hand1, card2a)
			return hand1, hand2, 1
		}

		card1b, hand1 = cardlib.PopSlice(hand1)
		card2b, hand2 = cardlib.PopSlice(hand2)
		fmt.Printf("%s%s\t%s\n", indent, cardlib.ColorizeCard(card1b), cardlib.ColorizeCard(card2b))
		if len(hand1) == 0 {
			fmt.Printf("Player1 doesn't have enough cards for war. Hand is %v\n", hand1)
			hand2 = append(hand2, card1)
			hand2 = append(hand2, card2)
			hand2 = append(hand2, card1a)
			hand2 = append(hand2, card2a)
			hand2 = append(hand2, card1b)
			hand2 = append(hand2, card2b)
			return hand1, hand2, 2

		} else if len(hand2) == 0 {
			fmt.Printf("Player2 doesn't have enough cards for war. Hand is %v\n", hand2)
			hand1 = append(hand1, card1)
			hand1 = append(hand1, card2)
			hand1 = append(hand1, card1a)
			hand1 = append(hand1, card2a)
			hand1 = append(hand1, card1b)
			hand1 = append(hand1, card2b)
			return hand1, hand2, 1
		}

		card1c, hand1 = cardlib.PopSlice(hand1)
		card2c, hand2 = cardlib.PopSlice(hand2)
		fmt.Printf("%s%s\t%s\n", indent, cardlib.ColorizeCard(card1c), cardlib.ColorizeCard(card2c))
		hand1, hand2, winner = doRound(hand1, hand2, war_depth+1)
		if winner == 1 {
			hand1 = append(hand1, card1)
			hand1 = append(hand1, card2)
			hand1 = append(hand1, card1a)
			hand1 = append(hand1, card2a)
			hand1 = append(hand1, card1b)
			hand1 = append(hand1, card2b)
			hand1 = append(hand1, card1c)
			hand1 = append(hand1, card2c)
		} else if winner == 2 {
			hand2 = append(hand2, card1)
			hand2 = append(hand2, card2)
			hand2 = append(hand2, card1a)
			hand2 = append(hand2, card2a)
			hand2 = append(hand2, card1b)
			hand2 = append(hand2, card2b)
			hand2 = append(hand2, card1c)
			hand2 = append(hand2, card2c)
		} else {
			panic(winner)
		}
	} else if cmp > 0 {
		hand1 = append(hand1, card1)
		hand1 = append(hand1, card2)
		winner = 1
	} else if cmp < 0 {
		hand2 = append(hand2, card1)
		hand2 = append(hand2, card2)
		winner = 2
	}
	fmt.Printf("%sHand1: %v (%d cards)\n", indent, hand1, len(hand1))
	fmt.Printf("%sHand2: %v (%d cards)\n", indent, hand2, len(hand2))
	return hand1, hand2, winner
}

func cardCompare(card1 string, card2 string) (v int) {
	n1 := warCardValue(card1)
	n2 := warCardValue(card2)
	if n1 == n2 {
		v = 0
	} else if n1 > n2 {
		v = 1
	} else if n2 > n1 {
		v = -1
	}
	return v
}

func warCardValue(c string) (v int) {
	if c == "🂡" || c == "🂱" || c == "🃁" || c == "🃑" {
		v = 14
	} else if c == "🃒" || c == "🃂" || c == "🂲" || c == "🂢" {
		v = 2
	} else if c == "🂣" || c == "🂳" || c == "🃃" || c == "🃓" {
		v = 3
	} else if c == "🃔" || c == "🃄" || c == "🂴" || c == "🂤" {
		v = 4
	} else if c == "🂥" || c == "🂵" || c == "🃅" || c == "🃕" {
		v = 5
	} else if c == "🃖" || c == "🃆" || c == "🂶" || c == "🂦" {
		v = 6
	} else if c == "🂧" || c == "🂷" || c == "🃇" || c == "🃗" {
		v = 7
	} else if c == "🃘" || c == "🃈" || c == "🂸" || c == "🂨" {
		v = 8
	} else if c == "🂩" || c == "🂹" || c == "🃉" || c == "🃙" {
		v = 9
	} else if c == "🂺" || c == "🂪" || c == "🃚" || c == "🃊" {
		v = 10
	} else if c == "🂻" || c == "🂫" || c == "🃋" || c == "🃛" {
		v = 11
	} else if c == "🂽" || c == "🂭" || c == "🃍" || c == "🃝" {
		v = 12
	} else if c == "🂾" || c == "🂮" || c == "🃎" || c == "🃞" {
		v = 13
	}
	return v
}
