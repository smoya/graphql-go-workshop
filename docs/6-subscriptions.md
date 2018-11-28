# 6. Subscription to any comment on a Meetup.com Event

> Subscriptions are a GraphQL feature that allows a server to send data to its clients when a specific event happens. 
> Subscriptions are usually implemented with WebSockets.

In this chapter you will bring realtime functionality to your GraphQL API by implementing GraphQL subscriptions.

The goal of this chapter is to create a subscription to the comments posted on a Meetup.com Event.

## 1. Edit your `graphql.schema` file.

Let's add some new types. 
As before, you can take a look to the model located in the [github.com/smoya/graphql-go-workshop/pkg/meetup](https://github.com/smoya/graphql-go-workshop/blob/master/pkg/meetup/model.go) package.

Add the following types:
```graphql schema
type Comment {
    id: Int!
    comment: String!
    created: String!
    likes: Int!
    member: Member!
}

type Subscription {
    commentPosted(groupName: String!, eventID: String!): Comment!
}
```

## 2. Edit your `gqlgen.yaml` file.

Before generating the code again, we would need to add few new lines to the `gqlgen.yaml` file.

Please add the following lines under `model`.
```yaml
  Comment:
    model: github.com/smoya/graphql-go-workshop/pkg/meetup.Comment
```

## 3. Regenerate your code

> Note 
> The `resolver.go` file is generated **only** when the file is not present. 

As the generator will end creating new Resolvers, you would need to choose 1 option from:

1. Moving the `resolver.go` file to a `resolver.go.copy` before executing `gqlgen` for later "copy and paste" the needed code to the generated code.
2. Removing the `resolver.go` file and redo all the previous steps.

## 4. Implement the new `Comment` and `Subscription` Resolver.

### CommentResolver
`gqlgen` would generate a new resolver called `commentResolver`. 

<details> 
  <summary>You would need to implement some methods:</summary> 

```go
// Comment returns a CommentResolver.
func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}

type commentResolver struct{ *Resolver }

func (commentResolver) Likes(ctx context.Context, obj *meetup.Comment) (int, error) {
	return obj.LikeCount, nil
}

func (commentResolver) Created(ctx context.Context, obj *meetup.Comment) (string, error) {
	return time.Unix(obj.Created/1000, 0).Format(time.RFC822), nil
}
```
</details>

### SubscriptionResolver
`gqlgen` would generate a new resolver called `subscriptionResolver`.



<details> 
  <summary>You would need to implement some methods:</summary>
  
```go
// Subscription returns a SubscriptionResolver
func (r *Resolver) Subscription() SubscriptionResolver {
return &subscriptionResolver{r}
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) CommentPosted(ctx context.Context, groupName string, eventID string) (<-chan meetup.Comment, error) {
// ideally this should go in to a redis or similar.
   sentComments := make(map[int]struct{})
   commentsChan := make(chan meetup.Comment)

   // ideally this should be configurable
   t := time.NewTicker(time.Second * 5)
   go func() {
       for {
           select {
           case <-ctx.Done():
               return
           case <-t.C:
               comments, err := r.C.Comments(groupName, eventID)
               if err != nil {
                   log.Printf("error finding comments for group %s and event %s. %s\n", groupName, eventID, err.Error())
                   continue
               }

               for _, c := range comments {
                   if _, ok := sentComments[c.ID]; !ok {
                       commentsChan <- *c
                       sentComments[c.ID] = struct{}{}
                   }
               }
           }
       }

   }()

return commentsChan, nil
}
```
</details>


## 5. Play

Run `go run server/server.go`.
Open your browser and navigate to [http://localhost:8080](http://localhost:8080). 

Execute the following subscription:

```graphql
subscription{
  commentPosted(groupName: "Golang-Barcelona", eventID: "256537826"){
    id
    comment
    created
    likes
    member {
      name
    }
  }
}
```