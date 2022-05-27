package types

// Money представляет собой денежную сумму в минимальных единицах (дирамы, центы, копейки и т.д)
type Money int64

type MinBalance Money

// Код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Номер карты
type PAN string

// Payment представляет информайию о платеже
type Payment struct {
	ID int
	Amount Money
}

// Инфо о карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // использовали Money
	MinBalance MinBalance
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type PaymentSource struct {
	Type string
	Number string
	Balance Money
}