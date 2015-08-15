package main

import ()

type Message struct {
	ID          int64  `json:id`
	Type        string `json:type`
	Points      string `json:points`
	PhoneNumber string `json:phone_number`
	GroupID     int64  `json:group_id`
	CreatedOn   string `json:created_on`
}

func NewMessage() (*Message, error) {
	return &Message{}, nil
}

func (this *Message) Load() bool {
	return true
}

func (this *Message) Save() bool {
	return true
}
