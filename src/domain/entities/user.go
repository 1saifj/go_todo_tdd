package entities

type User struct {
	Base     `gorm:"embedded"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
