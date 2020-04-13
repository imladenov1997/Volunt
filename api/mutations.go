package api
import (
	// "github.com/imladenov1997/volunt/components"
	"github.com/imladenov1997/volunt/graph/model"
)

/* Define the main struct holding all functions */

type Mutations struct {}

/* Define all functions here */

// func (m *Mutations) CreateExchange(totalBillCurrency *string, totalBillValue *float64, toBillCurrency *string, toBillValue *float64) *components.Exchange {
// 	totalBill := components.NewTotalBill(*totalBillCurrency, *totalBillValue)
// 	foreignBill := components.NewForeignBill(*toBillCurrency, *toBillValue)

// 	exchange := components.NewExchangeBill(totalBill, foreignBill)

// 	return exchange
// }

func (m *Mutations) CreateExchange(totalBillCurrency *string, totalBillValue *float64, toBillCurrency *string, toBillValue *float64) *model.Exchange {
	totalBill := model.NewTotalBill(totalBillCurrency, totalBillValue)
	
	if (toBillCurrency != nil) {

	}

	foreignBill := model.NewForeignBill(toBillCurrency, toBillValue)

	exchange := model.NewExchangeBill(totalBill, foreignBill)

	return exchange
}