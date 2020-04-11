package queries

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Cause",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)