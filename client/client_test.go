package client

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	expected := &Credentials{
		Consumer: &Consumer{
			key:    "consumer_key",
			secret: "consumer_secret",
		},
		Application: &Application{
			accessToken:  "access_token",
			accessSecret: "access_token_secret",
		},
	}

	observed, err := loadConfig("../resources/test_credentials.yaml")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("observed: %+v, expected: %+v", observed, expected)
	}
}
