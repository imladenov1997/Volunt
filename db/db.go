package db

import (
	"github.com/imladenov1997/volunt/graph/model" // temporary until the DB layer gets implemented
	"errors"
)

var GlobalExchange *model.Exchange

type DB struct {}

func (db DB) CreateExchange(exchange *model.Exchange) error {
	GlobalExchange = exchange
	return nil
}

func (db DB) GetExchange(ID *string) (*model.Exchange, error) {
	// TODO - implement 

	if GlobalExchange == nil {
		return nil, errors.New("No such exchange")
	} else if *GlobalExchange.ID != *ID {
		return nil, errors.New("No such exchange")
	}

	return GlobalExchange, nil
}