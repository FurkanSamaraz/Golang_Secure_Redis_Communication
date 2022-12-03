package companies

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/block"

	"net/http"

	"github.com/go-redis/redis/v8"
)

type Companies struct {
	Companie string `json:"companie"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Companiescreate(w http.ResponseWriter, r *http.Request) {

	companies := new(Companies)
	reqBodys, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBodys, &companies)
	payloads, err := json.Marshal(companies)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(companies)

	newDatas, err := json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newDatas))
	}
	blockpayload, _ := block.BlocksEncrypt(string(payloads), "send-create-companies---")
	if err := redisClient.Publish(ctx, "send-create-companies", blockpayload).Err(); err != nil {
		panic(err)
	}

}

type Companiesup struct {
	ID       int    `json:"id"`
	Companie string `json:"companie"`
}

func Companiesupdate(w http.ResponseWriter, r *http.Request) {
	companies := new(Companiesup)
	reqBodys, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBodys, &companies)
	payloads, err := json.Marshal(companies)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(companies)

	newDatas, err := json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newDatas))
	}
	blockpayload, _ := block.BlocksEncrypt(string(payloads), "send-create-companies---")
	if err := redisClient.Publish(ctx, "send-update-companies", blockpayload).Err(); err != nil {
		panic(err)
	}
}

func Companiesdelete(w http.ResponseWriter, r *http.Request) {
	companies := new(Companiesup)
	reqBodys, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBodys, &companies)
	payloads, err := json.Marshal(companies)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(companies)

	newDatas, err := json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newDatas))
	}
	blockpayload, _ := block.BlocksEncrypt(string(payloads), "send-create-companies---")

	if err := redisClient.Publish(ctx, "send-delete-companies", blockpayload).Err(); err != nil {
		panic(err)
	}
}
