package main

import (
	r "github.com/dancannon/gorethink"
	"log"
	"net/http"
)

type Channel struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name,omitempty"`
}

type User struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name,omitempty"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "srchat",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Creating srchat database in RethinkDB...")
	r.DBCreate("srchat").Exec(session)
	r.DB("srchat").TableCreate("channels").Exec(session)
	r.DB("srchat").TableCreate("messages").Exec(session)
	r.DB("srchat").TableCreate("users").Exec(session)

	router := NewRouter(session)

	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)
	router.Handle("channel add", addChannel)

	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)
	router.Handle("user edit", editUser)

	log.Println("Listening on port 4000...")
	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
