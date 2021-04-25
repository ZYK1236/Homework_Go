/**
 ******************************************************************************
 * File Name          : 登陆
 * Author             : 张宇恺
 * Description        : 根据传入的 username, password 去查表
 ******************************************************************************
 */

package controller

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris/src/database"
	"iris/src/model"
	"time"
)

var mySecret = []byte("secret")

type LoginController struct{}

type User struct {
	UserName string `json:"username"`
	Pwd      string `json:"pwd"`
}

func (lc *LoginController) Get(ctx iris.Context) model.ResponseModel {
	var user = User{}

	sql := "select username,pwd from user where id=1;"
	err := database.DB.Get(&user, sql)

	if err != nil {
		fmt.Println("loginController database.DB.Get error")
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "loginController 数据库操作错误",
		}
	}

	username := ctx.URLParam("username")
	pwd := ctx.URLParam("pwd")

	if user.UserName == username && user.Pwd == pwd {
		token := CreateToken()
		return model.ResponseModel{
			Data: token,
			Code: 1,
			Msg:  "token 生成，25分钟后过期",
		}
	} else {
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "账号/密码不对",
		}
	}
}

func (lc *LoginController) Options(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization")
	ctx.Header("Access-Control-Max-Age", "3600")
}

func CreateToken() string {
	now := time.Now()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": now.Unix(),
		"exp": now.Add(1500 * time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)

	return tokenString
}
