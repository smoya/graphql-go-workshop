# 4. Fetching RSVPs from a Meetup.com Event

At this point, your GraphQL API is able to fetch Meetup.com Groups filtered by their name as their list of events.
You would want to go further and fetch the RSVPs of those events as well.

Meetup.com API does append that data to the payload we fetched for getting Events data. Instead, it provides another endpoint that brings you data about RSVPs.

## 1. Edit your `graphql.schema` file.

Let's add some new types. 
As before, you can take a look to the model located in the [github.com/smoya/graphql-go-workshop/pkg/meetup](https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/model.go) package.

Add the following types:
```graphql schema
type Rsvp {
    created: String!
    updated: String!
    response: String!
    guests: Int
    member: Member!
}

enum RsvpResponse{
  yes
  no
}

type Member {
    id: Int!
    name: String!
    isHost: Boolean!
}
```

Add the following field to the `Event` type:
```graphql schema
rsvp(response: RsvpResponse): [Rsvp]
```

## 2. Edit your `gqlgen.yaml` file.

Before generating the code again, we would need to add few new lines to the `gqlgen.yaml` file.

Please add the following lines under `model`.
```yaml
  Rsvp:
    model: github.com/smoya/graphql-go-workshop/pkg/meetup.RSVP
  Member:
    model: github.com/smoya/graphql-go-workshop/pkg/meetup.Member
```

## 3. Regenerate your code

> Note 
> The `resolver.go` file is generated **only** when the file is not present. 

As the generator will end creating new Resolvers, you would need to choose 1 option from:

1. Moving the `resolver.go` file to a `resolver.go.copy` before executing `gqlgen` for later "copy and paste" the needed code to the generated code.
2. Removing the `resolver.go` file and redo all the previous steps.

## 4. Implement the Resolver methods.

### GroupResolver
`gqlgen` would generate a new method on the `eventResolver` called `Rsvp`. 

<details> 
  <summary>Use the Meetup.com client for fetching data about the RSVP of such event</summary> 
  
```go
func (r *eventResolver) Rsvp(ctx context.Context, obj *meetup.Event, response *RsvpResponse) ([]*meetup.RSVP, error) {
	return r.C.RSVPs(obj.Group.Urlname, obj.ID, (*string)(response))
}
```
</details>

### RSVPResolver
`gqlgen` would generate a new resolver called `rsvpResolver`.

<details> 
  <summary>You would need to implement some methods:</summary> 

```go
func (r *rsvpResolver) Created(ctx context.Context, obj *meetup.RSVP) (string, error) {
	return time.Unix(obj.Created/1000, 0).Format(time.RFC822), nil
}
func (r *rsvpResolver) Updated(ctx context.Context, obj *meetup.RSVP) (string, error) {
	return time.Unix(obj.Updated/1000, 0).Format(time.RFC822), nil
}
func (r *rsvpResolver) Response(ctx context.Context, obj *meetup.RSVP) (RsvpResponse, error) {
	return RsvpResponse(obj.Response), nil
}
```
</details>

## 5. Play

Run `go run server/server.go`.
Open your browser and navigate to [http://localhost:8080](http://localhost:8080). 

Execute the following query: 
```graphql
{
  group(name: "Golang-Barcelona") {
    name
    events(status: upcoming) {
      name
      rsvp(response: no) {
        member {
          name
        }
        response
        created
        updated
      }
    }
  }
}
```