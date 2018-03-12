package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Komei1009/todoList/model"
	"github.com/Komei1009/todoList/controller/validation"
	"net/http"
)

// 新規ユーザー作成
func CreateUserController(c *gin.Context) {
	// リクエストパラメーター取得
	user, ok := validation.ToUser(c)
	if !ok {
		return
	}
	// 作成済みユーザーか？
	if model.ExistUserByName(user.Name) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "登録済みのユーザーネームです",
		})
		return
	}

	user.Pass = model.ToHash(user.Pass)

	// DBinsert
	err := model.CreateUser(user.Name, user.Pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "データベースエラー",
		})
		return
	}

	token, err := model.CreateTokenString(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "アクセストークンを作成できませんでした",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// 既存ユーザーログイン
func LoginController(c *gin.Context) {
	// リクエストパラメータチェック
	user, ok := validation.ToUser(c)
	if !ok {
		return
	}

	user.Pass = model.ToHash(user.Pass)

	// ログインチェック
	if !model.CheckLogin(user.Name, user.Pass) {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "ユーザー名またはパスワードが間違っています",
		})
		return
	}

	// トークンを生成
	token, err := model.CreateTokenString(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}