package controller

import (
	"database/sql"
	"log"

	"github.com/redis/go-redis/v9"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/api_tools?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal()
	}
	return db
}

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	return rdb
}
