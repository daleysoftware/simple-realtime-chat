package main

import (
	r "github.com/dancannon/gorethink"
	"github.com/mitchellh/mapstructure"
)

const (
	ChannelStop = iota
	UserStop
	MessageStop
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	go func() {
		err := r.Table("channels").Insert(channel).Exec(client.session)

		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()
}

func subscribeChannel(client *Client, data interface{}) {
	stop := client.NewStopChannel(ChannelStop)
	result := make(chan r.ChangeResponse)

	cursor, err := r.Table("channels").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(client.session)

	if err != nil {
		client.send <- Message{"error", err.Error()}
	}

	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			result <- change

		}
	}()

	go func() {
		for {
			select {
			case <-stop:
				cursor.Close()
				return
			case change := <-result:
				if change.NewValue != nil && change.OldValue == nil {
					client.send <- Message{"channel add", change.NewValue}
				}
			}
		}
	}()
}

func unsubscribeChannel(client *Client, data interface{}) {
	client.StopForKey(ChannelStop)
}

func subscribeUser(client *Client, data interface{}) {
	stop := client.NewStopChannel(UserStop)
	result := make(chan r.ChangeResponse)

	cursor, err := r.Table("users").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(client.session)

	if err != nil {
		client.send <- Message{"error", err.Error()}
	}

	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			result <- change

		}
	}()

	go func() {
		for {
			select {
			case <-stop:
				cursor.Close()
				return
			case change := <-result:
				if change.NewValue != nil && change.OldValue != nil {
					client.send <- Message{"user edit", change.NewValue}
				}
				if change.NewValue != nil && change.OldValue == nil {
					client.send <- Message{"user add", change.NewValue}
				}
				if change.NewValue == nil && change.OldValue != nil {
					client.send <- Message{"user remove", change.OldValue}
				}
			}
		}
	}()
}

func unsubscribeUser(client *Client, data interface{}) {
	client.StopForKey(UserStop)
}

func editUser(client *Client, data interface{}) {
	var user User
	err := mapstructure.Decode(data, &user)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	user.Id = client.user.Id

	go func() {
		err := r.Table("users").Update(user).Exec(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
		}
		client.user = user
	}()
}
