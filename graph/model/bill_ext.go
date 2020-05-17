package model

import "github.com/google/uuid"

func NewBill(billCurrency *string, value *float64) *Bill {
	billID := uuid.New().String()
	currency := Currency{Name: *billCurrency}
	//people := make(map[string]interface{})
	bill := Bill{ID: billID, Currency: &currency, Value: value}

	return &bill
}

func (bill *Bill) getValue() *float64 {
	return bill.Value
}

func (bill *Bill) ChangeBillCurrency(currency *string) {
	bill.Currency = &Currency{Name: *currency}
}

