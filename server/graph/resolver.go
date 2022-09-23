package graph

import "github.com/Mickey327/graphqlapp/pkg/postgres"

//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct {
	Postgres *postgres.Postgres
}
