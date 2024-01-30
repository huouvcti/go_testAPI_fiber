package dal

import (
	"log"
	"testAPI/config"

	"gorm.io/gorm"
)

type ModelInterface interface {
	TableName() string
	GetModels() any
}

type ORMInterface interface {
	Create(board ModelInterface) error
	ReadAll(model ModelInterface) (any, error)
}

type ORM struct {
	DB *gorm.DB
}

func NewORM() (ORMInterface, error) {
	log.Println("New ORM")
	return &ORM{
		DB: config.DB,
	}, nil
}

func (o *ORM) Create(board ModelInterface) error {
	return o.DB.Create(&board).Error
}

func (o *ORM) ReadAll(model ModelInterface) (any, error) {
	results := model.GetModels()

	// o.DB.Table(model.TableName()).Order("ID desc").Find(&results).Error 또는
	return results, o.DB.Order("ID desc").Find(&results).Error
}
