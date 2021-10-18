package main

import (
	"fmt"
	"kumparan/controller"
	models "kumparan/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	r := gin.Default()
	models.ConnectDatabase()

	api := r.Group("/project/kumparan")
	{
		api.GET("/articles", controller.FindArticles)
		api.POST("/articles", controller.CreateArticles)
	}

	r.Run(":8889")
}
