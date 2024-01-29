package dal

import (
	"log"
	"testAPI/config"

	"gorm.io/gorm"
)

type ORMInterface interface {
	Create(board ModelInterface) error
	ReadAll(model []ModelInterface) ([]ModelInterface, error)
}

type ORM struct {
	DB *gorm.DB
}

func NewORM() (ORMInterface, error) {
	log.Println("dal dal")
	return &ORM{
		DB: config.DB,
	}, nil
}

func (o *ORM) Create(board ModelInterface) error {
	return o.DB.Create(&board).Error
}

func (o *ORM) ReadAll(model []ModelInterface) ([]ModelInterface, error) {
	log.Printf("%T\n", model)
	return model, o.DB.Order("ID desc").Find(&model).Error
}
