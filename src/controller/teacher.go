/**
 ******************************************************************************
 * File Name          : 查询学生对应老师 Controller
 * Author             : 张宇恺
 * Description        : 根据传入的 stuno 去查对应的老师
 ******************************************************************************
 */

package controller

import (
	"iris/src/database"
	"iris/src/model"
	logMsg "iris/src/utils"

	"github.com/kataras/iris/v12"
)

type TeacherController struct{}

// Teacher 记住这里一定要与 mysql 中的字段一致！！！
type Teacher struct {
	StuNo       int    `json:"stuno"`
	StuName     string `json:"stuname"`
	TeacherName string `json:"teachername"`
}

func (tc *TeacherController) GetTeacher(ctx iris.Context) model.ResponseModel {
	// 数据初始化
	stuno := ctx.URLParam("stuno")
	path := ctx.Path()
	if stuno == "" {
		defer logMsg.LogSuccessMsg(path, "Get")
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "no pass in stuno",
		}
	}

	teacher := Teacher{}

	// 执行 sql 语句
	sql := "select stuno,stuname,teachername from student,teacher where student.teacherid=teacher.id and student.stuno=?;"
	err := database.DB.Get(&teacher, sql, stuno)
	if err != nil {
		logMsg.LogErrorMsg(ctx.Path(), "database.DB.Get")
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "stuno is not found",
		}
	}
	defer logMsg.LogSuccessMsg(path, "Get")

	return model.ResponseModel{
		Data:       teacher,
		Code:       1,
		Msg:        "success",
		TotalCount: 1,
	}
}

func (tc *TeacherController) OptionsTeacher(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization")
	ctx.Header("Access-Control-Max-Age", "3600")
}
