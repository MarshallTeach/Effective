package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID			string `gorm:"primary_key" json:"id"`
	Username 	string `json:"username"`
	Password    string `json:"password"`
}

func RegisterUsers(data map[string]interface{}) error {
	user := User{
		ID:       data["id"].(string),
		Username: data["username"].(string),
		Password: data["password"].(string),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func IsUsernameExist(username string) (bool, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.Username != "" {
		return true, nil
	}
	return false, err
}