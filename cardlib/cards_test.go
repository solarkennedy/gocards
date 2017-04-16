package cardlib

import (
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestCardValue(t *testing.T) {
	assertEqual(t, CardValue("🂡"), CardValue("🃑"), "Card value didn't match")
}