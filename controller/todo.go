package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Komei1009/todoList/controller/validation"
	"github.com/Komei1009/todoList/model"
	"net/http"
)

// タスク新規追加
func TodoAdd(c *gin.Context){
	userName, ok := model.AuthorityCheck(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "不正アクセスです。",
		})
		return
	}
	// リクエストパラメータ取得
	todo, ok := validation.TodoInputCheck(c)
	if !ok{
		return
	}

	// タスク名の重複防止
	if model.ExistTaskName(userName,todo) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "登録済みのタスクです。",
		})
		return
	}
	// DBにタスク追加
	err := model.NewTodo(userName, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "データベースエラー",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "新規タスクが登録されました。",
	})
}

// タスク完了or削除チェック
func TodoControll(c *gin.Context){
	user, ok := model.AuthorityCheck(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "不正アクセスです。",
		})
		return
	}
	// リクエストパラメータ取得
	controll, ok := validation.TodoControllCheck(c)
	if !ok {
		return
	}

	// タスクがない時"”
	if !model.ExistTaskName(user, controll.TaskName) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "タスクが見つかりませんでした",
		})
		return
	}

	// タスク完了or削除の処理
	if controll.Status == "change" {
		model.CompletedTodo(user, controll.TaskName)
		c.JSON(http.StatusOK, gin.H{
			"ok": "タスクの状態を変更しました。",
		})
	} else {
		model.RemoveTodo(user, controll.TaskName)
		c.JSON(http.StatusOK, gin.H{
			"ok": "タスクを削除しました。",
		})
	}

}

// 表示するTodo
func TodoDisplay(c *gin.Context){
	user, ok := model.AuthorityCheck(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "不正アクセスです。",
		})
		return
	}
	// リクエストパラメータ取得
	mode, ok := validation.TodoDisplayCheck(c)
	if !ok{
		return
	}

	// タスク表示
	c.JSON(http.StatusOK, gin.H{
		"todos" : model.DisplayTodo(user, mode),
	})
}
