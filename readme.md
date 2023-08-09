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

## 年终总奖

2022年计划是写20篇的博客和四篇的视频教学，没有达成目标，主要原因是自己偷懒了。
这一年的技术重心是云原生的方向，
在学习过程中发现搞云原生成本有点高了，
我准备了一台4核八G的服务器，做到后面发现资源还是不够用，
而且云原生涉及的技术栈太深搞起来不那么容易出成果。
2023年的工作重心打算放在go语言的学习上，然后会花一部分精力继续搞云原生。
2023年的目标依旧是20篇的博客和四篇的视频教学。

最后祝各位新的一年大展宏兔，钱兔无量。
