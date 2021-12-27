package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

type query struct{}

func (_ *query) Hello() string { return "Hello" }

func main() {
	s := `
			type Query {
				hello: String!
			}
	`
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
