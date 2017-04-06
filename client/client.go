package client

import (
	"errors"

	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"github.com/spf13/viper"
)

// Consumer wraps user credentials
type Consumer struct {
	key    string
	secret string
}

// Application credentials
type Application struct {
	accessToken  string
	accessSecret string
}

// Credentials wraps consumer and application credentials
type Credentials struct {
	Consumer    *Consumer
	Application *Application
}

// Client wraps twittergo.Client for
// getting information on credential usage
// and twitter.StreamService to stream twitter data
type Client struct {
	Account *twittergo.Client
	Stream  *twitter.StreamService
}

// NewClient returns a new Client
// for account usage and stream data
func NewClient() (*Client, error) {

	// Load configuration file
	credentials, err := loadConfig("./twitter.yaml")
	if err != nil {
		return nil, err
	}

	// Since we're are using two twitter clients from different packages
	// a lot of data is being overlaped and duplicated
	// The next step would be to isolate the functionalities we want to keep
	// and create a clean client
	config := &oauth1a.ClientConfig{
		ConsumerKey:    credentials.Consumer.key,
		ConsumerSecret: credentials.Consumer.secret,
	}

	user := oauth1a.NewAuthorizedConfig(credentials.Application.accessToken, credentials.Application.accessSecret)
	token := oauth1.NewToken(credentials.Application.accessToken, credentials.Application.accessSecret)

	httpClient := oauth1.NewConfig(credentials.Consumer.key, credentials.Consumer.secret).Client(oauth1.NoContext, token)

	return &Client{
		twittergo.NewClient(config, user),
		twitter.NewClient(httpClient).Streams,
	}, nil
}

// loadConfig values from config.yaml
func loadConfig(filepath string) (*Credentials, error) {

	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("no configuration file loaded")
	}

	consumer, err := consumer()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve consumer credentials: %s", err)
	}

	application, err := application()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve application credentials: %s", err)
	}

	credentials := &Credentials{
		Consumer:    consumer,
		Application: application,
	}

	// We don't need Viper anymore since our credentials are already load
	viper.Reset()

	return credentials, nil
}

// consumer returns the Consumer credentials
func consumer() (*Consumer, error) {

	keys := []string{
		"consumer_key",
		"consumer_secret",
	}

	c := &Consumer{}

	for _, k := range keys {
		if !viper.IsSet(k) {
			return nil, fmt.Errorf("could not load: %s", k)
		}

		if k == "consumer_key" {
			c.key = viper.GetString("consumer_key")
		} else {
			c.secret = viper.GetString("consumer_secret")
		}
	}
	return c, nil
}

// newApplication returns the Application credentials
func application() (*Application, error) {

	keys := []string{
		"access_token",
		"access_token_secret",
	}

	a := &Application{}

	for _, k := range keys {
		if !viper.IsSet(k) {
			return nil, fmt.Errorf("could not load: %s", k)
		}

		if k == "access_token" {
			a.accessToken = viper.GetString("access_token")
		} else {
			a.accessSecret = viper.GetString("access_token_secret")
		}
	}
	return a, nil
}
