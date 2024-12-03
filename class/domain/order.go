package domain

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	CustomerId uint
	VoucherId  uint
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
