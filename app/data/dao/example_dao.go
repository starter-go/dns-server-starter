package dao

import (
	"github.com/starter-go/dns-server-starter/app/data/dxo"
	"github.com/starter-go/dns-server-starter/app/data/entity"
	"gorm.io/gorm"
)

// ExampleDAO ...
type ExampleDAO interface {
	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)
}
