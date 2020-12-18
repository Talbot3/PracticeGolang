# golang 练习项目

## Env
go env

## Golang install by [`g`](https://github.com/voidint/g)
g install 1.14.2

## Go seed
function
* 1.default directory
* 2.default makefile

## Run

```bash
cd packages/xx
go build
# go run .
```

## Go Modules

- 开启GO111MODULE
用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。

> GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。

> 在使用模块的时候，GOPATH 是无意义的，不过它还是会把下载的依赖储存在 $GOPATH/src/mod 中，也会把 go install 的结果放在 $GOPATH/bin 中。


```bash
go mod init创建一个新模块，初始化go.mod描述它的文件。
go build，go test以及其他软件包构建命令go.mod根据需要添加新的依赖项。
go list -m all 打印当前模块的依赖项。
go get 更改所需的依赖版本（或添加新的依赖）。
go mod tidy 删除未使用的依赖项。
```

## 引用 

- [Go语言圣经](https://books.studygolang.com/gopl-zh/)
- [using-go-modules](https://blog.golang.org/using-go-modules)
- [golang-package](https://studygolang.com/articles/5831)
- [Golang Package 与 Module 简介](https://www.jianshu.com/p/07ffc5827b26)
- [Go 指南](https://tour.go-zh.org/basics/1)