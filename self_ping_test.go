package selfping_test

import (
	"testing"

	selfping "github.com/prongbang/heroku-self-ping"
)

func TestStartInterval(t *testing.T) {

	counter := 0

	options := selfping.Options{
		Interval: 1 * 1000,
		Verbose:  false,
	}

	selfping.SetInterval(func() {
		counter++
	}, options.Interval, false)

	for {
		if counter == 1 {
			return
		}
	}
}
