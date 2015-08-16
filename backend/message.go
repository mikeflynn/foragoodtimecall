package main

import (
	"strconv"
)

type Message struct {
	ID          int64  `json:id`
	Type        string `json:type`
	Points      int    `json:points`
	PhoneNumber string `json:phone_number`
	ContestID   int64  `json:contest_id`
	CreatedOn   string `json:created_on`
}

type MessageList struct {
	Items []*Message `json:items`
}

var messageTable string = "message"

func MessageNew() (*Message, error) {
	return &Message{}, nil
}

func (this *Message) Load() bool {
	where := map[string]interface{}{}

	if this.ID != 0 {
		where["id"] = this.ID
	}

	rows, err := MySQLSelect(messageTable, where, []string{"id", "type", "phonenumber", "points", "payload", "contest_id", "created_on"})
	if err != nil {
		return false
	}

	this.ID = rows[0]["id"]
	this.Type = rows[0]["type"]
	this.Points, _ = strconv.ParseInt(rows[0]["points"], 10, 64)
	this.PhoneNumber = rows[0]["phonenumber"]
	this.ContestID, _ = strconv.ParseInt(rows[0]["contest_id"], 10, 64)
	this.CreatedOn = rows[0]["created_on"]

	return true
}

func (this *Message) Save() bool {
	success := false

	if this.ID == 0 {
		success, newID := MySQLInsert(messageTable, Struct2Map(this))

		if !success {
			return false
		}

		this.ID = newID
	} else {
		success = MySQLUpdate(messageTable, Struct2Map(this), map[string]interface{}{"id": this.ID})
	}

	return success
}
