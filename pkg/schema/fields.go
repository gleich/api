package schema

import (
	"github.com/Matt-Gleich/lumber"
	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"last_name": &graphql.Field{
		Description: "My last name",
		Type:        graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			lumber.Info("Responding to the last_name field")
			return "Gleich", nil
		},
	},
	"first_name": &graphql.Field{
		Description: "My first name",
		Type:        graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			lumber.Info("Responding to the first_name field")
			return "Matt", nil
		},
	},
	"full_name": &graphql.Field{
		Description: "My full name",
		Type:        graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			lumber.Info("Responding to the full_name field")
			return "Matt Gleich", nil
		},
	},
}
