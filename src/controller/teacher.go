package controller

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"

	"github.com/kataras/iris/v12"
)

type TeacherController struct{}

// 记住这里一定要与 mysql 中的字段一致！！！
type Teacher struct {
	StuNo       int    `json:"stuno"`
	StuName     string `json:"stuname"`
	TeacherName string `json:"teachername"`
}

func (cc *TeacherController) GetTeacher(ctx iris.Context) model.ResponseModel {
	stuno := ctx.URLParam("stuno")
	if stuno == "" {
		defer fmt.Println("path:/student/teacher ----> GET ✅")
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "no pass in stuno",
		}
	}

	teacher := Teacher{}

	sql := "select stuno,stuname,teachername from student,teacher where student.teacherid=teacher.id and student.stuno=?;"
	err := database.DB.Get(&teacher, sql, stuno)
	if err != nil {
		fmt.Println("database.DB.Get error, stuno is not found❌")
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "stuno is not found",
		}
	}
	defer fmt.Println("path:/student/Teacher ----> GET ✅")

	return model.ResponseModel{
		Data:       teacher,
		Code:       1,
		Msg:        "success",
		TotalCount: 1,
	}
}
