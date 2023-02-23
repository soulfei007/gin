package auth_service

import (
	"gin-learn/models"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() bool {
	return models.CheckAuth(a.Username, a.Password)
}
