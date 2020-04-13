package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/imladenov1997/volunt/graph/generated"
	"github.com/imladenov1997/volunt/graph/model"
)

func (r *mutationResolver) CreateExchange(ctx context.Context, totalBillCurrency string, totalBillValue float64, toBillCurrency *string, toBillValue *float64) (*model.Exchange, error) {
	exchange := r.mutations.CreateExchange(&totalBillCurrency, &totalBillValue, toBillCurrency, toBillValue)

	return exchange, nil
}

func (r *mutationResolver) AddPerson(ctx context.Context, value float64) (*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddForeignBill(ctx context.Context, currency string, value string) (model.Bill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SampleQuery(ctx context.Context, test *string) (*int, error) {
	result := 5
	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
