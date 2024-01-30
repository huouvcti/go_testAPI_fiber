package dal

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model        // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt을 포함)
	Title      string `gorm:"size:20;not null"`
	Content    string `gorm:"size:100;not null"`
}

func (m *List) TableName() string {
	return "lists"
}

func (m *List) GetModels() any {
	return []List{}
}
