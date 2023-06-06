package main

import (
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitialiseRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       redisDb,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic("Redis :: Failed to connect to Redis:" + err.Error())
	}
}

func GetRedisClient() *redis.Client {
	return rdb
}

func SetInRedis(key string, value string) {
	rdb := GetRedisClient()
	err := rdb.Set(key, value, redisKeyLifetime).Err()
	if err != nil {
		log.Println("Redis :: Unable to set in Redis with key: ", key, "due to: ", err.Error())
	}
}

func GetFromRedis(key string) string {
	rdb := GetRedisClient()
	val, err := rdb.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return ""
		} else {
			log.Println("Redis :: Unable to get from Redis with key: ", key, "due to: ", err.Error())
		}
	}
	return val
}

func FormRedisKey(cargoToml string, mainRs string) string {
	project := cargoToml + mainRs
	base64dProject := base64Encoder(project)

	hashedEncodedProject := GetMd5StringOfInput(base64dProject)
	return hashedEncodedProject
}
