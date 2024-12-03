package domain

type User struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	Username string `gorm:"unique" example:"admin"`
	Password string `example:"password"`
}
