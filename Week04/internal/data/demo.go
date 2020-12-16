package data

import "github.com/jinzhu/gorm"

type Demo struct {
	gorm.Model
	Blob string
}
