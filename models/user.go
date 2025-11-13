package models

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   int
	Email string
}
