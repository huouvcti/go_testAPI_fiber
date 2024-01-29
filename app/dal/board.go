package dal

import (
	"gorm.io/gorm"
)

type ModelInterface interface {
	TableName()
	GetTable() any
}

type Board struct {
	gorm.Model        // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt을 포함)
	Title      string `gorm:"size:20;not null"`
	Content    string `gorm:"size:100;not null"`
	Img        string `gorm:"size:100;not null"`
	UserName   string `gorm:"size:10;not null"`
}

func (m Board) TableName() {
	println("dddd")
}

func (m Board) GetTable() any {
	return &Board{}
}
