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
			"description": &graphql.Field{
				Description: "Description of the account",
				Type:        graphql.String,
			},
			"name": &graphql.Field{
				Description: "Name for the website/service",
				Type:        graphql.String,
			},
		},
	},
)

// A blank description
const noDescription = "n/a"

// A social media account
type socialMediaAccount struct {
	URL         string `graphql:"url"`
	Username    string `graphql:"username"`
	Description string `graphql:"description"`
	Name        string `graphql:"name"`
}

var socials = []socialMediaAccount{
	{
		URL:         "https://twitter.com/matt_gleich",
		Username:    "matt_gleich",
		Description: "Home of my crappy tweets/retweets",
		Name:        "twitter",
	},
	{
		URL:         "https://github.com/gleich",
		Username:    "gleich",
		Description: "All of my open-source and personal work",
		Name:        "github",
	},
	{
		URL:         "https://hub.docker.com/u/mattgleich",
		Username:    "mattgleich",
		Description: "Personal docker images",
		Name:        "dockerhub",
	},
	{
		URL:         "https://www.linkedin.com/in/matt-gleich",
		Username:    "matthew-gleich",
		Description: noDescription,
		Name:        "linkedin",
	},
	{
		URL:         "https://www.strava.com/athletes/30124266",
		Username:    "30124266",
		Description: "Running and cycling",
		Name:        "strava",
	},
	{
		URL:         "https://wakatime.com/@Matthew_Gleich",
		Username:    "Matthew_Gleich",
		Description: "Programming statistics",
		Name:        "wakatime",
	},
	{
		URL:         "https://www.reddit.com/user/MGleich",
		Username:    "MGleich",
		Description: noDescription,
		Name:        "reddit",
	},
	{
		URL:         "https://www.producthunt.com/@mattgleich",
		Username:    "mattgleich",
		Description: "Big projects",
		Name:        "producthunt",
	},
}

var SocialsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"account": &graphql.Field{
				Type:        socialMediaAccountType,
				Description: "Get a social media account by name",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, ok := p.Args["name"].(string)
					if ok {
						for _, social := range socials {
							if name == social.Name {
								return social, nil
							}
						}
					}
					return nil, nil
				},
			},
			"accounts": &graphql.Field{
				Type:        graphql.NewList(socialMediaAccountType),
				Description: "Get all social media accounts",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socials, nil
				},
			},
		},
	},
)
