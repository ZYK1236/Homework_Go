/**
  ******************************************************************************
  * File Name          : 查询所有课程 Controller
  * Author             : 张宇恺
  * Description        : 查表，给出表中所有课程 id 以及名字
  ******************************************************************************
*/

package controller

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"

	"github.com/kataras/iris/v12"
)

type FindCourseController struct{}

type FindCourse struct {
	Id         int    `json:"id"`
	CourseName string `json:"coursename"`
}

func (fc *FindCourseController) GetCourse(ctx iris.Context) model.ResponseModel {
	result := []FindCourse{}
	sql := "select id, coursename from cource;"
	err := database.DB.Select(&result, sql)
	if err != nil {
		fmt.Print("database.DB.Select error 查找所有课程 id 失败 ❌", err.Error())
		panic(err.Error())

	}

	defer fmt.Println("path:/find/course ----> GET ✅")

	return model.ResponseModel{
		Data:       result,
		Code:       1,
		Msg:        "success",
		TotalCount: len(result),
	}
}
