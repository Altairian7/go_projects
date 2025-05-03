package main

import (
	"log"
	"net/http"
	"simple-graphql-go/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	resolver := &graph.Resolver{}

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	log.Println("Initializing server...")
	log.Println("Server initialized successfully.")
}

func shutdown() {
	log.Println("Shutting down server...")
	log.Println("Server shut down successfully.")
	log.Println("Goodbye!")
}