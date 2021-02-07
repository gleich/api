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
		},
	},
)

// A blank description
const noDescription = "n/a"

// A social media account
type socialMediaAccount struct {
	Url         string `graphql:"url"`
	Username    string `graphql:"username"`
	Description string `graphql:"description"`
}

var SocialsType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "Social Media accounts",
		Name:        "Socials",
		Fields: graphql.Fields{
			"twitter": &graphql.Field{
				Description: "Home of my crappy tweets/retweets",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://twitter.com/MattGleich",
						Username:    "MattGleich",
						Description: "Home of my crappy tweets/retweets",
					}, nil
				},
			},
			"github": &graphql.Field{
				Description: "All my open-source and personal work",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://github.com/Matt-Gleich",
						Username:    "Matt-Gleich",
						Description: "All of my open-source and personal work",
					}, nil
				},
			},
			"dockerHub": &graphql.Field{
				Description: "My personal docker images",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://hub.docker.com/u/mattgleich",
						Username:    "mattgleich",
						Description: "Personal docker images",
					}, nil
				},
			},
			"linkedin": &graphql.Field{
				Description: noDescription,
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://www.linkedin.com/in/matthew-gleich",
						Username:    "matthew-gleich",
						Description: noDescription,
					}, nil
				},
			},
			"strava": &graphql.Field{
				Description: "Running and cycling",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://www.strava.com/athletes/30124266",
						Username:    "30124266",
						Description: "Running and cycling",
					}, nil
				},
			},
			"wakatime": &graphql.Field{
				Description: "Programming statistics",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://wakatime.com/@Matthew_Gleich",
						Username:    "Matthew_Gleich",
						Description: "Programming statistics",
					}, nil
				},
			},
			"reddit": &graphql.Field{
				Description: noDescription,
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://www.reddit.com/user/MGleich",
						Username:    "MGleich",
						Description: noDescription,
					}, nil
				},
			},
			"productHunt": &graphql.Field{
				Description: "Big projects",
				Type:        socialMediaAccountType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return socialMediaAccount{
						Url:         "https://www.producthunt.com/@mattgleich",
						Username:    "mattgleich",
						Description: "Big projects",
					}, nil
				},
			},
		},
	},
)
