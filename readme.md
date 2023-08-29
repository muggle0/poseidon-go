# golang 开发实战
web 开发是学习曲线不那么陡峭的入门方式，通过学习web开发来深入的学习golang

## 项目框架介绍

**gin**

gin 框架是后端web开发比较流行的框了;当然Beego的开发难度要远低于gin,可能更适合初学者，但是从学习的角度来讲,Beego并不适合拿来练手。

gin文档 https://www.kancloud.cn/shuangdeyu/gin_book/949453

gorm orm 框架 http://gorm.book.jasperxu.com/

注意了小伙伴们，这个系列教程不适合没有语言基础的同学

## 项目结构

## gin 使用教程

安装go 
gopath goroot
go env

安装gin 

go get -u github.com/gin-gonic/gin

go env -w GO111MODULE=on
go mod init
go mod download

GOPROXY="https://goproxy.io,direct"

## go 关键字
package

Go语言中存在两种包、一种是可执行程序的包、一种是类库函数的包。

可执行程序的包，编译完成后会生成一个可执行文件、静态库的包编译之后会生成一个.a为后缀的文件，自己不能执行只能被可执行包引用。

可执行程序的包必须以main作为包名，静态库的包名没有要求。

func()

go run main.go

go get -u github.com/pilu/fresh

go install github.com/pilu/fresh@latest

## go-admin-core



## 注意事项

内嵌结构体

*p p &p
