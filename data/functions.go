package data

import "errors"

func CreateUser(user []User) {
	UserList = append(UserList, user...)
}

func ReadUser(index int) (User, error) {
	if index < 0 || index > len(UserList) - 1 {
		return User{}, errors.New("invalid index")
	}
	return UserList[index], nil
}

func UpdateUser(index int, user User) {
	UserList[index] = user	
}

func DeleteUser(index int) {
	UserList = append(UserList[: index], UserList[index + 1:]...)
}