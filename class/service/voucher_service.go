package service

import (
	"project/class/domain"
	"project/class/repository"
)

type VoucherService interface {
	Create(voucher domain.Voucher) error
	All(page uint, limit uint, isActive string, area string, voucherType string) (int64, int, uint, uint, []domain.Voucher, error)
	Get(id uint) (domain.Voucher, error)
	Update(voucher domain.Voucher) error
	Delete(id uint) error
}

type voucherService struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo}
}

func (s *voucherService) Create(voucher domain.Voucher) error {
	return s.repo.Create(voucher)
}

func (s *voucherService) All(currentPage uint, itemsPerPage uint, isActive string, area string, voucherType string) (int64, int, uint, uint, []domain.Voucher, error) {
	if currentPage == 0 {
		currentPage = 1
	}

	if itemsPerPage == 0 {
		itemsPerPage = 6
	}

	return s.repo.All(currentPage, itemsPerPage, isActive, area, voucherType)
}

func (s *voucherService) Get(id uint) (domain.Voucher, error) {
	return s.repo.Get(id)
}

func (s *voucherService) Update(voucher domain.Voucher) error {
	return s.repo.Update(voucher)
}

func (s *voucherService) Delete(id uint) error {
	return s.repo.Delete(id)
}
