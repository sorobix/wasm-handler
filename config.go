// package main

// import "time"

// var wsPort = "3001"
// var redisHost = "redis:6379"
// var redisPassword = "" // no password set
// var redisDb = 0        // use default DB
// var redisKeyLifetime = time.Hour

package main

import (
	"os"
	"strconv"
	"time"
)

var wsPort = ""
var redisHost = ""
var redisPassword = ""
var redisDb = 0
var redisKeyLifetime = time.Hour

func init() {
	wsPort = os.Getenv("WS_PORT")
	redisHost = os.Getenv("REDIS_HOST")
	redisPassword = os.Getenv("REDIS_PASSWORD")

	dbStr := os.Getenv("REDIS_DB")
	if db, err := strconv.Atoi(dbStr); err == nil {
		redisDb = db
	}

	lifetimeStr := os.Getenv("REDIS_KEY_LIFETIME")
	if lifetime, err := time.ParseDuration(lifetimeStr); err == nil {
		redisKeyLifetime = lifetime
	}
}
