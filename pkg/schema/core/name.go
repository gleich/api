package core

import "github.com/graphql-go/graphql"

var NameType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "My name",
		Name:        "Name",
		Fields: graphql.Fields{
			"last": &graphql.Field{
				Description: "Last name",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Gleich", nil
				},
			},
			"first": &graphql.Field{
				Description: "First name",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Matt", nil
				},
			},
			"full": &graphql.Field{
				Description: "Full name",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Matt Gleich", nil
				},
			},
		},
	},
)
