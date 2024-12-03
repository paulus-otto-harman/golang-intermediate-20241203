package domain

type DiscountVoucher struct {
	ID        uint     `gorm:"primaryKey;autoIncrement" json:"-"`
	VoucherID uint     `json:"-"`
	Rate      *float32 `gorm:"precision:5;scale:2" json:"rate,omitempty"`
	Amount    *uint32  `gorm:"precision:4" json:"amount,omitempty"`
	Condition `gorm:"embedded" json:"terms_and_condition"`
}
