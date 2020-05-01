package api
import (
	// "github.com/imladenov1997/volunt/components"
	"github.com/imladenov1997/volunt/graph/model"
	"github.com/imladenov1997/volunt/db"
	"errors"
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

	person = exchange.AddPerson(value)

	return person, nil
}

func (m *Mutations) UpdatePersonalBill(exchangeID *string, personID *string, value *float64) (*model.Person, error) {
	database := db.DB{}
	_, exchangeErr := database.GetExchange(exchangeID)

	if exchangeErr != nil {
		return nil, exchangeErr
	}

	return nil, nil

}