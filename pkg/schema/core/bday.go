package core

import (
	"time"

	"github.com/Matt-Gleich/lumber"
	"github.com/graphql-go/graphql"
)

var (
	bday = time.Date(2004, time.April, 8, 4, 13, 0, 0, time.UTC).In(est())
	now  = time.Now().In(est())
)

// Get the Eastern Standard Timezone
func est() *time.Location {
	loc, err := time.LoadLocation("EST")
	lumber.Error(err, "Failed to get EST timezone for birthday date")
	return loc
}

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
					return time.Date(now.Year(), time.April, 8, 0, 0, 0, 1, time.UTC).In(est()), nil
				},
			},
		},
	},
)
