package account

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {

	s := &Stats{
		id:                 "01",
		name:               "foo",
		RateLimit:          10,
		RateLimitRemaining: 5,
	}

	observed := s.String()
	expected := fmt.Sprintf(
		"ID:                    01\n"+
			"Name:                  foo\n"+
			"RateLimit:             10\n"+
			"RateLimitRemaining:    5\n"+
			"RateLimitReset:        %v", s.RateLimitReset,
	)

	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("expected %+v observed %+v", observed, expected)
	}
}
