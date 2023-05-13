package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
