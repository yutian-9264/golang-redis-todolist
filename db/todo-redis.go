package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

const todoIDCounter = "todoid"
const todoIDsSet = "todos-id-set"
const statusPending = "pending"

type Todo struct {
	ID     string
	Desc   string
	Status string
}

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

func ListTodos(status string) []Todo {
	rdb := NewRedisClient()
	defer rdb.Close()

	todoHashKeys, err := rdb.SMembers(ctx, todoIDsSet).Result()
	if err != nil {
		log.Fatal("查找todo任务的名称失败", err)
	}

	todos := []Todo{}

	for _, todoHashkey := range todoHashKeys {
		id := strings.Split(todoHashkey, ":")[1]

		todoMap, err := rdb.HGetAll(ctx, todoHashkey).Result()
		if err != nil {
			log.Fatalf("读“%s”todo任务失败 - %v\n", todoHashkey, err)
		}

		var todo Todo
		if status == "" {
			todo = Todo{id, todoMap["desc"], todoMap["status"]}
			todos = append(todos, todo)
		} else {
			if status == todoMap["status"] {
				todo = Todo{id, todoMap["desc"], todoMap["status"]}
				todos = append(todos, todo)
			}
		}
	}

	if len(todos) == 0 {
		fmt.Println("没有Todo任务")
		return nil
	}

	return todos
}

func UpdateTodo(id, desc, status string) {
	rdb := NewRedisClient()
	defer rdb.Close()

	exists, err := rdb.SIsMember(ctx, todoIDsSet, "todo:"+id).Result()
	if err != nil {
		log.Fatalf("确定todo(%s)是否存在前出现错误(%v)", id, err)
	}

	if !exists {
		log.Fatalf("您想要修改的todo(%s)不存在", id)
	}

	updatedTodo := map[string]interface{}{}

	if status != "" {
		updatedTodo["status"] = status
	}

	if desc != "" {
		updatedTodo["desc"] = desc
	}

	err = rdb.HSet(ctx, "todo:"+id, updatedTodo).Err()
	if err != nil {
		log.Fatalf("无法更新该todo(%s)", id)
	}
	fmt.Printf("该todo(%s)已更新\n", id)

}
