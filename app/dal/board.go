package dal

import (
	"log"
	"strconv"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model        // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt을 포함)
	Title      string `gorm:"column:title;size:20;not null"`
	Content    string `gorm:"column:content;size:100;not null"`
	Img        string `gorm:"column:img;size:100;not null"`
	UserName   string `gorm:"column:user_name;size:10;not null"`
}

func (m *Board) TableName() string {
	return "boards"
}

func (m *Board) GetModels() any {
	return []Board{}
}

func (m *Board) GetModel() ModelInterface {
	return &Board{}
}

func (m *Board) SetID(id string) {
	uint64_id, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		log.Println(err.Error())
		return
	}
	m.ID = uint(uint64_id)
}
