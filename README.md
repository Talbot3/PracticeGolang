# Env
go env

## Golang install by [`g`](https://github.com/voidint/g)
g install 1.14.2

## Go seed
function
* 1.default directory
* 2.default makefile

## run

```
make build
./build/main add 5 5
```

## 根据个人开发环境`go env`更改配置 `GOROOT`

```bash
#macosx
GOROOT := /Users/:username/.g/go
```

## Go Modules

- fswatch
- hello

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