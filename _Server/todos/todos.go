package todos

import (
	"context"
	"encoding/json"
	"fmt"
	"main/block"
	"main/companies"
	"main/database"
	"main/model"
	"main/userlogin"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
var A = companies.Redistoken()

func Todocreate() {
	db := database.Data()

	subscriber := redisClient.Subscribe(ctx, "send-create-todo")

	todos := model.Todos{}
	userlogin.TokenControl(A)
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-create-todo--------")

		if err := json.Unmarshal([]byte(blockmsg), &todos); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", todos)
		db.Create(&model.Todos{Todoname: todos.Todoname})

	}

}

type Todos struct {
	ID       int    `json:"id"`
	Todoname string `json:"todoname"`
}

func Todoupdate() {
	db := database.Data()
	userlogin.TokenControl(A)
	subscriber := redisClient.Subscribe(ctx, "send-update-todo")

	todo := Todos{}
	com := model.Todos{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-create-todo--------")

		if err := json.Unmarshal([]byte(blockmsg), &todo); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", todo)
		db.First(&com, todo.ID)

		db.Model(&todo).Updates(model.Todos{Todoname: todo.Todoname})

	}
}
func Tododelete() {
	db := database.Data()
	userlogin.TokenControl(A)
	subscriber := redisClient.Subscribe(ctx, "send-delete-companies")

	todo := Todos{}
	com := model.Todos{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-create-todo--------")

		if err := json.Unmarshal([]byte(blockmsg), &todo); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", todo)

		db.Delete(&com, todo.ID) // id 1 olan kullanıcı silinecektir.
		fmt.Println(com)

	}
}
