package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var wsPort = ""
var redisHost = ""
var redisPassword = ""
var redisDb = 0
var redisKeyLifetime = time.Hour

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

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
