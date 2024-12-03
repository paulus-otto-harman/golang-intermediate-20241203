package domain

import (
	"gorm.io/gorm"
	"time"
)

type Redemption struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	CustomerId uint
	VoucherId  uint
	OrderId    *uint
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
