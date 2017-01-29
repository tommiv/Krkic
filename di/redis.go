package di

import (
    "gopkg.in/redis.v5"
    log "gopkg.in/Sirupsen/logrus.v0"
)

var RedisImages *redis.Client
var RedisOther  *redis.Client

func SetupRedis() {
    RedisImages = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       6,
    })

    _, err := RedisImages.Ping().Result()
    if err == nil {
        log.Info("Redis images bucket online")
    } else {
        panic(err)
    }

    RedisOther = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       7,
    })

    _, err = RedisOther.Ping().Result()
    if err == nil {
        log.Info("Redis other bucket online")
    } else {
        panic(err)
    }
}
