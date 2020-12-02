package main

import (
	"github.com/gin-gonic/gin"
	redisDriver "github.com/go-redis/redis/v8"
	"github.com/kenretto/sessions"
	"github.com/kenretto/sessions/redis"
)

func main() {
	r := gin.Default()
	storeDB := redisDriver.NewClient(&redisDriver.Options{
		Addr: "localhost:6379",
	})
	store := redis.NewStore(storeDB, []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
