package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schemaFiles []string = []string{
	"geojson.graphql",
	"schema.graphql",
}

const address string = ":8000"

func main() {
	schemaString, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		log.Fatalf("%s", err)
	}

	schema := graphql.MustParseSchema(string(schemaString), &Resolver{}, graphql.UseFieldResolvers())
	http.Handle("/query", &relay.Handler{Schema: schema})

	http.Handle("/", http.FileServer(http.Dir("./views")))

	log.Printf("running on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
