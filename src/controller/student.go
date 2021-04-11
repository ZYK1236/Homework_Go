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

func (sc *StudentController) GetMessage(ctx iris.Context) model.ResponseModel {
	var beginId int

	// get query param
	pageNum := ctx.URLParam("pageNum")
	if pageNum == "" {
		beginId = 1
	} else {
		beginId, _ = strconv.Atoi(pageNum)
		beginId = (beginId - 1) * 5
	}

	// defined sql
	student := []Student{}
	count := 0
	selectSql := "select * from student where id>? and id<=?"
	countSql := "select count(*) from student"

	// exec sql
	error := database.DB.Select(&student, selectSql, beginId, beginId+5)
	error_count := database.DB.Get(&count, countSql)

	if error != nil || error_count != nil {
		fmt.Println("database.DB.Select / Get error ❌")
		panic(error.Error())
	}

	defer fmt.Println("path:/student/message ----> GET ✅")

	result := model.ResponseModel{
		Data:       student,
		Code:       1,
		Msg:        "success",
		TotalCount: count,
	}

	return result
}
