package meetup

// Venue represents a Meetup.com venue.
type Venue struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Lat                  float64 `json:"lat"`
	Lon                  float64 `json:"lon"`
	Repinned             bool    `json:"repinned"`
	Address              string  `json:"address_1"`
	City                 string  `json:"city"`
	Country              string  `json:"country"`
	LocalizedCountryName string  `json:"localized_country_name"`
}

// Group represents a Meetup.com group.
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

// Event represents a Meetup.com event.
type Event struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Created       int64  `json:"created"`
	Duration      int    `json:"duration"`
	RsvpLimit     int    `json:"rsvp_limit"`
	Status        string `json:"status"`
	Time          int64  `json:"time"`
	LocalDate     string `json:"local_date"`
	LocalTime     string `json:"local_time"`
	Updated       int64  `json:"updated"`
	UtcOffset     int    `json:"utc_offset"`
	WaitlistCount int    `json:"waitlist_count"`
	YesRsvpCount  int    `json:"yes_rsvp_count"`
	Venue         Venue  `json:"venue"`
	Group         Group  `json:"group"`
	Link          string `json:"link"`
	Description   string `json:"description"`
	Visibility    string `json:"visibility"`
}

// Member represents a member from a Meetup.com group.
type Member struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	EventContext EventContext `json:"event_context"`
}

// EventContext contains info about the relation between a Meetup.com event and a member who rspv'ed.
type EventContext struct {
	Host bool `json:"host"`
}

// RSVP represents a rsvp in a Meetup.com event.
type RSVP struct {
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
	Response string `json:"response"`
	Guests   int    `json:"guests"`
	Member   Member `json:"member"`
}

// Comment represents a commend in any Meetup.com group or event.
type Comment struct {
	ID        int    `json:"id"`
	Comment   string `json:"comment"`
	Link      string `json:"link"`
	Created   int64  `json:"created"`
	LikeCount int    `json:"like_count"`
	Member    Member `json:"member"`
}
