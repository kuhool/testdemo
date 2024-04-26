package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SECRET = "taoshihan"

type UserClaims struct {
	Id         uint      `json:"id"`
	Pid        uint      `json:"pid"`
	Username   string    `json:"username"`
	RoleId     uint      `json:"role_id"`
	CreateTime time.Time `json:"create_time"`
	jwt.StandardClaims
}

func MakeCliamsToken(obj UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, obj)
	tokenString, err := token.SignedString([]byte(SECRET))
	return tokenString, err
}
func ParseCliamsToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
