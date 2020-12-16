package biz

import (
	"main/internal/data"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Demor interface {
	Find() []*data.Demo
}

type DemoBiz struct {
	DB *gorm.DB
}

func NewDemoBiz(db *gorm.DB) *DemoBiz {
	return &DemoBiz{DB: db}
}

func (d *DemoBiz) Find() []*data.Demo {
	var datas []*data.Demo

	d.DB.Find(&datas)
	return datas
}
