package model

func NewTotalBill(billCurrency *string, value *float64) *TotalBill {
	currency := Currency{Name: *billCurrency}
	people := make([]*Person, 0)
	totalBill := TotalBill{Currency: &currency, Value: value, People: people}

	return &totalBill
}

func NewForeignBill(billCurrency *string, value *float64) *ForeignBill {
	var currencyName string
	currencyName = ""

	if (billCurrency != nil) {
		currencyName = *billCurrency
	}

	currency := Currency{Name: currencyName}
	foreignBill := ForeignBill{Currency: &currency, Value: value}

	return &foreignBill
}

func NewExchangeBill(totalBill *TotalBill, foreignBill *ForeignBill) *Exchange {
	var exchangeRate float64
	exchangeRate = 0.0

	if (foreignBill.Value != nil) {
		exchangeRate = *totalBill.Value / *foreignBill.Value
	}

	exc := Exchange{ExchangeFromBill: totalBill, ExchangeToBill: foreignBill, ExchangeRate: &exchangeRate}

	return &exc
}