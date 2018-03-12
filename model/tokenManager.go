package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go/request"
	"crypto/sha256"
	"encoding/hex"
)

/* トークン情報 */
type TokenInfo struct {
	Id    int
	Token string
}

var secretKey = secret


//	トークンの作成
func CreateTokenString(name string) (string, error) {

	//	アルゴリズムの指定
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user": name,
		"exp":  time.Now().Add(time.Hour * 48).Unix(),
	}


	//  トークンに対して署名の付与
	return token.SignedString([]byte(secretKey))
}


//	トークンの検証
func AuthorityCheck(c *gin.Context) (string, bool) {
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})

	if err == nil {
		claims := token.Claims.(jwt.MapClaims)
		return claims["user"].(string), true
	} else {
		return "", false
	}
}

// ハッシュ化
func ToHash(pass string) string {
	converted := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(converted[:])
}
