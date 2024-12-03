package service

import "project/class/repository"

type Service struct {
	Auth          AuthService
	RedeemVoucher RedeemVoucherService
	Voucher       VoucherService
}

func NewService(repo repository.Repository) Service {
	return Service{
		Auth:          NewAuthService(repo.Auth),
		RedeemVoucher: NewRedeemVoucherService(repo.RedeemVoucher),
		Voucher:       NewVoucherService(repo.Voucher),
	}
}
