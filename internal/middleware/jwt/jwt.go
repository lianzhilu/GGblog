package jwt

import (
	errmsg "GGblog/internal/errormsg"
	"GGblog/internal/setting"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte(setting.AppConf.JwtKey)

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return JwtKey, nil
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成token
func SetToken(username string) (string, int) {
	c := MyClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			Issuer:    "ggblog",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// 验证token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, keyFunc)
	claims, _ := token.Claims.(*MyClaims)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.New(errmsg.GetErrorMessage(errmsg.ERROR_INVALID_TOKEN))
	}
	return claims, nil
}

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var code int
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errmsg.GetErrorMessage(code),
			})
			ctx.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errmsg.GetErrorMessage(code),
			})
			ctx.Abort()
			return
		}

		parsedToken, err := ParseToken(checkToken[1])
		if err != nil {
			code = errmsg.ERROR_TOKEN_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errmsg.GetErrorMessage(code),
			})
			ctx.Abort()
			return
		}
		ctx.Set("username", parsedToken.Username)
		ctx.Next()
	}
}
