package repository

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"project/class/domain"
	"project/class/util"
	"time"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{db}
}

func (r *VoucherRepository) Create(voucher domain.Voucher) error {
	return r.db.Create(&voucher).Error
}

func (r *VoucherRepository) All(page uint, limit uint, isActive string, area string, voucherType string) (int64, int, uint, uint, []domain.Voucher, error) {
	var count int64
	r.db.Model(&domain.Voucher{}).Scopes(withArea(area), withExpiry(isActive), hasType(voucherType)).Count(&count)
	pages := int(math.Ceil(float64(count) / float64(limit)))

	var vouchers []domain.Voucher

	result := r.db.Scopes(util.Paginate(page, limit), withArea(area), withExpiry(isActive), hasType(voucherType)).Find(&vouchers)
	return count, pages, page, limit, vouchers, result.Error
}

func (r *VoucherRepository) Get(id uint) (domain.Voucher, error) {
	var voucher domain.Voucher
	err := r.db.First(&voucher, id).Error
	return voucher, err
}

func (r *VoucherRepository) Update(voucher domain.Voucher) error {
	return r.db.Save(&voucher).Error
}

func (r *VoucherRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Voucher{}, id).Error
}

func withArea(area string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("area ILIKE ?", fmt.Sprintf("%%%s%%", area))
	}
}

func withExpiry(isActive string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch isActive {
		case "0":
			return db.Where("expires_at <= ?", time.Now())
		case "1":
			return db.Where("starts_at <= ? AND expires_at >= ?", time.Now(), time.Now())
		default:
			return db
		}
	}
}

func hasType(voucherType string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch voucherType {
		case "0":
			return db.Where("redeem_points IS NULL")
		case "1":
			return db.Where("redeem_points IS NOT NULL")
		default:
			return db
		}
	}
}
