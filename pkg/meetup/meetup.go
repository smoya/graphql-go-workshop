package meetup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	apiURL = "https://api.meetup.com"
)

type param struct {
	key   string
	value string
}

// Client is a client for communicating with the Meetup.com v3 API.
type Client struct {
	c      *http.Client
	apiKey string
}

// NewClient creates a new Meetup.com v3 API client.
func NewClient(c *http.Client, apiKey string) *Client {
	return &Client{
		c:      c,
		apiKey: apiKey,
	}
}

func (c *Client) doGet(path string, params ...param) (*http.Response, error) {
	u, _ := url.Parse(apiURL)
	u.Path = path
	q := u.Query()
	q.Set("sign", "true")
	q.Set("key", c.apiKey)

	for _, p := range params {
		q.Set(p.key, p.value)
	}
	u.RawQuery = q.Encode()

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to meetup.com API returned a %d status code", r.StatusCode)
	}

	return r, err
}

// Group returns a Meetup group.
func (c *Client) Group(groupName string) (*Group, error) {
	r, err := c.doGet(groupName)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close() // nolint: errcheck

	var group *Group
	err = json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		return nil, errors.Wrap(err, "invalid or missing payload")
	}

	return group, nil
}

// Events returns the listing of all Meetup Events hosted by a target group.
func (c *Client) Events(groupName string, status *string) ([]*Event, error) {
	var params []param
	if status != nil {
		params = append(params, param{
			key:   "status",
			value: *status,
		})
	}

	r, err := c.doGet(fmt.Sprintf("%s/events", groupName), params...)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close() // nolint: errcheck

	var events []*Event
	err = json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		return nil, errors.Wrap(err, "invalid or missing payload")
	}

	if len(events) == 0 {
		return nil, nil
	}

	return events, nil
}

// RSVPs return the list of RSVP of an event.
func (c *Client) RSVPs(groupName, eventID string, response *string) ([]*RSVP, error) {
	var params []param
	if response != nil {
		params = append(params, param{
			key:   "response",
			value: *response,
		})
	}
	r, err := c.doGet(fmt.Sprintf("%s/events/%s/rsvps", groupName, eventID), params...)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close() // nolint: errcheck

	var rsvps []*RSVP
	err = json.NewDecoder(r.Body).Decode(&rsvps)
	if err != nil {
		return nil, errors.Wrap(err, "invalid or missing payload")
	}

	return rsvps, nil
}

// Comments returns the list of comments of an event.
func (c *Client) Comments(groupName, eventID string) ([]*Comment, error) {
	r, err := c.doGet(fmt.Sprintf("%s/events/%s/comments", groupName, eventID))
	if err != nil {
		return nil, err
	}

	defer r.Body.Close() // nolint: errcheck

	var comments []*Comment
	err = json.NewDecoder(r.Body).Decode(&comments)
	if err != nil {
		return nil, errors.Wrap(err, "invalid or missing payload")
	}

	return comments, nil
}
