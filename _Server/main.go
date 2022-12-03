package main

import (
	"main/companies"
	"main/database"
	"main/usercreate"
	"main/userlogin"
)

func execute() {
	go usercreate.Rediscreate()
	go userlogin.Redisgetuser()
	go companies.Companiescreate()
	go companies.Companiesupdate()
	go companies.Companiesdelete()
}

var wait chan string

func main() {
	database.Data()
	execute()
	<-wait
}
