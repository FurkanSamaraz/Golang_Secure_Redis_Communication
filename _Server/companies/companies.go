package companies

import (
	"context"
	"encoding/json"
	"fmt"
	"main/block"
	"main/database"
	"main/model"
	"main/userlogin"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Redistoken() string {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, _ := rdb.Get(ctx, "send-user-token").Result()

	fmt.Println("send-user-token", val)
	return val
}

var A = Redistoken()

func Companiescreate() {

	db := database.Data()

	subscriber := redisClient.Subscribe(ctx, "send-create-companies")

	companies := model.Companies{}
	userlogin.TokenControl(A)
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-create-companies---")
		if err := json.Unmarshal([]byte(blockmsg), &companies); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", companies)
		db.Create(&model.Companies{Companie: companies.Companie})

	}

}

type Companies struct {
	ID       int    `json:"id"`
	Companie string `json:"companie"`
}

func Companiesupdate() {
	db := database.Data()
	userlogin.TokenControl(A)
	subscriber := redisClient.Subscribe(ctx, "send-update-companies")

	companies := Companies{}
	com := model.Companies{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-create-companies---")
		if err := json.Unmarshal([]byte(blockmsg), &companies); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", companies)
		db.First(&com, companies.ID)

		db.Model(&companies).Updates(model.Companies{Companie: companies.Companie})

	}
}

func Companiesdelete() {
	db := database.Data()
	userlogin.TokenControl(A)
	subscriber := redisClient.Subscribe(ctx, "send-delete-companies")

	companies := Companies{}
	com := model.Companies{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-create-companies---")

		if err := json.Unmarshal([]byte(blockmsg), &companies); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", companies)

		db.Delete(&com, companies.ID) // id 1 olan kullanıcı silinecektir.
		fmt.Println(com)

	}
}
