package core

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MineClaims struct {
	Id    uint `json:"id"`
	IsVip bool `json:"is_vip"`
	jwt.StandardClaims
}

// GetToken 获取 token
func GetToken(id uint, wechatOpenId string, IsVip bool) (string, error) {
	mySigningKey := []byte("jewelry")

	// Create the Claims
	claims := MineClaims{
		Id:    id,
		IsVip: IsVip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 7 * time.Hour * time.Duration(1)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func AuthToken(tokenString string) (*MineClaims, bool) {
	if tokenString == "" {
		return nil, false
	}

	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &MineClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("jewelry"), nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(*MineClaims); ok && token.Valid {

		return claims, true
	}
	return nil, false
}

func GetLoginUser(tokenString string) uint {

	var loginUserId uint = 0
	if tokenString != "" {
		customClaims, _ := AuthToken(tokenString)
		loginUserId = customClaims.Id
	}

	return loginUserId
}
