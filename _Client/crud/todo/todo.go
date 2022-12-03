package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/block"

	"net/http"

	"github.com/go-redis/redis/v8"
)

type Todos struct {
	Todoname string `json:"todoname"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Todocreate(w http.ResponseWriter, r *http.Request) {

	todos := new(Todos)
	reqBodys, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBodys, &todos)
	payloads, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(todos)

	newDatas, err := json.Marshal(todos)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newDatas))
	}
	blockpayload, _ := block.BlocksEncrypt(string(payloads), "send-create-todo--------")

	if err := redisClient.Publish(ctx, "send-create-todo", blockpayload).Err(); err != nil {
		panic(err)
	}

}

type Todosup struct {
	ID       int    `json:"id"`
	Todoname string `json:"todoname"`
}

func Todoupdate(w http.ResponseWriter, r *http.Request) {
	todos := new(Todosup)
	reqBodys, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBodys, &todos)
	payloads, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(todos)

	newDatas, err := json.Marshal(todos)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newDatas))
	}
	blockpayload, _ := block.BlocksEncrypt(string(payloads), "send-create-todo--------")

	if err := redisClient.Publish(ctx, "send-update-todo", blockpayload).Err(); err != nil {
		panic(err)
	}
}

func Tododelete(w http.ResponseWriter, r *http.Request) {
	todos := new(Todosup)
	reqBodys, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBodys, &todos)
	payloads, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(todos)

	newDatas, err := json.Marshal(todos)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newDatas))
	}
	blockpayload, _ := block.BlocksEncrypt(string(payloads), "send-create-todo--------")

	if err := redisClient.Publish(ctx, "send-delete-todo", blockpayload).Err(); err != nil {
		panic(err)
	}
}
