package handler

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//JSON : 受け取るJSON
type JSON struct {
	ID string `json:"id" binding:"required"`
}

// LoginHandler : JWTの発行
func LoginHandler(c *gin.Context) {

	keyData, err := ioutil.ReadFile("./secret.key")
	if err != nil {
		panic(err)
	}

	var json JSON
	c.BindJSON(&json)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat": time.Now().Unix(),
		"id":  json.ID,
	})

	tokenString, err := token.SignedString([]byte(keyData))
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"status": "OK",
		"token":  tokenString,
	})
	return
}

// RequireTokenAuthenticationHandler : Tokenの検証
func RequireTokenAuthenticationHandler(c *gin.Context) {

	keyData, err := ioutil.ReadFile("./secret.key")
	if err != nil {
		panic(err)
	}

	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(keyData)
		return b, nil
	})

	if err == nil {
		claims := token.Claims.(jwt.MapClaims)
		msg := fmt.Sprintf("こんにちは、「 %s 」さん", claims["id"])
		c.JSON(200, gin.H{"message": msg})
	} else {
		c.JSON(401, gin.H{"error": fmt.Sprint(err)})
	}
}
