package model

import "github.com/google/uuid"

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

//func (exchange *Exchange) GetTotalBill() *TotalBill {
//	return exchange.ExchangeFromBill
//}
//
//func (bill *TotalBill) HasPerson(personID *string) (bool, error) {
//	if (personID == nil) {
//		return false, errors.New("No valid PersonID")
//	}
//
//	_, hasVal := bill.People[*personID]
//
//	return hasVal, nil
//}

func (person *Person) CheckID(personID *string) bool {
	return person.ID == *personID
}