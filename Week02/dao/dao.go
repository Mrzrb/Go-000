package dao

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Dao struct {
	*gorm.DB
}

func New() *Dao {
	return &Dao{}
}

type User struct {
	Id   string
	Name string
}

func (d *Dao) GetById(ctx context.Context, id int32) ([]*User, error) {
	var (
		users []*User
		err   error
	)
	if err = d.Where("id=?", id).Find(&users); err != nil {
		return nil, errors.Errorf("error occured when get Users %w", err)
	}

	return users, nil
}
