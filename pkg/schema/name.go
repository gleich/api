package schema

import "github.com/graphql-go/graphql"

var nameType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "My name",
		Name:        "Name",
		Fields: graphql.Fields{
			"last": &graphql.Field{
				Description: "My last name",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Gleich", nil
				},
			},
			"first": &graphql.Field{
				Description: "My first name",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Matt", nil
				},
			},
			"full": &graphql.Field{
				Description: "My full name",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Matt Gleich", nil
				},
			},
		},
	},
)
