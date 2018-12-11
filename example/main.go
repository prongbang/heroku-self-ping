package main

import (
	"fmt"
	"net/http"
	"time"

	selfping "github.com/prongbang/heroku-self-ping"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {

	selfping.HerokuSelfPing("https://google.com", selfping.Options{
		Interval: 10 * 1000, // 10 Seconds
		Verbose:  false,
	})

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
