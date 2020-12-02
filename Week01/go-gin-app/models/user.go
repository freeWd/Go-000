package models

type User struct {
	Model
	Name string `json:"username"`
}

func GetUsers() (users []User) {
	db.Find(&users)
	return
}

func GetUserByID(id string) (user User) {
	db.Where("id = ?", id).First(&user)
	return
}
