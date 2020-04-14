package api

import (
	"github.com/imladenov1997/volunt/graph/model"
	"github.com/imladenov1997/volunt/db"
)

type Queries struct{}

func (q *Queries) GetExchange(id *string) (*model.Exchange, error) {
	database := db.DB{}
	exchange, exchangeErr := database.GetExchange(id)

	return exchange, exchangeErr
}