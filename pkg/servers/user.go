package server

import (
	"errors"
	"golang.org/x/net/context"
)

type User struct {
	ID   int64
	Name string
}

func (User) Get(ctx context, id int64) (user User, err error) {
	u := user{
		ID:   int64(100),
		Name: "Fred",
	}
	return u, nil

}
