package dal

import (
	"log"
	"testAPI/config"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model        // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt을 포함)
	Title      string `gorm:"size:20;not null"`
	Content    string `gorm:"size:100;not null"`
	Img        string `gorm:"size:100;not null"`
	UserName   string `gorm:"size:10;not null"`
}

type BoardORMInterface interface {
	AddBoard(board *Board) error
	GetBoard() (data []Board, err error)
}

type BoardORM struct {
	DB *gorm.DB
}

func NewBoardORM() (BoardORMInterface, error) {

	log.Println("dal dal")
	return &BoardORM{
		DB: config.DB,
	}, nil
}

func (o *BoardORM) AddBoard(board *Board) error {
	return o.DB.Create(&board).Error
}

func (o *BoardORM) GetBoard() (data []Board, err error) {
	return data, o.DB.Order("ID desc").Find(&data).Error
}
