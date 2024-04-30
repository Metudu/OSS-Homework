package data

type User struct {
	UserName string `json:"userName"`
	UserAge int `json:"userAge"`
	UserOccupation string `json:"userOccupation"`
}

var UserList = []User{}