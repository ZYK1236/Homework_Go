/**
 ******************************************************************************
 * File Name          : 查询学生对应信息
 * Author             : 张宇恺
 * Description        : 查表，给出所有学生的身份信息
 ******************************************************************************
 */

package controller

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"
	logMsg "iris/src/utils"

	"strconv"

	"github.com/kataras/iris/v12"
)

// Student 标签千万不要空格！！ 5555
type Student struct {
	Id        int    `json:"id"`
	Stuno     int    `json:"stuno"`
	Stuname   string `json:"stuname"`
	Sex       int    `json:"sex"`
	Birthday  int    `json:"birthday"`
	CourseId  int    `json:"courseid"`
	TeacherId int    `json:"teacherid"`
}

type StudentController struct{}

func (sc *StudentController) GetMessage(ctx iris.Context) model.ResponseModel {
	var beginId int
	var student []Student

	// get query param
	path := ctx.Path()
	pageNum := ctx.URLParam("pageNum")

	if pageNum == "" {
		beginId = 1
	} else {
		beginId, _ = strconv.Atoi(pageNum)
		beginId = (beginId - 1) * 5
	}

	// defined sql
	count := 0
	selectSql := "select * from student where id>? and id<=?"
	countSql := "select count(*) from student"

	// exec sql
	err := database.DB.Select(&student, selectSql, beginId, beginId+5)
	errorCount := database.DB.Get(&count, countSql)

	if err != nil || errorCount != nil {
		fmt.Println("database.DB.Select / Get error ❌")
		return model.ResponseModel{
			Data: nil,
			Code: 0,
			Msg:  "error",
		}
	}

	defer logMsg.LogSuccessMsg(path, "Get")

	result := model.ResponseModel{
		Data:       student,
		Code:       1,
		Msg:        "success",
		TotalCount: count,
	}

	return result
}

func (sc *StudentController) OptionsMessage(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization")
	ctx.Header("Access-Control-Max-Age", "3600")
}
