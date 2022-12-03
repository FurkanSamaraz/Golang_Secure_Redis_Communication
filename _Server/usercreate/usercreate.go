package usercreate

import (
	"context"
	"encoding/json"
	"fmt"
	"main/block"
	"main/database"
	"main/model"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

type Users struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Rediscreate() {
	db := database.Data()

	subscriber := redisClient.Subscribe(ctx, "send-user-create")

	user := model.Users{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-user-create-------")

		if err := json.Unmarshal([]byte(blockmsg), &user); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
		db.Create(&model.Users{Name: user.Name, Password: user.Password})
		fmt.Println(msg)
	}

}
