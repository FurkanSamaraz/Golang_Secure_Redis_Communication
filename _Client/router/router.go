package router

import (
	"main/crud/companies"
	"main/crud/todo"
	userlogin "main/login"
	usercreate "main/user"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()

	r.HandleFunc("/create", usercreate.Usercreate)
	r.HandleFunc("/login", userlogin.Userlogin)

	r.HandleFunc("/companiescreate", companies.Companiescreate)
	r.HandleFunc("/companiesupdate", companies.Companiesupdate)
	r.HandleFunc("/companiesdelete", companies.Companiesdelete)

	r.HandleFunc("/todocreate", todo.Todocreate)
	r.HandleFunc("/todoupdate", todo.Todoupdate)
	r.HandleFunc("/tododelete", todo.Tododelete)
	http.ListenAndServe(":8080", r)

}
