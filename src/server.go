package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arx-8/try-go-graphql/src/external"
	"github.com/arx-8/try-go-graphql/src/graphql/directive"
	"github.com/arx-8/try-go-graphql/src/graphql/generated"
	"github.com/arx-8/try-go-graphql/src/graphql/resolver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect DB
	db, err := external.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolver.Resolver{
				DB: db,
			},
			Directives: generated.DirectiveRoot{
				Binding: directive.Binding,
			},
		}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
