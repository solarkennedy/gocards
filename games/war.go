package games

import (
	"fmt"
	"github.com/solarkennedy/gocards/cardlib"
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

	for len(hand1) > 0 && len(hand2) > 0 {
		round_counter++
		fmt.Printf("Round %d...\n", round_counter)
		hand1, hand2 = doRound(hand1, hand2, 0)
	}
}

func doRound(hand1 []string, hand2 []string, war_depth int) ([]string, []string) {
	card1 := ""
	card2 := ""
	card1, hand1 = cardlib.PopSlice(hand1)
	card2, hand2 = cardlib.PopSlice(hand2)
	cmp := cardCompare(card1, card2)
	fmt.Printf("Player1: %s	Player 2: %s\n", cardlib.ColorizeCard(card1), cardlib.ColorizeCard(card2))
	if cmp == 0 {
		fmt.Println("WAR")
	} else if cmp > 0 {
		fmt.Println("Player 1 has the higher card")
		hand1 = append(hand1, card1)
		hand1 = append(hand1, card2)
	} else if cmp < 0 {
		fmt.Println("Player 2 has the higher card")
		hand2 = append(hand2, card1)
		hand2 = append(hand2, card2)
	}
	return hand1, hand2
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
