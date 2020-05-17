package api

import (
	"errors"
	"github.com/imladenov1997/volunt/db"
	// "github.com/imladenov1997/volunt/components"
	"github.com/imladenov1997/volunt/graph/model"
)

/* Define the main struct holding all functions */

type Mutations struct {}

func (m *Mutations) CreateExchange(totalBillCurrency *string, totalBillValue *float64, toBillCurrency *string, toBillValue *float64) (exchange *model.Exchange,  errMsg error) {
	errMsg = nil
	database := db.DB{}

	defer func() {
		if err := recover(); err != nil {
			errMsg = errors.New(err.(string))
			exchange = nil
		}
	}()
	
	totalBill := model.NewBill(totalBillCurrency, totalBillValue)

	foreignBill := model.NewBill(toBillCurrency, toBillValue)
	
	exchange = model.NewExchangeBill(totalBill, foreignBill)

	database.CreateExchange(exchange)

	return exchange, errMsg
}

func (m *Mutations) AddPerson(exchangeID *string, value *float64) (exchangePair *model.ExchangePair, errMsg error) {
	var exchange model.Exchange

	database := db.DB{}
	err := database.GetExchange(exchangeID).Decode(&exchange)

	if err != nil {
		return nil, err
	}

	toBillValue := exchange.GetToBillValue()

	if toBillValue == nil {
		return nil, errors.New("Target amount missing")
	}

	if *value > *toBillValue {
		return nil, errors.New("Added amount higher than total")
	}

	exchangePair = exchange.AddPerson(value)
	err = database.UpsertPersonToExchange(exchangeID, exchangePair)

	if err != nil {
		return nil, err
	}

	return exchangePair, nil
}

func (m *Mutations) UpdatePersonalBill(exchangeID *string, personID *string, value *float64) (*model.ExchangePair, error) {
	var exchange model.Exchange

	database := db.DB{}
	exchangeErr := database.GetExchange(exchangeID).Decode(&exchange)

	if exchangeErr != nil {
		return nil, exchangeErr
	}

	result := exchange.UpdatePersonalBill(personID, value)

	if result != nil {
		return nil, result
	}

	exchangePair, err := exchange.GetPersonalBill(personID)
	err = database.UpsertPersonToExchange(exchangeID, exchangePair)

	if err != nil {
		return nil, err
	}

	return exchangePair, nil
}

func (m *Mutations) UpdateExchangeCurrency(exchangeID *string, currency *string, value *float64) (*model.Exchange, error) {
	var exchange model.Exchange

	database := db.DB{}
	exchangeErr := database.GetExchange(exchangeID).Decode(&exchange)

	if exchangeErr != nil {
		return nil, exchangeErr
	}

	if currency == nil {
		return nil, errors.New("Currency not provided")
	}

	err := exchange.UpdateExchangeCurrency(currency, value)

	return &exchange, err

}
