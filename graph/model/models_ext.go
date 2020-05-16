package model

import (
	"errors"
	"github.com/google/uuid"
)

func NewBill(billCurrency *string, value *float64) *Bill {
	billID := uuid.New().String()
	currency := Currency{Name: *billCurrency}
	//people := make(map[string]interface{})
	bill := Bill{ID: billID, Currency: &currency, Value: value}

	return &bill
}

func NewExchangeBill(totalBill *Bill, foreignBill *Bill) *Exchange {
	var exchangeRate float64
	exchangeRate = 0.0

	if *foreignBill.Value == 0.0 {
		panic("Division By Zero")
	}

	billID := uuid.New().String()
	exchangeRate = *totalBill.Value / *foreignBill.Value


	people := make(map[string]interface{})

	exc := Exchange{ID: &billID, FromBill: totalBill, ToBill: foreignBill, ExchangeRate: &exchangeRate, People: people}

	return &exc
}

func (exchange *Exchange) AddPersonRef(person *Person, exchangePair *ExchangePair) *ExchangePair {
	exchange.People[person.ID] = exchangePair

	return exchangePair
}

func (exchange *Exchange) AddPerson(value *float64) *ExchangePair {
	personID := uuid.New().String()
	exchangePairID := uuid.New().String()

	person := &Person{ID: personID}
	toValue := *value / *exchange.ExchangeRate
	exchangePair := &ExchangePair{ID: &exchangePairID, Owner: person, FromValue: value, ToValue: &toValue}

	exchange.AddPersonRef(person, exchangePair)

	return exchangePair
}

func (exchange *Exchange) GetToBillValue() *float64 {
	if exchange.ToBill == nil {
		return nil
	}

	return exchange.ToBill.getValue()
}

func (exchange *Exchange) GetTotalBill() *Bill {
	return exchange.FromBill
}

func (person *Person) CheckID(personID *string) bool {
	return person.ID == *personID
}

func (bill *Bill) getValue() *float64 {
	return bill.Value
}

func (exchangePair *ExchangePair) UpdateFromValue(value *float64) {
	exchangePair.FromValue = value
}

func (exchangePair *ExchangePair) UpdateToValue(value *float64) {
	exchangePair.ToValue = value
}

func (exchange *Exchange) UpdatePersonalBill(personID *string, fromValue *float64) error {
	if personID == nil {
		return errors.New("Person ID error")
	}

	exchangePairInterface, exists := exchange.People[*personID]
	exchangePair := exchangePairInterface.(ExchangePair)

	if !exists {
		return errors.New("Person not found in this exchange")
	}

	rate := exchange.ExchangeRate

	if rate != nil && *rate == 0.0 {
		return errors.New("Division by Zero")
	}

	toValue := *fromValue / *rate


	exchangePair.UpdateFromValue(fromValue)
	exchangePair.UpdateToValue(&toValue)

	return nil
}

func (exchange *Exchange) GetPersonalBill(personID *string) (*ExchangePair, error) {
	if personID == nil {
		return nil, errors.New("personID is nil")
	}

	exchangePair := exchange.People[*personID].(*ExchangePair)

	if exchangePair == nil {
		return nil, errors.New("No such Exchange Pair")
	}

	return exchangePair, nil

}

func (exchange *Exchange) UpdateExchangeCurrency(currency *string, value *float64) error {
	exchange.FromBill.ChangeBillCurrency(currency)

	if value == nil {
		return nil
	}

	err := exchange.UpdateExchangeRate(value)

	return err
}

func (bill *Bill) ChangeBillCurrency(currency *string) {
	bill.Currency = &Currency{Name: *currency}
}

func (exchange *Exchange) UpdateExchangeRate(fromBillValue *float64) error {
	if fromBillValue == nil { return errors.New("Non-existent value") }

	if exchange.ToBill.Value == nil || *exchange.ToBill.Value == 0 {
		return errors.New("Invalid target value")
	}

	newExchangeRate := *fromBillValue / *exchange.ToBill.Value

	if newExchangeRate == 0.0 {
		return errors.New("Exchange rate cannot be 0")
	}

	exchange.ExchangeRate = &newExchangeRate

	for _, v := range exchange.People {
		pair := v.(ExchangePair)
		newToVal := *pair.FromValue / newExchangeRate
		pair.UpdateToValue(&newToVal)
	}

	return nil

}
