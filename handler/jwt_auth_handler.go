package handler

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// LoginHandler : JWTの発行
func LoginHandler(c *gin.Context) {

	userName := c.Param("user_name")
	passWord := c.Param("pass_word")

	signBytes, err := ioutil.ReadFile("../demo.rsa")
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	if userName == "test" && passWord == "test" {
		// create token
		token := jwt.New(jwt.SigningMethodRS256)

		// set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "test"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString(signKey)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, tokenString)
	}
}
