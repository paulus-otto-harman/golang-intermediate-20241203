package repository

import "gorm.io/gorm"

type Repository struct {
	Auth          AuthRepository
	RedeemVoucher RedeemVoucherRepository
	Voucher       VoucherRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Auth:          *NewAuthRepository(db),
		RedeemVoucher: *NewRedeemVoucherRepository(db),
		Voucher:       *NewVoucherRepository(db),
	}
}
