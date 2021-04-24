# 使用 go, iris, mysql 搭建的乞丐版后台系统

## 接口说明
* /student/message 返回所有的学生信息
* /student/course?stuno=xxx 返回某个学生选的课程
* /student/teacher?stuno=xxx 返回某个学生选的老师
* /student/upload 上传某个学生信息  
* /find/teacher 返回全部老师
* /find/course 返回全部课程
* /record 返回当前访问人数

## 编译
> CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
