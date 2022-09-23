package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/Mickey327/graphqlapp/pkg/postgres"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Mickey327/graphqlapp/graph"
	"github.com/Mickey327/graphqlapp/graph/generated"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	database, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal("Couldn't create database")
	}
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Postgres: database}}))

	srv.AddTransport(&transport.Websocket{Upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
