# 3. Fetching Events from a Meetup.com Group

At this point, your GraphQL API is able to fetch Meetup.com Groups filtered by their name.
Even though this is somehow useful, you will want to fetch the events from those Groups as well.

Meetup.com API does append that data to the payload we fetched for getting Groups data. Instead, it provides another endpoint that brings you data about events.

## 1. Edit your `graphql.schema` file.

Let's add some new types. This time we are adding some more lines compared to the 1st chapter. 
As before, you can take a look to the model located in the [github.com/smoya/graphql-go-workshop/pkg/meetup](https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/model.go) package.

Add the following types:
```graphql schema
type Event {
  id: String!
  name: String!
  description: String
  created: String!
  duration: Int
  rsvpLimit: Int
  status: EventStatus!
  time: String!
  waitlistCount: Int
  yesRSVPCount: Int
  venue: Venue!
  link: String
}

enum EventStatus {
    cancelled
    draft
    past,
    proposed
    suggested
    upcoming
}

type Venue {
    id: Int!
    name: String!
    address: String
    city: String
    country: String
}
```

Add the following field to the `Group` type:
```graphql schema
events(status: EventStatus): [Event]
```

## 2. Edit your `gqlgen.yaml` file.

Before generating the code again, we would need to add few new lines to the `gqlgen.yaml` file.

Please add the following lines under `model`.
```yaml
  Event:
    model: github.com/smoya/graphql-go-workshop/pkg/meetup.Event
  Venue:
    model: github.com/smoya/graphql-go-workshop/pkg/meetup.Venue
```

## 3. Regenerate your code

> Note 
> The `resolver.go` file is generated **only** when the file is not present. 

As the generator will end creating new Resolvers, you would need to choose 1 option from:

1. Moving the `resolver.go` file to a `resolver.go.copy` before executing `gqlgen` for later "copy and paste" the needed code to the generated code.
2. Removing the `resolver.go` file and redo all the previous steps.

## 4. Implement the Resolver methods.

### GroupResolver
`gqlgen` would generate a new method on the `groupResolver` called `Events`.

<details> 
  <summary>YUse the Meetup.com client for fetching data about the events of such group:</summary> 

```go
func (r *groupResolver) Events(ctx context.Context, obj *meetup.Group, status *EventStatus) ([]*meetup.Event, error) {
	return r.C.Events(obj.Urlname, (*string)(status))
}
```
</details>

### EventResolver
`gqlgen` would generate a new resolver called `eventResolver`.

<details> 
  <summary>You would need to implement some methods:</summary> 

```go
type eventResolver struct{ *Resolver }

func (r *eventResolver) Created(ctx context.Context, obj *meetup.Event) (string, error) {
	return time.Unix(obj.Created/1000, 0).Format(time.RFC822), nil
}
func (r *eventResolver) Time(ctx context.Context, obj *meetup.Event) (string, error) {
	return time.Unix(obj.Time/1000, 0).Format(time.RFC822), nil
}
func (r *eventResolver) Status(ctx context.Context, obj *meetup.Event) (EventStatus, error) {
	return EventStatus(obj.Status), nil
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
      time
      waitlistCount
      yesRSVPCount
    }
  }
}
```