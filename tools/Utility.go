package tools

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetToday(format string) (todayString string) {
	today := time.Now()
	todayString = today.Format(format)
	return
}

func EncryptionSha256(data []byte) string {
	hash   	 :=  	sha256.New()
	hash.Write(data)
	Encrypt  := 	base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return Encrypt
}

func GenerateToken(userId uint) (string, error) {
	tokenDuration, err := strconv.Atoi(os.Getenv("token_duration"))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	claim := jwt.MapClaims{}
	claim["authorized"]   =   true
	claim["user_id"]  	  = userId
	claim["exp"] 		  =   time.Now().Add(time.Hour * time.Duration(tokenDuration)).Unix()
	token  :=  jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(os.Getenv("token_secret")))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method:  %v", token.Header["alg"])
		}
		return []byte(os.Getenv("token_secret")), nil
	})
	if err != nil {
		return  err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("token_secret")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}