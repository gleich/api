package core

import (
	"time"

	"github.com/graphql-go/graphql"
)

var (
	bday = time.Date(2004, time.April, 9, 5, 14, 0, 0, time.Local)
	now  = time.Now().Local()
)

var BdayType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "My birthday",
		Name:        "Birthday",
		Fields: graphql.Fields{
			"date": &graphql.Field{
				Description: "The date of my birthday",
				Type:        graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return bday, nil
				},
			},
			"age": &graphql.Field{
				Description: "My age in years",
				Type:        graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return time.Since(bday).Hours() / 8760, nil
				},
			},
			"isBirthday": &graphql.Field{
				Description: "My age in hours",
				Type:        graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return now.Day() == 8 && now.Month() == time.April, nil
				},
			},
			"nextBirthday": &graphql.Field{
				Description: "The date of my next birthday",
				Type:        graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return time.Date(now.Year(), time.April, 9, 0, 0, 0, 0, time.Local), nil
				},
			},
		},
	},
)
