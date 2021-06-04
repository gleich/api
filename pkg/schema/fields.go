package schema

import (
	"github.com/gleich/api/pkg/schema/core"
	"github.com/gleich/api/pkg/schema/github"
	"github.com/gleich/api/pkg/util"
	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"name": &graphql.Field{
		Description: "Name",
		Type:        core.NameType,
		Resolve:     util.NoResolve,
	},
	"bday": &graphql.Field{
		Description: "Birthday",
		Type:        core.BdayType,
		Resolve:     util.NoResolve,
	},
	"socials": &graphql.Field{
		Description: "A social media account by name",
		Type:        graphql.NewNonNull(core.SocialsType),
		Resolve:     util.NoResolve,
	},
	"github": &graphql.Field{
		Description: "Informaton from my GitHub account",
		Type:        github.AccountType,
		Resolve:     util.NoResolve,
	},
}
