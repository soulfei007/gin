package models

type Auth struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)

	return auth.ID > 0
}
