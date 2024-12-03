package domain

type Customer struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `json:"name"`
	RewardPoints uint      `json:"reward_points"`
	Redemptions  []Voucher `gorm:"many2many:redemptions;" json:"redemptions"`
	Vouchers     []Voucher `gorm:"many2many:orders;" json:"vouchers"`
}

func CustomerSeed() []Customer {
	return []Customer{
		{Name: "Customer Satu", RewardPoints: 10},
		{Name: "Customer Dua", RewardPoints: 100},
		{Name: "Customer Tiga", RewardPoints: 50},
		{Name: "Customer Empat", RewardPoints: 30},
		{Name: "Customer Lima", RewardPoints: 25},
		{Name: "Customer Enam", RewardPoints: 100},
		{Name: "Customer Tujuh", RewardPoints: 500},
		{Name: "Customer Delapan", RewardPoints: 25},
		{Name: "Customer Sembilan", RewardPoints: 75},
		{Name: "Customer Sepuluh", RewardPoints: 200},
	}
}
