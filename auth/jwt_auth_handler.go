package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}
