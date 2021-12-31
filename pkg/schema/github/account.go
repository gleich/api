package github

import (
	"github.com/gleich/api/pkg/db"
	"github.com/graphql-go/graphql"
)

var AccountType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "My GitHub account",
		Name:        "GitHub",
		Fields: graphql.Fields{
			"followers": &graphql.Field{
				Description: "Followers for my account",
				Type:        graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Followers, nil
				},
			},
			"email": &graphql.Field{
				Description: "Email for my account",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Email, nil
				},
			},
			"username": &graphql.Field{
				Description: "Username for my account",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Username, nil
				},
			},
			"repos": &graphql.Field{
				Description: "Repositories in my account",
				Type:        graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Repos, nil
				},
			},
			"contributions": &graphql.Field{
				Description: "Number of contributions I've made in the last year",
				Type:        graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Contributions, nil
				},
			},
			"hireable": &graphql.Field{
				Description: "If I am hireable",
				Type:        graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Hireable, nil
				},
			},
			"location": &graphql.Field{
				Description: "Where am I!?!?",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Location, nil
				},
			},
			"organizations": &graphql.Field{
				Description: "Number of organizations I am apart of",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Organizations, nil
				},
			},
			"website": &graphql.Field{
				Description: "Link to my website",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Website, nil
				},
			},
			"company": &graphql.Field{
				Description: "The company I am currently apart of.",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var account db.Account
					db.DB.First(&account, 1)
					return account.Company, nil
				},
			},
		},
	},
)
