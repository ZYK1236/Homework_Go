/**
  ******************************************************************************
  * File Name          : 查询学生对应课程 Controller
  * Author             : 张宇恺
  * Description        : 根据传入的 stuno 去查对应的课程
  ******************************************************************************
*/

package controller

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"

	"github.com/kataras/iris/v12"
)

type CourseController struct{}

// 记住这里一定要与 mysql 中的字段一致！！！
type Course struct {
	StuNo      int    `json:"stuno"`
	StuName    string `json:"stuname"`
	CourseName string `json:"coursename"`
}

func (cc *CourseController) GetCourse(ctx iris.Context) model.ResponseModel {
	stuno := ctx.URLParam("stuno")
	if stuno == "" {
		defer fmt.Println("path:/student/course ----> GET ✅")

		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "no pass in stuno",
		}
	}

	course := Course{}

	sql := "select stuno,stuname,coursename from student,cource where student.courseid=cource.id and student.stuno=?;"
	err := database.DB.Get(&course, sql, stuno)
	if err != nil {
		fmt.Println("database.DB.Get error, stuno is not found❌")
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "stuno is not found",
		}
	}
	defer fmt.Println("path:/student/course ----> GET ✅")

	return model.ResponseModel{
		Data: course,
		Code: 1,
		Msg:  "success",
		TotalCount: 1,
	}
}
