package types

type Money int64

type PAN string

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Color string

const (
	Silver Color = "Silver"
	Gold   Color = "Gold"
)

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    Color
	Name     string
	Active   bool
}