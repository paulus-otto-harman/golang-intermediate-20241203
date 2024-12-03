package domain

import (
	"gorm.io/gorm"
	"project/class/util"
	"time"
)

type Voucher struct {
	ID           uint                 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string               `gorm:"size:20;unique" json:"name"`
	Code         string               `gorm:"size:15;unique" json:"code"`
	RedeemPoints *uint                `json:"redeem_points"`
	FreeShipping *FreeShippingVoucher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"free_shipping,omitempty"`
	Discount     *DiscountVoucher     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"discount,omitempty"`
	StartsAt     time.Time            `json:"starts_at"`
	ExpiresAt    time.Time            `json:"expires_at"`
	Area         string               `gorm:"size:15" json:"area"`
	CreatedAt    time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time            `json:"-"`
	DeletedAt    gorm.DeletedAt       `json:"-"`
}

func VoucherSeed() []Voucher {
	return []Voucher{
		{
			Name: "afooj10q0", Code: "code001", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoaj10q5", Code: "code002", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afooj25q5", Code: "code003", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoaj25q0", Code: "code004", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afooj10c5", Code: "code005", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoaj10c5", Code: "code006", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afooj25c5", Code: "code007", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoaj25c5", Code: "code008", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoojt10q5", Code: "code009", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoajt10q5", Code: "code010", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoojt25q5", Code: "code011", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoajt25q0", Code: "code012", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoojt10c5", Code: "code013", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoajt10c5", Code: "code014", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoojt25c5", Code: "code015", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoajt25c5", Code: "code016", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		// //
		{
			Name: "xfooj10q5", Code: "code017", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoaj10q5", Code: "code018", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfooj25q5", Code: "code019", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoaj25q0", Code: "code020", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfooj10c5", Code: "code021", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoaj10c5", Code: "code022", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfooj25c5", Code: "code023", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoaj25c5", Code: "code024", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoojt10q5", Code: "code025", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoajt10q5", Code: "code026", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoojt25q5", Code: "code027", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoajt25q0", Code: "code028", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoojt10c5", Code: "code029", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoajt10c5", Code: "code030", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoojt25c5", Code: "code031", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoajt25c5", Code: "code032", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		// // //
		{
			Name: "adro10j10q5", Code: "code033", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10j10q5", Code: "code034", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10j25q5", Code: "code035", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10j25q0", Code: "code036", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10j10c5", Code: "code037", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10j10c5", Code: "code038", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro10j25c5", Code: "code039", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10j25c5", Code: "code040", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro10jt10q5", Code: "code041", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10jt10q5", Code: "code042", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10jt25q5", Code: "code043", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10jt25q0", Code: "code044", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10jt10c5", Code: "code045", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10jt10c5", Code: "code046", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro10jt25c5", Code: "code047", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10jt25c5", Code: "code048", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10j10q5", Code: "code049", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10j10q5", Code: "code050", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10j25q5", Code: "code051", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10j25q0", Code: "code052", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10j10c5", Code: "code053", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10j10c5", Code: "code054", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10j25c5", Code: "code055", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10j25c5", Code: "code056", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10jt10q5", Code: "code057", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10jt10q5", Code: "code058", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10jt25q5", Code: "code059", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10jt25q0", Code: "code060", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10jt10c5", Code: "code061", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10jt10c5", Code: "code062", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10jt25c5", Code: "code063", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10jt25c5", Code: "code064", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		//// // // //
		{
			Name: "adro25j10q5", Code: "code065", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25j10q5", Code: "code066", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25j25q5", Code: "code067", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25j25q0", Code: "code068", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25j10c5", Code: "code069", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25j10c5", Code: "code070", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro25j25c5", Code: "code071", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25j25c5", Code: "code072", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro25jt10q5", Code: "code073", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25jt10q5", Code: "code074", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25jt25q5", Code: "code075", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25jt25q0", Code: "code076", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25jt10c5", Code: "code077", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25jt10c5", Code: "code078", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro25jt25c5", Code: "code079", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25jt25c5", Code: "code080", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25j10q5", Code: "code081", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25j10q5", Code: "code082", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25j25q5", Code: "code083", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25j25q0", Code: "code084", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25j10c5", Code: "code085", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25j10c5", Code: "code086", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25j25c5", Code: "code087", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25j25c5", Code: "code088", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25jt10q5", Code: "code089", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25jt10q5", Code: "code090", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25jt25q5", Code: "code091", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25jt25q0", Code: "code092", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25jt10c5", Code: "code093", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25jt10c5", Code: "code094", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25jt25c5", Code: "code095", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25jt25c5", Code: "code096", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		//// // // // //
		{
			Name: "adao10j100q5", Code: "code097", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10j100q5", Code: "code098", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10j250q5", Code: "code099", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10j250q0", Code: "code100", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10j100c5", Code: "code101", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10j100c5", Code: "code102", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao10j250c5", Code: "code103", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10j250c5", Code: "code104", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao10jt100q5", Code: "code105", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10jt100q5", Code: "code106", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10jt250q5", Code: "code107", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10jt250q0", Code: "code108", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10jt100c5", Code: "code109", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10jt100c5", Code: "code110", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao10jt250c5", Code: "code111", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10jt250c5", Code: "code112", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10j100q5", Code: "code113", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10j100q5", Code: "code114", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10j250q5", Code: "code115", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10j250q0", Code: "code116", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10j100c5", Code: "code117", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10j100c5", Code: "code118", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10j250c5", Code: "code119", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10j250c5", Code: "code120", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10jt100q5", Code: "code121", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10jt100q5", Code: "code122", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10jt250q5", Code: "code123", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10jt250q0", Code: "code124", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10jt100c5", Code: "code125", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10jt100c5", Code: "code126", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10jt250c5", Code: "code127", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10jt250c5", Code: "code128", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		// // // // // // //
		{
			Name: "adao25j200q5", Code: "code129", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25j200q5", Code: "code130", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25j400q5", Code: "code131", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25j400q0", Code: "code132", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25j200c5", Code: "code133", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25j200c5", Code: "code134", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao25j400c5", Code: "code135", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25j400c5", Code: "code136", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao25jt200q5", Code: "code137", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25jt200q5", Code: "code138", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25jt400q5", Code: "code139", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25jt400q0", Code: "code140", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25jt200c5", Code: "code141", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25jt200c5", Code: "code142", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao25jt400c5", Code: "code143", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25jt400c5", Code: "code144", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25j200q5", Code: "code145", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25j200q5", Code: "code146", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25j400q5", Code: "code147", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25j400q0", Code: "code148", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25j200c5", Code: "code149", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25j200c5", Code: "code150", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25j400c5", Code: "code151", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25j400c5", Code: "code152", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25jt200q5", Code: "code153", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25jt200q5", Code: "code154", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25jt400q5", Code: "code155", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25jt400q0", Code: "code156", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25jt200c5", Code: "code157", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25jt200c5", Code: "code158", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25jt400c5", Code: "code159", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25jt400c5", Code: "code160", RedeemPoints: util.Ptr(uint(50)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		// // // // // // // //
		{
			Name: "afoob10q0", Code: "code161", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Bali",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoaj10q1", Code: "code162", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afooj25q1", Code: "code163", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoab25q0", Code: "code164", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Bali",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afooj10c1", Code: "code165", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoaj10c1", Code: "code166", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afooj25c1", Code: "code167", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoaj25c1", Code: "code168", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoojt10q1", Code: "code169", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoajt10q1", Code: "code170", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoojt25q1", Code: "code171", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoap25q0", Code: "code172", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Papua",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "afoojt10c1", Code: "code173", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoajt10c1", Code: "code174", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoojt25c1", Code: "code175", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "afoajt25c1", Code: "code176", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfooj10q1", Code: "code177", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoaj10q1", Code: "code178", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfooj25q1", Code: "code179", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoab25q0", Code: "code180", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Bali",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfooj10c1", Code: "code181", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoaj10c1", Code: "code182", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfooj25c1", Code: "code183", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoaj25c1", Code: "code184", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoojt10q1", Code: "code185", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoajt10q1", Code: "code186", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoojt25q1", Code: "code187", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoap25q0", Code: "code188", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Papua",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xfoojt10c1", Code: "code189", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoajt10c1", Code: "code190", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoojt25c1", Code: "code191", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xfoajt25c1", Code: "code192", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			FreeShipping: util.Ptr(FreeShippingVoucher{Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		// // //
		{
			Name: "adro10j10q1", Code: "code193", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10j10q1", Code: "code194", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10j25q1", Code: "code195", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10b25q0", Code: "code196", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10j10c1", Code: "code197", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10j10c1", Code: "code198", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro10j25c1", Code: "code199", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10j25c1", Code: "code200", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro10jt10q1", Code: "code201", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10jt10q1", Code: "code202", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10jt25q1", Code: "code203", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra10p25q0", Code: "code204", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro10jt10c1", Code: "code205", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10jt10c1", Code: "code206", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro10jt25c1", Code: "code207", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra10jt25c1", Code: "code208", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10j10q1", Code: "code209", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10j10q1", Code: "code210", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10j25q1", Code: "code211", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10b25q0", Code: "code212", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10j10c1", Code: "code213", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10j10c1", Code: "code214", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10j25c1", Code: "code215", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10j25c1", Code: "code216", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10jt10q1", Code: "code217", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10jt10q1", Code: "code218", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10jt25q1", Code: "code219", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra10p25q0", Code: "code220", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro10jt10c1", Code: "code221", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10jt10c1", Code: "code222", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro10jt25c1", Code: "code223", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra10jt25c1", Code: "code224", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(10)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		//// // // //
		{
			Name: "adro25j10q1", Code: "code225", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25j10q1", Code: "code226", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25j25q1", Code: "code227", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25b25q0", Code: "code228", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25j10c1", Code: "code229", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25j10c1", Code: "code230", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro25j25c1", Code: "code231", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25j25c1", Code: "code232", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro25jt10q1", Code: "code233", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25jt10q1", Code: "code234", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25jt25q1", Code: "code235", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adra25p25q0", Code: "code236", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "adro25jt10c1", Code: "code237", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25jt10c1", Code: "code238", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "adro25jt25c1", Code: "code239", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "adra25jt25c1", Code: "code240", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25j10q1", Code: "code241", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25j10q1", Code: "code242", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25j25q1", Code: "code243", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25b25q0", Code: "code244", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25j10c1", Code: "code245", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25j10c1", Code: "code246", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25j25c1", Code: "code247", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25j25c1", Code: "code248", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25jt10q1", Code: "code249", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25jt10q1", Code: "code250", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25jt25q1", Code: "code251", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdra25p25q0", Code: "code252", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdro25jt10c1", Code: "code253", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25jt10c1", Code: "code254", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 10, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdro25jt25c1", Code: "code255", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: true, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdra25jt25c1", Code: "code256", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Rate: util.Ptr(float32(25)), Condition: Condition{IsEither: false, MinPurchase: 25, PaymentMethod: "cod"}}),
		},
		//// // // // //
		{
			Name: "adao10j100q1", Code: "code257", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10j100q1", Code: "code258", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10j250q1", Code: "code259", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10b250q0", Code: "code260", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10j100c1", Code: "code261", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10j100c1", Code: "code262", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao10j250c1", Code: "code263", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10j250c1", Code: "code264", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao10jt100q1", Code: "code265", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10jt100q1", Code: "code266", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10jt250q1", Code: "code267", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa10p250q0", Code: "code268", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao10jt100c1", Code: "code269", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10jt100c1", Code: "code270", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao10jt250c1", Code: "code271", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa10jt250c1", Code: "code272", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10j100q1", Code: "code273", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10j100q1", Code: "code274", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10j250q1", Code: "code275", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10b250q0", Code: "code276", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10j100c1", Code: "code277", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10j100c1", Code: "code278", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10j250c1", Code: "code279", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10j250c1", Code: "code280", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10jt100q1", Code: "code281", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10jt100q1", Code: "code282", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10jt250q1", Code: "code283", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa10p250q0", Code: "code284", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao10jt100c1", Code: "code285", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10jt100c1", Code: "code286", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 100, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao10jt250c1", Code: "code287", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: true, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa10jt250c1", Code: "code288", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(10)), Condition: Condition{IsEither: false, MinPurchase: 250, PaymentMethod: "cod"}}),
		},
		// // // // // // //
		{
			Name: "adao25j200q1", Code: "code289", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25j200q1", Code: "code290", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25j400q1", Code: "code291", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25b400q0", Code: "code292", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25j200c1", Code: "code293", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25j200c1", Code: "code294", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao25j400c1", Code: "code295", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25j400c1", Code: "code296", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao25jt200q1", Code: "code297", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25jt200q1", Code: "code298", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25jt400q1", Code: "code299", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adaa25p400q0", Code: "code300", RedeemPoints: nil,
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "adao25jt200c1", Code: "code301", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25jt200c1", Code: "code302", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "adao25jt400c1", Code: "code303", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "adaa25jt400c1", Code: "code304", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-11-15"), ExpiresAt: util.Date("2024-12-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25j200q1", Code: "code305", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25j200q1", Code: "code306", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25j400q1", Code: "code307", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25b400q0", Code: "code308", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Bali",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25j200c1", Code: "code309", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25j200c1", Code: "code310", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25j400c1", Code: "code311", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25j400c1", Code: "code312", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "DKI Jakarta",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25jt200q1", Code: "code313", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25jt200q1", Code: "code314", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25jt400q1", Code: "code315", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdaa25p400q0", Code: "code316", RedeemPoints: nil,
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Papua",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "qris"}}),
		},
		{
			Name: "xdao25jt200c1", Code: "code317", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25jt200c1", Code: "code318", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 200, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdao25jt400c1", Code: "code319", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: true, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
		{
			Name: "xdaa25jt400c1", Code: "code320", RedeemPoints: util.Ptr(uint(100)),
			StartsAt: util.Date("2024-10-15"), ExpiresAt: util.Date("2024-11-15"), Area: "Jawa Timur",
			Discount: util.Ptr(DiscountVoucher{Amount: util.Ptr(uint32(25)), Condition: Condition{IsEither: false, MinPurchase: 400, PaymentMethod: "cod"}}),
		},
	}
}
