package main

import (
	// "github.com/go-redis/redis"
	"github.com/gin-gonic/gin"	
	"github.com/no7dw/fan404/db"

	"net/http"
)
var redisClient = db.InitRedis()

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value , err := redisClient.Get(user).Result()
		if value != "" && err == nil {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			err := redisClient.Set(user, json.Value, 0).Err()
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
