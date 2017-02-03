package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
	"log"
)

type FindHandler func(string) (Handler, bool)

type Client struct {
	send         chan Message
	socket       *websocket.Conn
	findHandler  FindHandler
	session      *r.Session
	stopChannels map[int]chan bool
	user         User
}

func (c *Client) NewStopChannel(stopKey int) chan bool {
	c.StopForKey(stopKey)
	stop := make(chan bool)
	c.stopChannels[stopKey] = stop
	return stop
}

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func (c *Client) StopForKey(key int) {
	if ch, found := c.stopChannels[key]; found {
		ch <- true
		delete(c.stopChannels, key)
	}
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}

		log.Printf("Rx name:'%s' data:'%#v'", message.Name, message.Data)
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for message := range client.send {
		log.Printf("Tx name:'%s' data:'%#v'", message.Name, message.Data)
		if err := client.socket.WriteJSON(message); err != nil {
			break
		}
	}

	client.socket.Close()
}

func (c *Client) Close() error {
	for _, ch := range c.stopChannels {
		ch <- true
	}

	close(c.send)
	err := r.Table("users").Get(c.user.Id).Delete().Exec(c.session)

	if err != nil {
		fmt.Println(err)
	}
	return err
}

func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) (*Client,
	error) {
	var user User
	user.Name = "Anonymous"
	res, err := r.Table("users").Insert(user).RunWrite(session)

	if err != nil {
		return nil, err
	}
	if len(res.GeneratedKeys) > 0 {
		user.Id = res.GeneratedKeys[0]
	}

	return &Client{
		user:         user,
		send:         make(chan Message),
		socket:       socket,
		findHandler:  findHandler,
		session:      session,
		stopChannels: make(map[int]chan bool),
	}, nil
}
