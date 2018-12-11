# Heroku self ping for Golang

## Installation:

```go
go get github.com/prongbang/heroku-self-ping
```

## Usage:

```go
import selfping "github.com/prongbang/heroku-self-ping"

selfping.HerokuSelfPing("https://google.com", selfping.Options{
    Interval: 20 * 60 * 1000, // 20 Minute
    Verbose:  true,           // Show log
})
```

### Reference
- [Heroku self ping for Node.js](https://github.com/Neamar/heroku-self-ping)
