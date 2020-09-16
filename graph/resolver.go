package graph

//go:generate go run github.com/99designs/gqlgen
//(So we can run go generate ./...)

import (
	"context"
	"fmt"
	"gqlgen-todos/graph/model"
	"os"

	"github.com/jackc/pgx/v4"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is where we declare any dependencies for our app like our database, it gets initialized once in server.go when we create the graph.
type Resolver struct {
	todos []*model.Todo
	users []*model.User
	db    *pgx.Conn
}

// ConnectDB is a test
func (r *Resolver) ConnectDB(ctx context.Context) {
	if r.db != nil {
		fmt.Printf("ConnectDB connection already open\n")
		return
	}

	fmt.Printf("ConnectDB connecting...\n")
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	// defer conn.Close(ctx) // can't do this, keep the connection open then?
	if err != nil {
		panic(err)
	}

	r.db = conn
}
