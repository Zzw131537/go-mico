package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 鉴权
var jwtSecret = []byte("TodoList")

type Claim struct {
	Id uint32 `json:"id"`
	jwt.StandardClaims
}

// 签发用户token

func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	//	fmt.Println("expireTime :", expireTime)
	claims := Claim{
		Id: uint32(id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todoList",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(tokenClaims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("签发token: ", token)
	return token, err
}

// 验证用户token

func ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claim); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
