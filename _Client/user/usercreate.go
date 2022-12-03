package usercreate

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/block"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Usercreate(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	payload, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(user)

	newData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}
	blockpayload, _ := block.BlocksEncrypt(string(payload), "send-user-create--------")

	if err := redisClient.Publish(ctx, "send-user-create", blockpayload).Err(); err != nil {
		panic(err)
	}

}
