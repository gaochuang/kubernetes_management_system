package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"kubernetes_management_system/models/user"
	"time"
)

var mySecret = []byte("my_secret_creat")

type Claims struct {
	ID          uint
	Username    string
	Role        string
	BufferTimer int64
	jwt.RegisteredClaims
}

func ReleaseToken(user user.User) (string, error) {
	claims := &Claims{
		ID:       user.ID,
		Username: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //timeout is 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "user token",
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(mySecret)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ParseToken(token string) (*jwt.Token, *Claims, error) {

	fmt.Printf("tocke %s\n", token)
	claims := &Claims{}

	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	fmt.Printf("###tk %v \n", tk)
	return tk, claims, err
}