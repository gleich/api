package schema

import (
	"github.com/Matt-Gleich/api/pkg/schema/core"
	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"name": &graphql.Field{
		Description: "Name",
		Type:        core.NameType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "", nil
		},
	},
	"bday": &graphql.Field{
		Description: "Birthday",
		Type:        core.BdayType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "", nil
		},
	},
}
