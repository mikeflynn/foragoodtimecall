package main

import (
	"strconv"
)

type User struct {
	ID          int64  `json:id`
	Name        string `json:name`
	Image       string `json:image`
	FBID        int64  `json:fbid`
	Email       string `json:email`
	CreatedOn   string `json:created_on`
	PhoneNumber string `json:phonenumber`
	Points      int    `json:points`
}

type UserContest struct {
	UserID      int64  `json:user_id`
	ContestID   int64  `json:contest_id`
	PhoneNumber string `json:phonenumber`
	Points      int    `json:points`
	CreatedOn   string `json:created_on`
}

type UserList struct {
	Items []*User `json:items`
}

var userTable string = "user"

func UserListByContest(contestID int64) (*UserList, error) {
	queryStr := "SELECT u.id, u.name, u.image, u.fbid, u.email, u.created_on, uc.phonenumber, uc.points"
	queryStr += "FROM user_contest AS uc"
	queryStr += "LEFT JOIN user AS u ON (u.id = uc.user_id)"
	queryStr += "WHERE uc.contest_id = ?"

	list := &UserList{}

	results, err := MySQLQueryRows(queryStr, []interface{}{"contest_id": contestID})
	if err != nil {
		return list, err
	}

	for _, row := range results {
		list.Items = append(list.Items, rowToUser(row))
	}

	return list, nil
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

	this = rowToUser(userRow[0])

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

func rowToUser(row map[string]string) *User {
	aUser := &User{}

	if val, ok := row["id"]; ok {
		aUser.ID, _ = strconv.ParseInt(row["id"], 10, 64)
	}

	if val, ok := row["name"]; ok {
		aUser.Name = row["name"]
	}

	if val, ok := row["image"]; ok {
		aUser.Image = row["image"]
	}

	if val, ok := row["fbid"]; ok {
		aUser.FBID, _ = strconv.ParseInt(row["fbid"], 10, 64)
	}

	if val, ok := row["email"]; ok {
		aUser.Email = row["email"]
	}

	if val, ok := row["created_on"]; ok {
		aUser.CreatedOn = row["created_on"]
	}

	if val, ok := row["phonenumber"]; ok {
		aUser.PhoneNumber = row["phonenumber"]
	}

	if val, ok := row["points"]; ok {
		aUser.Points = row["points"]
	}

	return aUser
}
