package main

import (
	"strconv"
)

type User struct {
	ID        int64  `json:id`
	Name      string `json:name`
	Image     string `json:image`
	FBID      int64  `json:fbid`
	Email     string `json:email`
	CreatedOn string `json:created_on`
}

var userTable string = "user"

func NewUser() (*User, error) {
	return &User{}, nil
}

func (this *User) Login() bool {
	return true
}

func (this *User) Logout() bool {
	return true
}

func (this *User) Load() bool {
	where := map[string]interface{}{}

	if this.ID != 0 {
		where["id"] = this.ID
	}

	if this.FBID != 0 {
		where["fbid"] = this.FBID
	}

	if this.Email != "" {
		where["email"] = this.Email
	}

	userRow, err := MySQLSelect(userTable, where, []string{"id", "name", "image", "fbid", "email", "created_on"})
	if err != nil {
		return false
	}

	this.ID, _ = strconv.ParseInt(userRow[0]["id"], 10, 64)
	this.Name = userRow[0]["name"]
	this.Image = userRow[0]["image"]
	this.FBID, _ = strconv.ParseInt(userRow[0]["fbid"], 10, 64)
	this.Email = userRow[0]["email"]
	this.CreatedOn = userRow[0]["created_on"]

	return true
}

func (this *User) Save() bool {
	success := false

	if this.ID == 0 {
		success, newID := MySQLInsert(userTable, Struct2Map(this))

		if !success {
			return false
		}

		this.ID = newID
	} else {
		success = MySQLUpdate(userTable, Struct2Map(this), map[string]interface{}{"id": this.ID})
	}

	return success
}
