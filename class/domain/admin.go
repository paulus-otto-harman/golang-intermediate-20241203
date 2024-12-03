package domain

type Admin struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	UserID uint
	User   User
}

func AdminSeed() []Admin {
	return []Admin{
		{
			Name: "Super Admin",
			User: User{
				Username: "admin",
				Password: "password",
			},
		},
	}
}
