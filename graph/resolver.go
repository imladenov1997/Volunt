package graph

import (
	"github.com/imladenov1997/volunt/api"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	mutations api.Mutations
	queries api.Queries
}

