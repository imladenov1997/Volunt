package api
import (
	"github.com/imladenov1997/volunt/components"
)

/* Define all the function types here */

type createExchangeFn func(totalBillCurrency *string, totalBillValue *float64, toBillCurrency *string, toBillValue *float64) *components.Exchange

/* Define the main struct holding all functions */

type Mutations struct {
	
}

/* Define all functions here */

func CreateExchange(totalBillCurrency *string, totalBillValue *float64, toBillCurrency *string, toBillValue *float64) *components.Exchange {
	totalBill := components.NewTotalBill(*totalBillCurrency, *totalBillValue)
	foreignBill := components.NewForeignBill(*toBillCurrency, *toBillValue)

	exchange := components.NewExchangeBill(totalBill, foreignBill)

	return exchange
}