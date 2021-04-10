package controller

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"

	"strconv"
	"github.com/kataras/iris/v12"
)

// 标签千万不要空格！！ 5555
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

func (mc *StudentController) GetMessage(ctx iris.Context) *model.ResponseModel {
	var beginId int

	// 获取 query 参数
	param := ctx.URLParam("pageNum")
	if param == "" {
		beginId = 1
	} else {
		beginId, _ = strconv.Atoi(param)
	}

	// exec sql
	student := []Student{}
	sql := "select * from student where id>=? and id<=?"
	error := database.DB.Select(&student, sql, beginId, beginId+2)

	if error != nil {
		fmt.Println("database.DB.Select error ❌")
		panic(error.Error())
	}
	defer fmt.Println("path:/student/message ----> GET ✅")

	result := model.ResponseModel{
		Data: student,
		Code: 1,
		Msg:  "success",
	}

	return &result
}
