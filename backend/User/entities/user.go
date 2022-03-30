package entities

import (
	"time"

	"github.com/carloshcamilo/ego/config"
)

type User struct {
	Id        int        `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

func (u *User) CreateUser() (*User, error) {
	config.DB.Find(&u, "Email = ?", u.Email)
	config.DB.Create(&u)
	if err := config.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) FindUser() error {
	config.DB.Find(&u, u.Id)
	if err := config.DB.Error; err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUser() error {
	config.DB.Delete(&u, u.Id)
	if err := config.DB.Error; err != nil {
		return err
	}

	return nil
}
