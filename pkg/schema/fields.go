package schema

import (
	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"name": &graphql.Field{
		Description: "My name",
		Type:        nameType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "", nil
		},
	},
}
