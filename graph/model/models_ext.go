package model

import (
	"github.com/google/uuid"
	"errors"
)

func NewTotalBill(billCurrency *string, value *float64) *TotalBill {
	billID := uuid.New().String()
	currency := Currency{Name: *billCurrency}
	people := make(map[string]interface{})
	totalBill := TotalBill{ID: billID, Currency: &currency, Value: value, People: people}

	return &totalBill
}

func NewForeignBill(billCurrency *string, value *float64) *ForeignBill {
	var currencyName string
	currencyName = ""

	if (billCurrency != nil) {
		currencyName = *billCurrency
	}

	billID := uuid.New().String()
	currency := Currency{Name: currencyName}
	foreignBill := ForeignBill{ID: billID, Currency: &currency, Value: value}

	return &foreignBill
}

func NewExchangeBill(totalBill *TotalBill, foreignBill *ForeignBill) *Exchange {
	var exchangeRate float64
	exchangeRate = 0.0

	if (*foreignBill.Value == 0.0) {
		panic("Division By Zero")
	}
	
	billID := uuid.New().String()
	exchangeRate = *totalBill.Value / *foreignBill.Value

	exc := Exchange{ID: &billID, ExchangeFromBill: totalBill, ExchangeToBill: foreignBill, ExchangeRate: &exchangeRate}

	return &exc
}

func (totalBill *TotalBill) AddPersonRef(person *Person) *Person {
	totalBill.People[person.ID] = person

	return person
}

func (totalBill *TotalBill) AddPerson(value *float64) *Person {
	personID := uuid.New().String()
	billID := uuid.New().String()

	personalBill := PersonalBill{ID: billID, Currency: totalBill.Currency, Value: value}
	person := &Person{ID: personID, Bill: &personalBill}

	totalBill.AddPersonRef(person)

	return person
}

func (exchange *Exchange) GetTotalBill() *TotalBill {
	return exchange.ExchangeFromBill
}

func (bill *TotalBill) HasPerson(personID *string) (bool, error) {
	if (personID == nil) {
		return false, errors.New("No valid PersonID")
	}

	_, hasVal := bill.People[*personID]
	
	return hasVal, nil
}

func (person *Person) CheckID(personID *string) bool {
	return person.ID == *personID
}