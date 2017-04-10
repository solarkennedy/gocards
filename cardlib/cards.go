package cardlib

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
)

var Spades = []string{`🂡`, `🂢`, `🂣`, `🂤`, `🂥`, `🂦`, `🂧`, `🂨`, `🂩`, `🂪`, `🂫`, `🂭`, `🂮`}
var Hearts = []string{`🂱`, `🂲`, `🂳`, `🂴`, `🂵`, `🂶`, `🂷`, `🂸`, `🂹`, `🂺`, `🂻`, `🂽`, `🂾`}
var Diamonds = []string{`🃁`, `🃂`, `🃃`, `🃄`, `🃅`, `🃆`, `🃇`, `🃈`, `🃉`, `🃊`, `🃋`, `🃍`, `🃎`}
var Clubs = []string{`🃑`, `🃒`, `🃓`, `🃔`, `🃕`, `🃖`, `🃗`, `🃘`, `🃙`, `🃚`, `🃛`, `🃝`, `🃞`}

//const jokers = "🂿🂠 🃟🃏 "
//other := []string{"🃴"}

type Deck struct {
	cards []string
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func GetDeckOfCards() *Deck {
	cards := []string{}
	for i := 0; i <= 12; i++ {
		cards = append(cards, Spades[i])
	}
	for i := 0; i <= 12; i++ {
		cards = append(cards, Hearts[i])
	}
	for i := 0; i <= 12; i++ {
		cards = append(cards, Clubs[i])
	}
	for i := 0; i <= 12; i++ {
		cards = append(cards, Diamonds[i])
	}
	var deck = new(Deck)
	deck.cards = cards
	return deck
}

func PrintCard(c string) {
	if stringInSlice(c, Spades) || stringInSlice(c, Clubs) {
		fmt.Printf("%s ", color.BlackString(c))
	} else if stringInSlice(c, Hearts) || stringInSlice(c, Diamonds) {
		fmt.Printf("%s ", color.RedString(c))
	}
}

func (d *Deck) Print() {
	for _, c := range d.cards {
		PrintCard(c)
	}
	fmt.Println()
}

func (d *Deck) Shuffle() {
	dest := make([]string, len(d.cards))
	perm := rand.Perm(len(d.cards))
	for i := 0; i < len(d.cards); i++ {
		dest[i] = d.cards[perm[i]]
	}
	d.cards = dest
}

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
