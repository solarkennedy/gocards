package main

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
		if len(hand1)+len(hand2) != 52 {
			panic("BUG: somehow a card got lost!")
		}
	}
	fmt.Printf("\nPlayer %d won after %d rounds!\n", winner, round_counter)
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

func CardValue(c string) (v int) {
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

func main() {
	War()
}
