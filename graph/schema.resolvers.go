package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/imladenov1997/volunt/graph/generated"
	"github.com/imladenov1997/volunt/graph/model"
)

func (r *mutationResolver) CreateExchange(ctx context.Context, totalBillCurrency string, totalBillValue float64, toBillCurrency string, toBillValue float64) (*model.Exchange, error) {
	exchange, err := r.mutations.CreateExchange(&totalBillCurrency, &totalBillValue, &toBillCurrency, &toBillValue)

	return exchange, err
}

func (r *mutationResolver) AddPerson(ctx context.Context, exchangeID string, value float64) (*model.ExchangePair, error) {
	person, err := r.mutations.AddPerson(&exchangeID, &value)

	return person, err
}

func (r *mutationResolver) UpdatePersonalBill(ctx context.Context, exchangeID string, personID string, value float64) (*model.ExchangePair, error) {
	exchangePair, err := r.mutations.UpdatePersonalBill(&exchangeID, &personID, &value)

	return exchangePair, err
}

func (r *mutationResolver) UpdateForeignBill(ctx context.Context, exchangeID string, currency string, value string) (*model.Bill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTotalBill(ctx context.Context, exchangeID string, currency string, value string) (*model.Bill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ChangeCurrency(ctx context.Context, exchangeID string, currency string, value *float64) (*model.Exchange, error) {
	exchangePair, err := r.mutations.UpdateExchangeCurrency(&exchangeID, &currency, value)

	return exchangePair, err
}

func (r *queryResolver) GetExchange(ctx context.Context, id string) (*model.Exchange, error) {
	return r.queries.GetExchange(&id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
