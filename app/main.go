package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/db", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"db": getDB("testdb"),
		})
	})

	r.GET("/redis", func(c *gin.Context) {

		key := "test"

		client := getRedis()
		err := client.Set(key, "AAAAAAAAAAAAAAAAAAAA", 0).Err()
		if err != nil {
			fmt.Println("redis.Client.Set Error:", err)
		}

		val, err := client.Get(key).Result()
		if err != nil {
			fmt.Println("redis.Client.Get Error:", err)
		}

		c.JSON(200, gin.H{
			key: val,
		})
	})

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)
}

func getDB(dbName string) *gorm.DB {

	adapter := os.Getenv("DB_ADAPTER")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")

	if len(adapter) == 0 {
		adapter = "postgres"
		host = "postgresql"
		user = "root"
		pass = "root"
		port = "5432"
	}

	conn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, pass, port)

	db, err := gorm.Open(adapter, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connect: ", &db)
	return db
}

func getRedis() *redis.Client {

	host := os.Getenv("REDIS_HOST")
	pass := os.Getenv("REDIS_PASS")
	port := os.Getenv("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})

	return client
}
