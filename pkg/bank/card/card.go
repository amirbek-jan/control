package card

import (
	"a2/types"
)

func Vvod() types.Card {
	card := types.Card{
		ID:       1,
		PAN:      "5050 xxxx xxxx 9999",
		Currency: "TJS",
		Color:    "Gold",
		Name:     "Amirbek",
		Balance:  20000,
		Active:   false,
	}
	return Withdraw(&card)
}

func Withdraw(card *types.Card) types.Card {
	if !card.Active {
		return *card
	} else {
		card.Active = false
	}
	return *card
}