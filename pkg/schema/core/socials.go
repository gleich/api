package core

import (
	"github.com/graphql-go/graphql"
)

// Type for a social media account
var socialMediaAccountType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "A social media account",
		Name:        "socialMediaAccount",
		Fields: graphql.Fields{
			"url": &graphql.Field{
				Description: "URL to the account",
				Type:        graphql.String,
			},
			"username": &graphql.Field{
				Description: "User for the account",
				Type:        graphql.String,
			},
		},
	},
)

// A social media account
type socialMediaAccount struct {
	Url      string `graphql:"url"`
	Username string `graphql:"username"`
}

var SocialsType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "Social Media accounts",
		Name:        "socials",
		Fields: graphql.Fields{
			"twitter": &graphql.Field{
				Description: "Home of my crappy tweets",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:      "url",
						Username: "Matt-Gleich",
					}, nil
				},
			},
			"github": &graphql.Field{
				Description: "All my open-source and personal work",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:      "url",
						Username: "Matt-Gleich",
					}, nil
				},
			},
		},
	},
)
