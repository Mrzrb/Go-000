//+build wireinject

package internal

import (
	"main/internal/biz"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitDemoService(db *gorm.DB) *DemoService {
	wire.Build(NewDemoService, biz.NewDemoBiz)
	return nil
}
