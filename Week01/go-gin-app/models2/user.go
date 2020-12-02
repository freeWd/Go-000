package models2

import (
	"fmt"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"username"`
}

func GetUserById(id string) User {
	rows, err := db.Query("SELECT id, username FROM `user` WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows, "<--------")
	return User{}
}
