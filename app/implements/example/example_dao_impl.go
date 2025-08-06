package example

import (
	"github.com/starter-go/dns-server-starter/app/data/dao"
	"github.com/starter-go/dns-server-starter/app/data/dxo"
	"github.com/starter-go/dns-server-starter/app/data/entity"
	"gorm.io/gorm"
)

// DaoImpl ...
type DaoImpl struct {

	//starter:component
	_as func(dao.ExampleDAO) //starter:as("#")

	Agent dxo.DatabaseAgent //starter:inject("#")
}

func (inst *DaoImpl) _impl() dao.ExampleDAO {
	return inst
}

func (inst *DaoImpl) model() *entity.Example {
	return new(entity.Example)
}

func (inst *DaoImpl) modelList() []*entity.Example {
	return make([]*entity.Example, 0)
}

func (inst *DaoImpl) makeResult(o *entity.Example, res *gorm.DB) (*entity.Example, error) {
	err := res.Error
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Find ...
func (inst *DaoImpl) Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error) {
	m := inst.model()
	db = inst.Agent.DB(db)
	res := db.First(m, id)
	return inst.makeResult(m, res)
}
