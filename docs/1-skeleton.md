# 1. Skeleton

You are going to create the skeleton of your GraphQL API. A very basic version of our app that will expose a simple GraphQL Schema and a couple of endpoints.

## 1. Start by downloading or importing the `meetup` package.

You could either [download](https://github.com/smoya/graphql-go-workshop/tree/master/pkg/meetup) or import the `github.com/smoya/graphql-go-workshop/pkg/meetup` package. 
This will be useful for calling the Meetup.com API. It also contains the [model](https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/model.go) for Group, Event and RSVP among others.

## 2. Generate your `schema.graphql` file.

The very first step of any GraphQL API is the generation of the Schema.
Please refer to the [GraphQL documentation](https://graphql.org/learn/schema/) in order to learn about the GraphQL type system.

At this point you can take a look to the model located in the [github.com/smoya/graphql-go-workshop/pkg/meetup](https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/model.go) package and start a very basic schema to query data about 1 Meetup.com group. 

```graphql schema
type Query {
    group(name: String!): Group
}

type Group {
    id: Int!
    name: String!
    who: String
}
```

Save your schema file as `schema.graphql` in the root of your project.

## 3. Generate the basic skeleton code.

For simplicity purpose, we chose [gqlgen](https://gqlgen.com/getting-started/) as implementation library for our GraphQL API.
`gqlgen` is the right tool for generating a GraphQL API from scratch in just few minutes, hiding the complexity by auto-generating code, allowing us to be focused just in the domain logic implementation.    

1. Go get the package:

```bash
go get -u github.com/99designs/gqlgen
```

2. Run:

```bash
gqlgen init
```

This command will generate the basic skeleton of your new app. 

Take your time giving a closer look to those files now:
* `gqlgen.yml` — The gqlgen config file, knobs for controlling the generated code.
* `generated.go` — The GraphQL execution runtime, the bulk of the generated code.
* `models_gen.go` — Generated models required to build the graph. We will override these with the ones provided by the `github.com/smoya/graphql-go-workshop/pkg/meetup` package later.
* `resolver.go` — This is where your application code lives. `generated.go` will call into this to get the data the user has requested.
* `server/server.go` — This is a minimal entry point that sets up an http.Handler to the generated GraphQL server.

## 4. Execute your app

Run `go run server/server.go`. This will start the server.

Open your browser and navigate to [http://localhost:8080](http://localhost:8080). 
As you can see, `gqlgen` generated a handler at `/` that servers the GraphQL Playground. You can use it for making GraphQL Queries against our API.

Try making the following query:

```graphql schema
{
  group(name: "Golang-Barcelona") {
    name
  }
}
``` 

At this step you will end having an error when executing your query.
```json
{
  "data": {
    "group": null
  },
  "errors": [
    {
      "message": "internal system error",
      "path": [
        "group"
      ]
    }
  ]
}
```

Congratulations! You have just created your very first GraphQL API in Golang. Well, is not a useful API now but we will be adding value during the next chapters.