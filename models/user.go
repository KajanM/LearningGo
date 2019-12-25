package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.Id == 0 {
		return User{}, errors.New("new User must not include id")
	}
	user.Id = nextId
	nextId++
	users = append(users, &user)
	return user, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("user with Id '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	if u.Id == 0 {
		return User{}, errors.New("id is required")
	}
	for i, candidate := range users {
		if candidate.Id == id {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("could not find a user with id '%v'", id)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user with id '%v' not found", id)
}
