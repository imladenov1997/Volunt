package api

import (
	"github.com/imladenov1997/volunt/graph/model"
	"github.com/imladenov1997/volunt/db"
)

type Queries struct{}

func (q *Queries) GetExchange(id *string) (*model.Exchange, error) {
	var exchange model.Exchange

	database := db.DB{}
	exchangeErr := database.GetExchange(id).Decode(&exchange)

	return &exchange, exchangeErr
}