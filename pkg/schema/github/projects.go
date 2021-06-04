package github

import (
	"time"

	"github.com/graphql-go/graphql"
)

type projectData struct {
	OwnerAndName string `graphql:"ownerAndName"`
	Name         string `graphql:"name"`
	Owner        string `graphql:"owner"`
	URL          string `graphql:"url"`
	Language     struct {
		Name  string `graphql:"name"`
		Color string `graphql:"color"`
	} `graphql:"language"`
	Description string    `graphql:"description"`
	Updated     time.Time `graphql:"updated"`
	Created     time.Time `graphql:"created"`
	Stars       int       `graphql:"stars"`
}

var projectDataType = graphql.NewObject(
	graphql.ObjectConfig{
		Description: "One of my projects",
		Name:        "Project",
		Fields: graphql.Fields{
			"ownerAndName": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"owner": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"language": &graphql.Field{
				Type: graphql.NewObject(
					graphql.ObjectConfig{
						Name: "Language",
						Fields: graphql.Fields{
							"name": &graphql.Field{
								Type: graphql.String,
							},
							"color": &graphql.Field{
								Type: graphql.String,
							},
						},
					},
				),
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"updated": &graphql.Field{
				Type: graphql.DateTime,
			},
			"created": &graphql.Field{
				Type: graphql.DateTime,
			},
			"stars": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ProjectField = &graphql.Field{
	Type:        projectDataType,
	Description: "One of my projects",
	Args: graphql.FieldConfigArgument{
		"ownerAndName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// ownerAndName, ok := p.Args["ownerAndName"].(string)
		// if ok {
		// 	// Getting projects
		// 	var projects []projectData
		// 	db.DB.Query()
		// }
		return nil, nil
	},
}
