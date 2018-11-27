# graphql-go-workshop

[![Go Report Card](https://goreportcard.com/badge/github.com/smoya/graphql-go-workshop?style=flat-square)](https://goreportcard.com/report/github.com/smoya/graphql-go-workshop)

Source code of a simple GraphQL API that queries [Meetup.com API](https://secure.meetup.com/meetup_api) retrieving events among their rsvps. 

Used as part of a live coding talk at [GolangBCN November's 2018 Meetup](https://www.meetup.com/Golang-Barcelona/events/256537826/).

## Installation

### Docker

```bash
docker run -p 8080:8080 -e WORKSHOP_MEETUPAPIKEY=<YOUR_MEETUP_APIKEY> smoya/graphql-go-workshop:latest
```

### Source

This project requires Golang v1.11 or above since it relies in *modules*.

In order to build the app from source, make sure you clone this repo.
Then run `make build`. Alternatively you can run:

```bash
GO111MODULE=on go build -o bin/graphql-go-workshop
```

Then just run `WORKSHOP_MEETUPAPIKEY=<YOUR_MEETUP_APIKEY> bin/graphql-go-workshop`

## Usage

Open your browser and navigate to [http://localhost:8080](http://localhost:8080). 
You can use the GraphQL Playground in order to make GraphQL Queries against our API. 

![GraphQL Playground](https://user-images.githubusercontent.com/1083296/49106315-25959e00-f283-11e8-98a5-ee9ba7016cf4.jpg)

Alternatively you can make GraphQL queries with any other client using [http://localhost:8080/query](http://localhost:8080/query).

 

