package stream

import (
	"bytes"
	"fmt"
	"testing"

	"reflect"

	"github.com/alesr/twigopher/client"
	"github.com/dghubble/go-twitter/twitter"
)

func TestNewStream(t *testing.T) {

	var buf bytes.Buffer
	track := "foo"
	client := &client.Client{}

	expected := &Stream{out: &buf, track: track, client: client}
	observed := NewStream(client, &buf, track)

	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected %T observed %T", expected, observed)
	}
}

func TestTweetHandler(t *testing.T) {

	user := &twitter.User{Name: "bar"}
	tweet := &twitter.Tweet{Text: "foo", User: user}

	var buf bytes.Buffer
	stream := &Stream{out: &buf}

	stream.tweetHandler(tweet)

	expected := fmt.Sprintf("%s: %s\n\n\n", tweet.User.Name, tweet.Text)
	observed := buf.String()

	if observed != expected {
		t.Errorf("for tweet: %+v observed: %q expected: %q",
			tweet, observed, expected)
	}
}
