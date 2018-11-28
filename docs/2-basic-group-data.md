# 2. Fetching basic data about Meetup.com Groups 

As you may have noticed, the API is almost useless at this point. In this chapter you are going to move from a useless API to a meaningful one.
You will end returning basic Meetup.com Groups data instead of those ugly error messages. 

## 1. Take a look to the models provided by the `github.com/smoya/graphql-go-workshop/pkg/meetup` package

Navigate to [https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/model.go](https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/model.go).
You will find Structs describing the basic models for our API. Let's be focused on the `Group` one.

```go
type Group struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Created           int64   `json:"created"`
	JoinMode          string  `json:"join_mode"`
	Lat               float64 `json:"lat"`
	Lon               float64 `json:"lon"`
	Urlname           string  `json:"urlname"`
	Who               string  `json:"who"`
	LocalizedLocation string  `json:"localized_location"`
	State             string  `json:"state"`
	Country           string  `json:"country"`
	Region            string  `json:"region"`
	Timezone          string  `json:"timezone"`
}
```

## 2. Tell `gqlgen` to reuse the model

In order to avoid maintaining two versions of our models, you can tell `gqlgen` to skip code generation for some of your models and use your own instead.

1. Edit the `gqlgen.yml` file adding the following content:
```yaml
models:
  Group:
    model: github.com/smoya/graphql-go-workshop/pkg/meetup.Group
```

2. Remove the `resolver.go` file. This will let the `gqlgen` to regenerate it again.
3. Run `gqlgen` in order to regenerate the code.

## 3. Fetch Group data from Meetup.com API

At this point we are ready to start using the methods found in [https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/meetup.go][https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/meetup.go] and start fetching data directly from Meetup.com API.

1. Grab your API Key from [Meetup.com](https://secure.meetup.com/meetup_api/key).
2. Generate a Client for calling the Meetup.com API. You can quickly do this by calling the `NewClient` method on the meetup library:
```go
// We could configure our http client with any desired option.
httpc := http.Client{
    Timeout: s.Timeout,
}
c := meetup.NewClient(&httpc, MEETUP_APIKEY_HERE)
```

3. Pass this client to the `Resolver` struct.
```go
// resolver.go
type Resolver struct{
	c *meetup.Client
}

// server.go
Resolver{C: c}
```

4. Let the `QueryResolver.Group` resolver method use that client for fetching a Meetup.com Group.
```go
func (r *queryResolver) Group(ctx context.Context, name string) (*meetup.Group, error) {
	return r.C.Group(name)
}
```

## 4. Play

Run `go run server/server.go`.
Open your browser and navigate to [http://localhost:8080](http://localhost:8080). 

Execute any available query. Example: 
```graphql
{
  group(name: "Golang-Barcelona") {
    name
  }
}
```

Congratulations! You have just created very first GraphQL Resolver.