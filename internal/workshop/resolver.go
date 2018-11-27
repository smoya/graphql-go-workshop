//go:generate gorunpkg github.com/99designs/gqlgen

package workshop

import (
	context "context"
	"time"

	meetup "github.com/smoya/graphql-go-workshop/pkg/meetup"
)

// Resolver is the root GraphQL Resolver.
type Resolver struct {
	C *meetup.Client
}

// Group returns a GroupResolver
func (r *Resolver) Group() GroupResolver {
	return &groupResolver{r}
}

// Event returns a EventResolver
func (r *Resolver) Event() EventResolver {
	return &eventResolver{r}
}

// Member returns a MemberResolver
func (r *Resolver) Member() MemberResolver {
	return &memberResolver{r}
}

// Query returns a QueryResolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Rsvp returns a RsvpResolver
func (r *Resolver) Rsvp() RsvpResolver {
	return &rsvpResolver{r}
}

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
func (r *eventResolver) Rsvp(ctx context.Context, obj *meetup.Event, response *RsvpResponse) ([]*meetup.RSVP, error) {
	return r.C.RSVPs(obj.Group.Urlname, obj.ID, (*string)(response))
}

type memberResolver struct{ *Resolver }

func (r *memberResolver) IsHost(ctx context.Context, obj *meetup.Member) (bool, error) {
	return obj.EventContext.Host, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Group(ctx context.Context, name string) (*meetup.Group, error) {
	return r.C.Group(name)
}

type groupResolver struct{ *Resolver }

func (r *groupResolver) Events(ctx context.Context, obj *meetup.Group, status *EventStatus) ([]*meetup.Event, error) {
	return r.C.Events(obj.Urlname, (*string)(status))
}

type rsvpResolver struct{ *Resolver }

func (r *rsvpResolver) Created(ctx context.Context, obj *meetup.RSVP) (string, error) {
	return time.Unix(obj.Created/1000, 0).Format(time.RFC822), nil
}
func (r *rsvpResolver) Updated(ctx context.Context, obj *meetup.RSVP) (string, error) {
	return time.Unix(obj.Updated/1000, 0).Format(time.RFC822), nil
}
func (r *rsvpResolver) Response(ctx context.Context, obj *meetup.RSVP) (RsvpResponse, error) {
	return RsvpResponse(obj.Response), nil
}
