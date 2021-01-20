package schema

import (
	"github.com/Matt-Gleich/api/pkg/schema/core"
	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"name": &graphql.Field{
		Description: "My name",
		Type:        core.NameType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "", nil
		},
	},
	"bday": &graphql.Field{
		Description: "My birthday",
		Type:        core.BdayType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "", nil
		},
	},
}
