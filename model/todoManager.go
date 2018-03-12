package model

import (
	"github.com/jinzhu/gorm"
)

// タスク情報
type Todo struct {
	gorm.Model
	User		string	`json:"user"`
	TaskName 	string	`json:"task_name"`
	Status		string	`json:"status"`
}


// 新規タスク登録
func NewTodo (user string, todo string)error{
	task := Todo {
		User	 : user,
		TaskName : todo,
		Status: "active",
	}
	return DB.Create(&task).Error
}

// タスク名検索
func ExistTaskName(user string, todo string) bool {
	todos := []Todo{}
	DB.Find(&todos, "user = ? and task_name = ?", user,todo)

	return(len(todos) != 0)
}
// タスク完了
func CompletedTodo(user string, todo string)error{
	task := Todo{}
	err := DB.Where("user = ? and task_name = ?",user, todo ).First(&task).Error
	if err != nil {
		panic(err)
		return err
	}
	if task.Status == "active" {
		err = DB.Model(&task).Update(&task).Update("status", "completed").Error
	} else {
		err = DB.Model(&task).Update(&task).Update("status", "active").Error
	}

	return err
}

// タスク削除
func RemoveTodo(user string, todo string) bool {
	task := Todo{}

	DB.Where("user = ? and task_name = ?",user, todo).First(&task)
	DB.Delete(&task)
	return true
}

// タスク表示
func DisplayTodo(user string, mode string) []Todo {
	var displayTodo []Todo

	// 全てのタスク
	if mode == "all" {
		DB.Where("user = ?",user).Find(&displayTodo)
	} else {
		DB.Where("user = ? and status = ?",user, mode).Find(&displayTodo)
	}
	return displayTodo
}