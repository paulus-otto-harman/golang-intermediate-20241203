package domain

type Condition struct {
	IsEither      bool   `json:"is_either"`
	MinPurchase   uint32 `gorm:"precision:4" json:"min_purchase"`
	PaymentMethod string `gorm:"size:15" json:"payment_method"`
}
