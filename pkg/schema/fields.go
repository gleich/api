package schema

import (
	"github.com/Matt-Gleich/lumber"
	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"name": &graphql.Field{
		Description: "My name",
		Type:        nameType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			lumber.Info("Responding to name field")
			return "", nil
		},
	},
}
