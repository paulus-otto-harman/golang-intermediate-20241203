package database

import (
	"gorm.io/gorm"
	"project/class/domain"
)

func Migrate(db *gorm.DB) error {
	var err error

	if err = dropTables(db); err != nil {
		return err
	}

	if err = setupJoinTables(db); err != nil {
		return err
	}

	return db.AutoMigrate(
		&domain.Voucher{},
		&domain.FreeShippingVoucher{},
		&domain.DiscountVoucher{},
		&domain.Customer{},
		&domain.User{},
		&domain.Admin{},
	)

}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&domain.FreeShippingVoucher{},
		&domain.DiscountVoucher{},
		&domain.Order{},
		&domain.Redemption{},
		&domain.Voucher{},
		&domain.User{},
		&domain.Customer{},
		&domain.Admin{},
	)
}

func setupJoinTables(db *gorm.DB) error {
	var err error
	if err = db.SetupJoinTable(&domain.Customer{}, "Redemptions", &domain.Redemption{}); err != nil {
		return err
	}

	if err = db.SetupJoinTable(&domain.Customer{}, "Vouchers", &domain.Order{}); err != nil {
		return err
	}
	return nil
}
