package client

import (
	"errors"

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
	if err := loadConfig(); err != nil {
		return nil, err
	}

	c, err := newConsumer()
	if err != nil {
		return nil, err
	}

	a, err := newApplication()
	if err != nil {
		return nil, err
	}

	// Since we're are using two twitter clients from different packages
	// a lot of data is being overlaped and duplicated
	// The next step would be to isolate the functionalities we want to keep
	// and create a clean client
	config := &oauth1a.ClientConfig{
		ConsumerKey:    c.key,
		ConsumerSecret: c.secret,
	}

	user := oauth1a.NewAuthorizedConfig(a.accessToken, a.accessSecret)
	token := oauth1.NewToken(a.accessToken, a.accessSecret)

	httpClient := oauth1.NewConfig(c.key, c.secret).Client(oauth1.NoContext, token)

	return &Client{
		twittergo.NewClient(config, user),
		twitter.NewClient(httpClient).Streams,
	}, nil
}

// newConsumer returns the Consumer credentials
func newConsumer() (*Consumer, error) {

	if !viper.IsSet("consumer_key") {
		return nil, errors.New("could not load consumer key")
	}

	if !viper.IsSet("consumer_secret") {
		return nil, errors.New("could not load access consumer secret")
	}

	return &Consumer{
		viper.GetString("consumer_key"),
		viper.GetString("consumer_secret"),
	}, nil
}

// newApplication returns the Application credentials
func newApplication() (*Application, error) {

	if !viper.IsSet("access_token") {
		return nil, errors.New("could not load access token")
	}

	if !viper.IsSet("access_token_secret") {
		return nil, errors.New("could not load access token secret")
	}

	return &Application{
		viper.GetString("access_token"),
		viper.GetString("access_token_secret"),
	}, nil
}

// loadConfig values from config.yaml
func loadConfig() error {

	viper.SetConfigName("twitter")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("no configuration file loaded")
	}
	return nil
}
