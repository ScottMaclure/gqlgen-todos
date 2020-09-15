package main

import (
	"gofoo/gqlgen-todos/graph"
	"gofoo/gqlgen-todos/graph/generated"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	r.POST("/query", graphqlHandler())                // The API
	r.GET("/playground", playgroundHandler())         // The playground (not for prod!)
	r.StaticFile("favicon.ico", "./docs/favicon.ico") // Sigh
	r.Static("/client", "./docs")                     // The SPA client

	r.Run()
}
