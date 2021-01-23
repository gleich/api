package util

import "github.com/graphql-go/graphql"

var NoResolve = func(p graphql.ResolveParams) (interface{}, error) {
	return "", nil
}
