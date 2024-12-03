package domain

type FreeShippingVoucher struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"-"`
	VoucherID uint `json:"-"`
	Condition `gorm:"embedded" json:"terms_and_condition"`
}
