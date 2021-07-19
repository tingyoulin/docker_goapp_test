package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var rdb = newClient() // declare redis client

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(recordUserMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	router.GET("/ping1", func(c *gin.Context) {
		// member := c.Param("member")

		c.String(http.StatusOK, "Pong1")
	})

	router.GET("/get_visits", func(c *gin.Context) {
		// member := c.Param("member")
		userIP := c.ClientIP()
		val, err := rdb.Get(userIP).Result()

		switch {
		case err == redis.Nil:
			c.String(http.StatusBadRequest, fmt.Sprint("key does not exist"))
		case err != nil:
			c.String(http.StatusBadRequest, fmt.Sprint(err))
		case val == "":
			c.String(http.StatusBadRequest, fmt.Sprint("value is empty"))
		}

		c.String(http.StatusOK, fmt.Sprintf("Got!\n%s visits: %s", userIP, val))
	})

	router.GET("/delete/", func(c *gin.Context) {
		userIP := c.ClientIP()
		err := rdb.Del(userIP).Err()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprint(err))
		}
		c.String(http.StatusOK, fmt.Sprintf("%s deleted!", userIP))
	})

	return router
}

func recordUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIP := c.ClientIP()
		log.Printf("User IP: %s", userIP)

		err := rdb.Incr(userIP).Err()
		if err != nil {
			err := rdb.Set(userIP, 1, 0).Err()

			if err != nil {
				log.Println(fmt.Sprint(err))
			}

			log.Printf("create: %s\n", userIP)
			return
		}

	}
}

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // redis 是 docker-compose.yml 裡的 service name (or alias)
		Password: "0000",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatalln(err)
	}

	return client
}

// type Member struct {
// 	name   string
// 	visits string
// }

func main() {
	router := setupRouter()
	router.Run() // default localhost:8080
}
