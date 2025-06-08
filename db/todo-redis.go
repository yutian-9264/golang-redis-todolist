package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

const todoIDCounter = "todoid"
const todoIDsSet = "todos-id-set"
const statusPending = "pending"

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func CreateTodo(desc string) {
	rdb := NewRedisClient()
	defer rdb.Close()

	id, err := rdb.Incr(ctx, todoIDCounter).Result()
	if err != nil {
		log.Fatal("自增todo id失败", err)
	}
	todoid := "todo:" + strconv.Itoa(int(id))

	err = rdb.SAdd(ctx, todoIDsSet, todoid).Err()
	if err != nil {
		log.Fatal("添加todo id至Redis SET失败", err)
	}

	todo := map[string]interface{}{"desc": desc, "status": statusPending}
	err = rdb.HSet(ctx, todoid, todo).Err()
	if err != nil {
		log.Fatal("添加todo至Hash表失败", err)
	}
	fmt.Println("已成功创建一条todo!")

}
