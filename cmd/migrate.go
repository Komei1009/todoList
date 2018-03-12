package main

import "github.com/Komei1009/todoList/model"

func main() {
	model.DB.AutoMigrate(&model.User{})
	model.DB.AutoMigrate(&model.Todo{})
}
