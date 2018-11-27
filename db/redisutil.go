package db

import "github.com/go-redis/redis"

func InitRedis() *redis.Client {
       r := redis.NewClient(
               &redis.Options{
                       Addr: "localhost:6379",
                       Password: "",
                       DB: 0,})
       return r
}


