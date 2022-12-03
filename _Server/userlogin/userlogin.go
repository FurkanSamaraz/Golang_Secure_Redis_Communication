package userlogin

import (
	"context"
	"encoding/json"
	"fmt"
	"main/block"
	"main/database"
	"main/jwt"
	"main/model"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var Tokens string
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Redisgetuser() {
	db := database.Data()

	subscriber := redisClient.Subscribe(ctx, "send-user-data")

	user := model.Users{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		blockmsg, _ := block.BlocksDecrypt(msg.Payload, "send-user-data----------")

		if err := json.Unmarshal([]byte(blockmsg), &user); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
		ers := db.Where(&model.Users{Name: user.Name, Password: user.Password}).First(&user)

		if ers.Error == nil {
			Tokens, _ := jwt.CreateToken(uint64(user.ID))
			Redissetuser("loginn", Tokens)
		} else {
			Redissetuser("loginn", "false")
		}
	}

}
func TokenControl(comingtoken string) {
	if Tokens != comingtoken {
		fmt.Println("token false!")
		os.Exit(1)
	} else {
		fmt.Println("token true!")

	}

}

func Redissetuser(key string, value string) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

}
