package components

type TotalBill struct {
	Currency *Currency
	Value float64
	people *[]Person
}

type ForeignBill struct {
	Currency *Currency
	Value float64
}

type Exchange struct {
	ExchangeFromBill *TotalBill
	ExchangeToBill *ForeignBill
	exchangeRate float64
}

func NewTotalBill(billCurrency string, value float64) *TotalBill {
	currency := Currency{name: billCurrency}
	people := make([]Person, 0)
	totalBill := TotalBill{Currency: &currency, Value: value, people: &people}

	return &totalBill
}

func NewForeignBill(billCurrency string, value float64) *ForeignBill {
	currency := Currency{name: billCurrency}
	foreignBill := ForeignBill{Currency: &currency, Value: value}

	return &foreignBill
}

func NewExchangeBill(totalBill *TotalBill, foreignBill *ForeignBill) *Exchange {
	exchangeRate := totalBill.Value / foreignBill.Value
	exc := Exchange{ExchangeFromBill: totalBill, ExchangeToBill: foreignBill, exchangeRate: exchangeRate}

	return &exc
}