/**
 ******************************************************************************
 * File Name          : 查询所有老师 Controller
 * Author             : 张宇恺
 * Description        : 查表，给出表中所有老师 id 以及名字
 ******************************************************************************
 */

package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"iris/src/database"
	"iris/src/model"
)

type FindTeacherController struct{}

type FindTeacher struct {
	Id          int    `json:"id"`
	TeacherName string `json:"teachername"`
}

func (fc *FindTeacherController) GetTeacher() model.ResponseModel {
	result := []FindTeacher{}
	sql := "select id, teachername from teacher;"
	err := database.DB.Select(&result, sql)
	if err != nil {
		fmt.Print("database.DB.Select error 查找所有老师 id 失败 ❌", err.Error())
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "查找所有课程 id 失败",
		}
	}

	defer fmt.Println("path:/find/teacher ----> GET ✅")

	return model.ResponseModel{
		Data:       result,
		Code:       1,
		Msg:        "success",
		TotalCount: len(result),
	}
}

func (fc *FindTeacherController) OptionsTeacher(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization")
	ctx.Header("Access-Control-Max-Age", "3600")
}
