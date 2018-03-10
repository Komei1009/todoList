package model

import (
	"github.com/jinzhu/gorm"
)

// タスク情報
type Todo struct {
	gorm.Model
	TaskName 	string	`json:"task_name"`
	Status		string	`json:"status"`
}


// 新規タスク登録
func NewTodo (todo string)error{
	task := Todo {
		TaskName : todo,
		Status: "active",
	}
	return DB.Create(&task).Error
}

// タスク名検索
func ExistTaskName(todo string) bool {
	todos := []Todo{}
	DB.Find(&todos, "task_name = ?", todo)

	return(len(todos) != 0)
}
// タスク完了
func CompletedTodo(todo string)error{
	task := Todo{}
	err := DB.Where("task_name = ?",todo ).First(&task).Error
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
func RemoveTodo(todo string) bool {
	task := Todo{}

	DB.Where("task_name = ?", todo).First(&task)
	DB.Delete(&task)
	return true
}

// タスク表示
func DisplayTodo(mode string) []Todo {
	var allTodo []Todo

	// 全てのタスク
	if mode == "all" {
		DB.Find(&allTodo)
	} else {
		DB.Where("status = ?",mode).Find(&allTodo)
	}
	return allTodo
}