package model

import "github.com/jinzhu/gorm"

// ユーザー情報
type User struct {
	gorm.Model
	Name     string
	Password string
}

// ユーザー情報登録
func CreateUser(name string, password string) error {
	user := User{
		Name:     name,
		Password: password,
	}
	return DB.Create(&user).Error
}


// DBからユーザーネームを探す
func ExistUserByName(name string) bool {

	users := []User{}
	DB.Find(&users, "name = ?", name)

	return (len(users) != 0)
}

// ログインチェック
func CheckLogin(name string, password string) bool {

	users := []User{}
	DB.Find(&users, "name = ?", name)
	if len(users) != 0 {
		if password == users[0].Password {
			return true
		}
	}
	return false
}