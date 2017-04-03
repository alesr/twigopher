package account

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alesr/twigopher/client"
	"github.com/kurrik/twittergo"
)

// Stats wraps account usage information
type Stats struct {
	id                 string
	name               string
	RateLimit          uint32
	RateLimitRemaining uint32
	RateLimitReset     time.Time
}

// Stat retrieve account credential usage
func Stat(client *client.Client) *Stats {

	// Prepare GET request to get account info
	requestURL := "/1.1/account/verify_credentials.json"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatal("could not parse request: ", err)
	}

	// Send request
	resp, err := client.Account.SendRequest(req)
	if err != nil {
		log.Fatal("could not send request: ", err)
	}

	// Parse http response to twittergo.User struct
	// It has nice methods to manipulate the values
	user := &twittergo.User{}
	if err := resp.Parse(&user); err != nil {
		log.Fatal("failed to parse response: ", err)
	}

	// Forming our Stat struct with account data
	stats := &Stats{
		id:   user.IdStr(),
		name: user.Name(),
	}
	if resp.HasRateLimit() {
		stats.RateLimit = resp.RateLimit()
		stats.RateLimitRemaining = resp.RateLimitRemaining()
		stats.RateLimitReset = resp.RateLimitReset()
	} else {
		log.Println("could not parse rate limit from response")
	}
	return stats
}

// Implements the Stringer interface
func (s Stats) String() string {
	return fmt.Sprintf(
		"ID:                    %s\n"+
			"Name:                  %s\n"+
			"RateLimit:             %d\n"+
			"RateLimitRemaining:    %d\n"+
			"RateLimitReset:        %v",
		s.id, s.name, s.RateLimit, s.RateLimitRemaining, s.RateLimitReset,
	)
}
