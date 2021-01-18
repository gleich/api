package schema

import (
	"github.com/Matt-Gleich/lumber"
	"github.com/graphql-go/graphql"
)

// Initialize the schema
func Init() graphql.Schema {
	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: Fields,
	}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	lumber.Fatal(err, "Failed to create new schema")
	return schema
}
