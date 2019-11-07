# 安装go语言环境

## 安装

### 下载解压安装

国内下载地址：[https://studygolang.com/dl](https://studygolang.com/dl)

```
wget https://studygolang.com/dl/golang/go1.13.1.linux-amd64.tar.gz
mkdir /usr/local/go
tar -xvf go1.13.1.linux-amd64.tar.gz -C /usr/local/go
mv /usr/local/go/go /usr/local/go/go1.13
```

### 源码编译安装

> 如果git clone很慢的话可以访问地址`https://www.ipaddress.com/ip-lookup`来查询`github.com`和`github.global.ssl.fastly.net`的ip，然后配置`/etc/hosts`文件，添加如`192.30.253.112 github.com`和`199.232.5.194 github.global.ssl.fastly.net`

Go高版本的编译过程需要Go1.4的二进制来实现引导,因为go1.4是c编写的go工具链的最后一个发布版，所以它的编译不需要go的编译器，用gcc和glibc-devel，而1.5往后的版本都是通过go自身来编译的。

1. 如果装过go，先清除go的相关环境变量
   ```
   #先注释掉配置文件里相关的值,然后再手动unset环境变量
   env | grep -i go
   unset GOROOT
   ...
   ```
2. 安装git，并配置使能从github上拉代码（一般都有）
3. 安装`gcc`和`glibc-devel`【arch好像没有glibc-devel，但我还是编译成功了】
4. 下载go1.4.3源码并编译
   ```
   mkdir /usr/local/go && cd /usr/local/go
   wget https://dl.google.com/go/go1.4.3.src.tar.gz
   tar -xvf go1.4.3.src.tar.gz
   mv go go1.4.3
   cd go1.4.3/src
   # 这里必须禁用cgo，否则会报错（我也不知道为啥，我的操作系统是manjaro）
   CGO_ENABLED=0 ./make.bash
   # 成功的话打印如下日志
   ...
   ---
   Installed Go for linux/amd64 in /usr/local/go/go
   Installed commands in /usr/local/go/go/bin
   ```
5. 设置go引导环境变量
   ```
   vim ~/.bashrc
   export GOROOT_BOOTSTRAP=/usr/local/go/go1.4.3
   source ~/.bashrc
   ```
6. 安装go1.5以上任意版本（这里是1.13）
   ```
   cd /usr/local/go/
   git clone git@github.com:golang/go.git
   cd go
   git checkout -b release-branch.go1.13 origin   release-branch.go1.13
   cd go/src
   ./make.bash
   # 日志如下
   Building Go cmd/dist using /usr/local/go/go1.4.3.
   Building Go toolchain1 using /usr/local/go/go1.4.3.
   Building Go bootstrap cmd/go (go_bootstrap) using Go   toolchain1.
   Building Go toolchain2 using go_bootstrap and Go toolchain1.
   Building Go toolchain3 using go_bootstrap and Go toolchain2.
   Building packages and commands for linux/amd64.
   ---
   Installed Go for linux/amd64 in /usr/local/go/go
   Installed commands in /usr/local/go/go/bin
   ```

## 配置

```
vim ~/.bashrc
```

```
# go
#  GO111MODULE：是 Go modules 的开关
#  GOPROXY：默认是proxy.golang.org，国内访问不了。
#  GOPATH：有了gomod之后，它只是用来存储依赖的包
export GOROOT_BOOTSTRAP=/usr/local/go/go1.4.3
export GOROOT=/usr/local/go/go1.13
export GOPATH=/home/yjx/Repository/code/golang/package
export PATH=$GOROOT/bin:$PATH
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
```

```
source ~/.bash_profile
```

## 依赖管理

### go版本管理

可以使用`gvm`插件对go版本进行管理，但我还是习惯手动编译新的版本，然后再改环境变量，切换go版本。

### 包依赖管理

以前用的都是godep等工具进行管理，自go1.11 `go mod`出现，其他都可以放弃了，用起来不要太爽。

详细了解可参考文章：[干货满满的 Go Modules 和 goproxy.cn](https://github.com/EDDYCJY/blog/blob/master/talk/goproxy-cn.md)
