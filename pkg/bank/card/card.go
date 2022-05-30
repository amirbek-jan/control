package card

import (
	"a2/types"
)

func Vvod() types.Card {
	card := types.Card{
		ID:       1,
		PAN:      "5050 xxxx xxxx 6969",
		Currency: "TJS",
		Color:    "Gold",
		Name:     "Boozoorg",
		Balance:  20000,
		Active:   false,
	}
	return Withdraw(&card)
}

func Withdraw(card *types.Card) types.Card {
	if !card.Active {
		card.Active = true
	} else {
		card.Active = false
	}
	return *card
}