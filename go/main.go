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

	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
