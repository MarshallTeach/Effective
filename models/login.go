package models

import (
	"Effective/pkg/util"
	"github.com/jinzhu/gorm"
)

func Login(username, password string) (bool, error) {
	var user User
	err := db.Where(User{
		Username: username,
		Password: util.EncodeMD5(password),
	}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.Username != "" {
		return true, nil
	}

	return false, nil
}