/**
  ******************************************************************************
  * File Name          : 上传 Controller
  * Author             : 张宇恺
  * Description        : 上传 Controller 主要有以下几个功能:
												 1. 插入新数据到数据库
												 2. 如果 stuno 重复，会提示已经有该学生
  ******************************************************************************
*/

package controller

import (
	"fmt"
	"iris/src/database"
	"iris/src/model"
	logMsg "iris/src/utils"

	"github.com/kataras/iris/v12"
)

type UploadController struct{}

func (uc *UploadController) PostUpload(ctx iris.Context) model.ResponseModel {
	path := ctx.Path()
	// 数据初始化工作
	var stuno, sex, birthday, courseid, teacherid, stuname string
	var stunoArray []string

	keyArray := []string{"stuname", "sex", "birthday", "courseid", "teacherid"}
	valueArray := []interface{}{stuname, sex, birthday, courseid, teacherid}

	stuno = ctx.PostValue("stuno")
	for i := 0; i < len(keyArray); i++ {
		key := keyArray[i]
		valueArray[i] = ctx.PostValue(key)
	}
	valueArray = append([]interface{}{stuno}, valueArray...)

	stunoSql := "select stuno from student;"
	sql := "insert into student (stuno,stuname,sex,birthday,courseid,teacherid) values (?,?,?,?,?,?);"

	// 查当前 mysql 中 stuno 的全部信息
	stunoSqlErr := database.DB.Select(&stunoArray, stunoSql)
	if stunoSqlErr != nil {
		fmt.Println("database.DB.Select error, no stunoArray❌", stunoSqlErr.Error())
		return model.ResponseModel{
			Data:       nil,
			Msg:        "database.DB.Select error, no stunoArray",
			Code:       0,
			TotalCount: 0,
		}
	}

	// 如果 mysql 中 stuno 已经包括这名学生，不插入此条信息到 mysql
	for _, value := range stunoArray {
		if value == stuno {
			return model.ResponseModel{
				Data:       nil,
				Msg:        "included current studentNumber, deny to insert",
				Code:       2,
				TotalCount: 0,
			}
		}
	}

	// 执行 mysql 插入语句
	res, err := database.DB.Exec(sql, valueArray...)

	if err != nil {
		fmt.Println("database.DB.Exec error, insert error❌", err.Error())
		return model.ResponseModel{
			Data:       nil,
			Msg:        "database.DB.Exec error, sql 插入语句写的有问题",
			Code:       0,
			TotalCount: 0,
		}
	}
	defer logMsg.LogSuccessMsg(path, "Post")

	return model.ResponseModel{
		Data:       res,
		Code:       1,
		Msg:        "success",
		TotalCount: 1,
	}
}

func (uc *UploadController) OptionsUpload(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization")
	ctx.Header("Access-Control-Max-Age", "3600")
}
