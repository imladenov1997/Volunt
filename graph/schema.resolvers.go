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

func (r *mutationResolver) AddPerson(ctx context.Context, exchangeID string, value float64) (*model.Person, error) {
	person, err := r.mutations.AddPerson(&exchangeID, &value)

	return person, err
}

func (r *mutationResolver) UpdatePersonalBill(ctx context.Context, exchangeID string, personID string, value float64) (*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateForeignBill(ctx context.Context, exchangeID string, currency string, value string) (model.Bill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTotalBill(ctx context.Context, exchangeID string, currency string, value string) (model.Bill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ChangeCurrency(ctx context.Context, billID string, currency string, value string) (model.Bill, error) {
	panic(fmt.Errorf("not implemented"))
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) SampleQuery(ctx context.Context, test *string) (*int, error) {
	result := 5
	return &result, nil
}
