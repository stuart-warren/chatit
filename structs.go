package main

import (
	"time"
)

type User struct {
	Id        string
	Username  string
	Nicknme   string
	StatusMsg string
	Clients   []Client
}

type Client struct {
	Id       string
	Type     string
	Location Coords
}

type Coords struct {
	Lat float32
	Lon float32
}

type Room struct {
	Id     string
	Name   string
	Desc   string
	Usage  int
	Locked bool
}

type Message struct {
	Id        string    `json:"id"`
	ClientId  string    `json:"client"`
	ReplyToId string    `json:"replyto,omitempty"`
	Room      string    `json:"room" binding:"required"`
	Timestamp time.Time `json:"@timestamp"`
	Message   string    `json:"message" binding:"required"`
	Commands  []string  `json:"commands,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Mentions  []string  `json:"mentions,omitempty"`
	Uris      []string  `json:"uris,omitempty"`
}

type Subscription struct {
	Id     string
	Client *Client
	Room   *Room
}

type ServiceMap struct {
	Id      string
	Service string
	BaseUrl string
}
