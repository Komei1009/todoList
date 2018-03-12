package validation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Pass string
}


/*
	ユーザー情報入力チェック
*/
func ToUser(c *gin.Context) (*User, bool) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "nameが未入力です",
		})
		return nil, false
	}
	pass := c.PostForm("password")
	if pass == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "passが未入力です",
		})
		return nil, false
	}

	return &User{
		Name: name,
		Pass: pass,
	}, true
}

