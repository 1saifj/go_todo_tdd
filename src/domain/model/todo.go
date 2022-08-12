package model

type Todo struct {
	Base       `gorm:"embedded"`
	Title      string `json:"title" gorm:"not null"`
	Content    string `json:"content" gorm:"not null"`
	IsComplete bool   `json:"is_complete" gorm:"not null;default:false"`
}
