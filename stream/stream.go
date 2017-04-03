package stream

import (
	"fmt"

	"io"

	"github.com/alesr/twigopher/client"
	"github.com/dghubble/go-twitter/twitter"
)

// Stream ...
type Stream struct {
	out    io.Writer
	stream *twitter.Stream
	client *client.Client
	track  string
}

// NewStream ...
func NewStream(client *client.Client, out io.Writer, track string) *Stream {
	return &Stream{
		out:    out,
		track:  track,
		client: client,
	}
}

// Start straming tweets containing track values
func (s *Stream) Start() error {

	params := &twitter.StreamFilterParams{
		Track:         []string{s.track},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := s.client.Stream.Filter(params)
	if err != nil {
		return fmt.Errorf("failed to start streaming tweets: %s", err)
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = s.tweetHandler

	for message := range stream.Messages {
		demux.Handle(message)
	}
	return nil
}

// Close the stream
func (s *Stream) Close() {
	s.stream.Stop()
}

func (s *Stream) tweetHandler(tweet *twitter.Tweet) {
	fmt.Fprintf(
		s.out,
		"%s: %s\n\n\n",
		tweet.User.Name,
		tweet.Text,
	)
}
