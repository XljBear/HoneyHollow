package redis

import (
	"HoneyHollow/server/config"
	"fmt"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func GetRedis() *redis.Client {
	if redisDB == nil {
		var err error
		redisDB, err = InitRedis()
		if err != nil {
			//Todo log
			return nil
		}
	}
	return redisDB
}

func InitRedis() (client *redis.Client, err error) {
	redisHost := config.RedisHost
	redisPort := config.RedisPort

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: config.RedisPassword,
		DB:       config.RedisDBIndex,
	})
	_, err = client.Ping().Result()
	return
}
