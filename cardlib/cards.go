package cardlib

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
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

func ColorizeCard(c string) string {
	if stringInSlice(c, Spades) || stringInSlice(c, Clubs) {
		return fmt.Sprintf("%s ", color.BlackString(c))
	} else if stringInSlice(c, Hearts) || stringInSlice(c, Diamonds) {
		return fmt.Sprintf("%s ", color.RedString(c))
	}
	return ""
}

func (d *Deck) Print() {
	for _, c := range d.cards {
		fmt.Print(ColorizeCard(c))
	}
	fmt.Println()
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	dest := make([]string, len(d.cards))
	perm := rand.Perm(len(d.cards))
	for i := 0; i < len(d.cards); i++ {
		dest[i] = d.cards[perm[i]]
	}
	d.cards = dest
}

func (d *Deck) DealIntoHands(hands int, n int) [][]string {
	hand_array := make([][]string, hands)
	fmt.Printf("Dealing into %d hands of %d cards each...\n", hands, n)
	for h := 0; h < hands; h++ {
		hand_array[h] = make([]string, 0)
	}
	for len(d.cards) > 0 {
		for h := 0; h < hands; h++ {
			card := d.Pop()
			hand_array[h] = append(hand_array[h], card)
		}
	}
	return hand_array
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PopSlice(a []string) (string, []string) {
	x := ""
	x, a = a[len(a)-1], a[:len(a)-1]
	return x, a
}

func (d *Deck) Pop() string {
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
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