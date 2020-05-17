package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

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

func (exchange *Exchange) UpdatePersonalBill(personID *string, fromValue *float64) error {
	if personID == nil {
		return errors.New("Person ID error")
	}

	exchangePairInterface, exists := exchange.People[*personID]
	fmt.Println(&exchangePairInterface)
	fmt.Println(exchangePairInterface)
	exchangePair := exchangePairInterface.(ExchangePair)
	fmt.Println(&exchangePair)
	fmt.Println(exchangePair)

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

	exchange.People[*personID] = exchangePair

	return nil
}

func (exchange *Exchange) GetPersonalBill(personID *string) (*ExchangePair, error) {
	if personID == nil {
		return nil, errors.New("personID is nil")
	}

	exchangePairInterface, exists := exchange.People[*personID]


	if !exists {
		return nil, errors.New("Person not found in this exchange")
	}

	exchangePair := exchangePairInterface.(ExchangePair)

	return &exchangePair, nil

}

func (exchange *Exchange) UpdateExchangeCurrency(currency *string, value *float64) error {
	exchange.FromBill.ChangeBillCurrency(currency)

	if value == nil {
		return nil
	}

	err := exchange.UpdateExchangeRate(value)

	return err
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

// AWFUL WORKAROUND
type MongoExchange struct {
	ID           *string                 `json:"ID"`
	FromBill     *Bill                   `json:"fromBill"`
	ToBill       *Bill                   `json:"toBill"`
	ExchangeRate *float64                `json:"exchangeRate"`
	People       map[string]ExchangePair `json:"people"`
	PersonOnly   *ExchangePair           `json:"personOnly"`
}

func (mExchange *MongoExchange) ToGQLExchange() Exchange {
	gqlExchange := Exchange{
		ID:           mExchange.ID,
		FromBill:     mExchange.FromBill,
		ToBill:       mExchange.ToBill,
		ExchangeRate: mExchange.ExchangeRate,
		People:       map[string]interface{}{},
		PersonOnly:   mExchange.PersonOnly,
	}

	for key, value := range mExchange.People {
		gqlExchange.People[key] = value
	}

	return gqlExchange
}
