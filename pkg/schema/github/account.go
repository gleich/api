package github

import (
	"fmt"

	"github.com/Matt-Gleich/api/pkg/db"
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
					var followers int
					err := db.DB.QueryRow(fmt.Sprintf("SELECT followers FROM %v;", accountTable)).Scan(&followers)
					if err != nil {
						return 0, nil
					}
					return followers, err
				},
			},
			"email": &graphql.Field{
				Description: "Email for my account",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var email string
					err := db.DB.QueryRow(fmt.Sprintf("SELECT email FROM %v;", accountTable)).Scan(&email)
					if err != nil {
						return "", nil
					}
					return email, err
				},
			},
			"username": &graphql.Field{
				Description: "Username for my account",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var username string
					err := db.DB.QueryRow(fmt.Sprintf("SELECT username FROM %v;", accountTable)).Scan(&username)
					if err != nil {
						return "", nil
					}
					return username, err
				},
			},
			"repos": &graphql.Field{
				Description: "Repositories in my account",
				Type:        graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var repos int
					err := db.DB.QueryRow(fmt.Sprintf("SELECT repos FROM %v;", accountTable)).Scan(&repos)
					if err != nil {
						return 0, nil
					}
					return repos, err
				},
			},
			"contributions": &graphql.Field{
				Description: "Number of contributions I've made in the last year",
				Type:        graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var contributions int
					err := db.DB.QueryRow(fmt.Sprintf("SELECT contributions FROM %v;", accountTable)).Scan(&contributions)
					if err != nil {
						return 0, nil
					}
					return contributions, err
				},
			},
			"hireable": &graphql.Field{
				Description: "If I am hireable",
				Type:        graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var hireable bool
					err := db.DB.QueryRow(fmt.Sprintf("SELECT hireable FROM %v;", accountTable)).Scan(&hireable)
					if err != nil {
						return true, nil
					}
					return hireable, err
				},
			},
			"location": &graphql.Field{
				Description: "Where am I!?!?",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var location string
					err := db.DB.QueryRow(fmt.Sprintf("SELECT location FROM %v;", accountTable)).Scan(&location)
					if err != nil {
						return "", nil
					}
					return location, err
				},
			},
			"organizations": &graphql.Field{
				Description: "Number of organizations I am apart of",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var organizations string
					err := db.DB.QueryRow(fmt.Sprintf("SELECT organizations FROM %v;", accountTable)).Scan(&organizations)
					if err != nil {
						return "", nil
					}
					return organizations, err
				},
			},
			"website": &graphql.Field{
				Description: "Link to my website",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var website string
					err := db.DB.QueryRow(fmt.Sprintf("SELECT website FROM %v;", accountTable)).Scan(&website)
					if err != nil {
						return "", nil
					}
					return website, err
				},
			},
			"company": &graphql.Field{
				Description: "The company I am currently apart of.",
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var company string
					err := db.DB.QueryRow(fmt.Sprintf("SELECT company FROM %v;", accountTable)).Scan(&company)
					if err != nil {
						return "", nil
					}
					return company, err
				},
			},
		},
	},
)
