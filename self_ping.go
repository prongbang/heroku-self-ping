package selfping

import (
	"log"
	"net/http"
	"time"
)

type Options struct {
	Interval int
	Verbose  bool
}

func HerokuSelfPing(url string, options Options) {

	if url == "" {
		if options.Verbose {
			log.Println("heroku-self-ping: no url provided. Exiting.")
		}
		return
	}

	if options.Verbose {
		log.Println("heroku-self-ping: Setting up heartbeat to " + url + " every " + string(options.Interval) + "ms.")
	}

	SetInterval(func() {
		resp, err := http.Get(url)
		if options.Verbose {
			log.Println("heroku-self-ping: Sending heartbeat to " + url)
		}
		if err != nil {
			if options.Verbose {
				log.Printf("heroku-self-ping: Sending heartbeat error %s", err)
			}
		}
		defer resp.Body.Close()
	}, options.Interval, false)
}

func SetInterval(doFunc func(), milliseconds int, async bool) chan bool {

	interval := time.Duration(milliseconds) * time.Millisecond

	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				if async {
					go doFunc()
				} else {
					doFunc()
				}
			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()
	return clear

}
