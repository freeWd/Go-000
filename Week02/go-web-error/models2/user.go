package models2

import "github.com/pkg/errors"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"username"`
}

func GetUserById(id string) (User, error) {
	var user User
	err := db.QueryRow("SELECT username FROM `user` WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return user, errors.Wrap(err, "read failed")
	}
	return user, nil
}
