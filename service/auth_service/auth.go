package auth_service

import "Effective/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.Login(a.Username, a.Password)
}