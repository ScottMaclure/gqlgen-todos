package graph

//go:generate go run github.com/99designs/gqlgen
//(So we can run go generate ./...)

import "gofoo/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is where we declare any dependencies for our app like our database, it gets initialized once in server.go when we create the graph.
type Resolver struct {
	todos []*model.Todo
	users []*model.User
}
