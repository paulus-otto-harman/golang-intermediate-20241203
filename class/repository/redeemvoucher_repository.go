package repository

import (
	"gorm.io/gorm"
	"project/class/domain"
)

type RedeemVoucherRepository struct {
	db *gorm.DB
}

func NewRedeemVoucherRepository(db *gorm.DB) *RedeemVoucherRepository {
	return &RedeemVoucherRepository{db}
}

func (r *RedeemVoucherRepository) Create(customer domain.Customer) error {
	return r.db.Model(&domain.Customer{ID: 1}).
		Association("Redemptions").
		Append(&domain.Redemption{VoucherId: 1})
}
