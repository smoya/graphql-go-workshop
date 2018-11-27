package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/smoya/graphql-go-workshop/pkg/meetup"

	"github.com/99designs/gqlgen/handler"
	"github.com/smoya/graphql-go-workshop/internal/workshop"
)

type specification struct {
	Debug        bool
	Port         int           `default:"8080"`
	MeetupAPIKey string        `required:"true"`
	Timeout      time.Duration `default:"2s"`
}

func main() {
	var s specification
	err := envconfig.Process("workshop", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	// We could configure our http client with any desired option.
	httpc := http.Client{
		Timeout: s.Timeout,
	}
	c := meetup.NewClient(&httpc, s.MeetupAPIKey)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(workshop.NewExecutableSchema(workshop.Config{Resolvers: &workshop.Resolver{C: c}})))

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil))
}
