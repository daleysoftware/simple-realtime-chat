package main

import (
	r "github.com/dancannon/gorethink"
	"log"
	"net/http"
	"time"
)

type Channel struct {
	ChannelId string `json:"channelId" gorethink:"channelId,omitempty"`
	Name      string `json:"name" gorethink:"name,omitempty"`
}

type User struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name,omitempty"`
}

type Message struct {
	ChannelId string    `json:"channelId" gorethink:"channelId,omitempty"`
	Author    string    `json:"author" gorethink:"author,omitempty"`
	CreatedAt time.Time `json:"createdAt" gorethink:"createdAt,omitempty"`
	Body      string    `json:"body" gorethink:"body,omitempty"`
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
	r.DB("srchat").TableCreate("users").Exec(session)
	r.DB("srchat").TableCreate("messages").Exec(session)
	r.DB("srchat").Table("messages").IndexCreate("createdAt")

	router := NewRouter(session)

	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)
	router.Handle("channel add", addChannel)

	router.Handle("message subscribe", subscribeMessage)
	router.Handle("message unsubscribe", unsubscribeMessage)
	router.Handle("message add", addMessage)

	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)
	router.Handle("user edit", editUser)

	log.Println("Listening on port 5000...")
	http.Handle("/", router)

	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Panic(err.Error())
	}
}
