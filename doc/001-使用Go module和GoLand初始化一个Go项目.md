## 简介
Golang是一门优秀的语言，特别是在并发编程上，得益于它的协程和channel等，非常方便易用。它通过go module来管理包和依赖，本文介绍如何利用它创建一个项目。


##  重要的环境变量

Go通过环境变量来做项目上的管理和控制，通过命令go env可以查看相关变量：


```
$ go env
GO111MODULE="on"
GOARCH="amd64"
GOHOSTOS="darwin"
GOMODCACHE="/Users/larry/go/pkg/mod"
GOPATH="/Users/larry/go"
GOPROXY="https://mirrors.aliyun.com/goproxy/"
GOROOT="/Users/larry/Software/go"
GOTOOLDIR="/Users/larry/Software/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
PKG_CONFIG="pkg-config"
```
关键环境变量：

GOROOT：Go的安装目录，即可执行文件所在的目录；

GOPATH：工作目录并不是项目所有目录，编译后的二进制文件存放地，import包的搜索路径，主要包含bin、pkg、src；

GO111MODULE：启用go module管理项目，需要有go.mod和go.sum文件；

GOPROXY：下载依赖时的代理，必须配置，不然无法成功下载；

常用的代理有：


```
# 1. 七牛 CDN
export GOPROXY=https://goproxy.cn,direct

# 2. 阿里云
export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 3. 官方
export GOPROXY=https://goproxy.io,direct
```
## 初始化项目
用命令执行如下：


```
# 创建project的目录
$ mkdir pkslow_test
# 进入目录
$ cd pkslow_test/
# 初始化
$ go mod init pkslow.com/pkslow_test
go: creating new go.mod: module pkslow.com/pkslow_test

$ l
-rw-r--r--   1 larry  staff   39 Dec 13 21:07 go.mod
#查看文件内容
$ cat go.mod 
module pkslow.com/pkslow_test

go 1.15
```

执行或编译后执行：


```
# 直接run
$ go run main.go 
hello pkslow

# 编译成二进制文件
$ go build
t$ l
-rw-r--r--   1 larry  staff       39 Dec 13 21:07 go.mod
-rw-r--r--   1 larry  staff       74 Dec 13 21:12 main.go
-rwxr-xr-x   1 larry  staff  2146904 Dec 13 21:12 pkslow_test
# 执行二进制文件
$ ./pkslow_test 
hello pkslow
```

## 引入本地包

![image](doc/a1b2.jpg)










