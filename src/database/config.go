package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init() {
	var err error
	//定义mysql数据源，配置数据库地址，帐号以及密码
	dsn := "homeworkgo:DF354NH2!69iuis@tcp(homeworkgo.mysql.polardb.rds.aliyuncs.com)/homeworkgo?charset=utf8&parseTime=True&loc=Local"

	//根据数据源dsn和mysql驱动, 创建数据库对象
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。
	// 如果要检查数据源的名称是否真实有效，应该调用Ping方法
	err = DB.Ping()
	if err != nil {
		panic("sql can not connect.." + err.Error())
	}

	fmt.Println("数据库启动成功 ✅ --->", DB)
}
