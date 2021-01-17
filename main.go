package main

import (
	"net/http"

	"github.com/Matt-Gleich/lumber"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	fields := graphql.Fields{
		"last_name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				lumber.Info("Responding to last_name field")
				return "Gleich", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	lumber.Fatal(err, "Failed to create new schema")

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   false,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err = http.ListenAndServe(":8080", nil)
	lumber.Fatal(err, "Failed to listen and serve for requests")
}
