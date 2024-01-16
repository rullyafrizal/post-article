package model

import (
	"database/sql"
)

type Tabler interface {
	TableName() string
}

type Article struct {
	Id          int64        `gorm:"primaryKey,column:id"`
	Title       string       `gorm:"column:title"`
	Content     string       `gorm:"column:content"`
	Category    string       `gorm:"column:category"`
	Status      string       `gorm:"column:status"`
	CreatedDate sql.NullTime `gorm:"column:created_date"`
	UpdatedDate sql.NullTime `gorm:"column:updated_date"`
}

func (Article) TableName() string {
	return "posts"
}
