package validation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

// タスク情報
type Todo struct {
	TaskName 	string
	Status		string
}

// タスクのnullチェック
func TodoInputCheck(c *gin.Context)(string, bool) {
	todo := c.PostForm("task")
	fmt.Println(todo)
	if todo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "タスクが入力されていません。",
			})
		return "", false
	}
	return todo, true
}

// タスクへの操作内容チェック
func TodoControllCheck(c *gin.Context)(*Todo, bool){
	todo,_ := TodoInputCheck(c)

	controll := c.PostForm("controll")
	if controll == "change"{
		return &Todo{
			TaskName: todo,
			Status: controll,
		}, true
	}else if controll == "delete"{
		return &Todo{
			TaskName: todo,
			Status: controll,
		}, true
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "操作内容が入力されていません。",
		})
		return nil, false
	}
	return nil, false
}

// 表示するタスク内容チェック
func TodoDisplayCheck(c *gin.Context)(string, bool) {
	mode := c.PostForm("display")
	if mode == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "表示形式が入力されていません。",
		})
		return "", false
	} else if mode != "all" && mode != "active" && mode != "completed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "表示形式の入力が間違っています。",
		})
		return "", false
	}
	return mode, true
}
