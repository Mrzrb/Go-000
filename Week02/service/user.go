package service

import (
	"context"
	"main/dao"
)

type User struct {
}

func New() *User {
	return &User{}
}

func (u *User) GetById(ctx context.Context, id int32) ([]*dao.User, error) {
	var (
		users []*dao.User
		err   error
	)
	if users, err = dao.New().GetById(ctx, id); err != nil {
		return nil, err
	}
	return users, nil
}
