package controller

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"

	"github.com/kataras/iris/v12"
)

type FindTeacherController struct{}

type FindTeacher struct {
	Id          int    `json:"id"`
	TeacherName string `json:"teachername"`
}

func (fc *FindTeacherController) GetTeacher(ctx iris.Context) model.ResponseModel {
	result := []FindTeacher{}
	sql := "select id, teachername from teacher;"
	err := database.DB.Select(&result, sql)
	if err != nil {
		fmt.Print("database.DB.Select error 查找所有课程 id 失败 ❌", err.Error())
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
