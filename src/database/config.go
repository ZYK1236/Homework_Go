package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Init() {
	var db *sqlx.DB

	//定义mysql数据源，配置数据库地址，帐号以及密码
	dsn := "homeworkgo:DF354NH2!69iuis@tcp(homeworkgo.mysql.polardb.rds.aliyuncs.com)/homeworkgo?charset=utf8&parseTime=True&loc=Local"

	//根据数据源dsn和mysql驱动, 创建数据库对象
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库启动成功", db)
}