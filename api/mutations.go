package api

import (
	"errors"
	"github.com/imladenov1997/volunt/db"
	// "github.com/imladenov1997/volunt/components"
	"github.com/imladenov1997/volunt/graph/model"
	// "fmt"
)

/* Define the main struct holding all functions */

type Mutations struct {}

func (m *Mutations) CreateExchange(totalBillCurrency *string, totalBillValue *float64, toBillCurrency *string, toBillValue *float64) (exchange *model.Exchange,  errMsg error) {
	errMsg = nil
	
	defer func() {
		if err := recover(); err != nil {
			errMsg = errors.New(err.(string))
			exchange = nil
		}
	}()
	
	totalBill := model.NewBill(totalBillCurrency, totalBillValue)

	foreignBill := model.NewBill(toBillCurrency, toBillValue)
	
	exchange = model.NewExchangeBill(totalBill, foreignBill)

	db.GlobalExchange = exchange

	return exchange, errMsg
}

func (m *Mutations) AddPerson(exchangeID *string, value *float64) (person *model.ExchangePair, errMsg error) {
	database := db.DB{}
	exchange, err := database.GetExchange(exchangeID)

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

	person = exchange.AddPerson(value)

	return person, nil
}

func (m *Mutations) UpdatePersonalBill(exchangeID *string, personID *string, value *float64) (*model.ExchangePair, error) {
	database := db.DB{}
	exchange, exchangeErr := database.GetExchange(exchangeID)

	if exchangeErr != nil {
		return nil, exchangeErr
	}

	result := exchange.UpdatePersonalBill(personID, value)

	return nil, result
}

func (m *Mutations) UpdateExchangeCurrency(exchangeID *string, currency *string, value *float64) (*model.Exchange, error) {
	database := db.DB{}
	exchange, exchangeErr := database.GetExchange(exchangeID)

	if exchangeErr != nil {
		return nil, exchangeErr
	}

	if currency == nil {
		return nil, errors.New("Currency not provided")
	}

	err := exchange.UpdateExchangeCurrency(currency, value)

	return exchange, err

}
