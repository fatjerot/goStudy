# 基本结构和基本数据类型

## 项目结构

有了`go mod`后，项目结构不需要像以前那么复杂了。

**$GOPATH**：

```shell
[yjx@manjaro-yjx go]$ tree -L 4 $GOPATH
/home/yjx/Repository/code/golang/package
├── bin
│   ├── dlv
│   ├── fillstruct
│   ├── gocode
│   ├── gocode-gomod
│   ...
├── pkg
│   ├── mod
│   │   ├── 9fans.net
│   │   │   └── go@v0.0.0-20181112161441-237454027057
│   │   ├── cache
│   │   │   ├── download
│   │   │   ├── lock
│   │   │   └── vcs
│   │   ├── github.com
│   │   │   ├── acroca
│   │   │   ├── cosiner
│   │   │   ├── cweill
│   │   │   ...
│   │   └── golang.org
│   └── sumdb
│       └── sum.golang.org
│           └── latest
└── src
    ├── github.com
    └── golang.org
        └── x
            └── tools
```

- src：放源码
- pkg：自从用了`go mod`后就没找到`.a`文件，以后再深入吧
- bin：`go install`后生成的可执行文件

**其他目录只要有`go.mod`文件，都可以算一个工作目录**

以后想起了再补充。。。

## 语言结构

### 基础语法

```
package main

import "fmt"

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
}
```

- 包声明
- 引入包
- 函数&方法
- 注释
- 语句 & 表达式
- 行分隔符
- 标识符
- 关键字

### init和main

golang里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。

- 相同点：两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
- 不同点：init可以应用于任意包中，且可以重复定义多个。main函数只能用于main包中，且只能定义一个。

执行顺序：

- 对同一个go文件的init()调用顺序是从上到下的
- 对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数
- 对不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()
- 如果package存在依赖，则先调用最早被依赖的package中的init()
- 最后调用main函数

### 初始化顺序

![](https://upload-images.jianshu.io/upload_images/1537644-f79036ee350856bf.png)

- 当 A 没有引入其他包的情况下执行顺序是：`package > const > var > init() > main() > 类型的方法 > 其余函数`
- 当 A 里引入了包 B 的时候会有一些变化，程序会先执行包 B 里面的`import，const，var，init（）`，执行完毕之后跳回 A ，继续按照顺序执行。
- 如果包B里面引入了C包，如上规则执行
- 普通方法函数按调用排序，然后再用字母排序。

**init和main**：

相同点：

- 两个函数在定义时不能有任何的参数和返回值
- 该函数只能由 go 程序自动调用，不可以被引用

不同点：

- init 可以应用于任意包中，且可以重复定义多个。
- main 函数只能用于 main 包中，且只能定义一个。

### 包的概念

#### 推荐结构

- 项目根目录包名main，文件名main.go
- 子目录下文件里的包名和目录名一致，单个文件则和目录名保持一致，多个文件随意

#### 引用格式：

```
# 单个包
import "fmt"
# 多个包
import (
   "fmt"
   "os"
)
```

#### 可见性规则：

当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，那这个标识符就被定义为public（如：`fmt.Println("Hello, World!")`）；如果是以小写字母开头，在包内可见，但在包外(即其他包导入该包后)是不可使用的。

#### import

公有包导入就不提了，主要说一下私有包的导入：

1. 被引入.go文件和引入.go文件处于同一项目，同一目录下
   
   - 这种情况其实两个文件都是在同一个包下的，可以直接引用小写字母开头的方法
2. 被引入.go文件和引入.go文件处于同一项目不同目录下
   - 格式：`项目的module名/从项目mod文件到该文件所在的目录的路径【不是包名】`
   ```
   # 项目结构
    ├── 001basic_datatype
    │   ├── data_type.go
    │   ├── hello
    │   │   └── hello.go
    │   └── README.md
    ├── go.mod
    ├── main.go
    └── README.md
   ```
    ```
    # go.mod
    module go_study_basic

    go 1.13
    ```
   ```
   # main.go引用data_type.go
   import datatype "go_study_basic/001basic_datatype"
   # data_type.go引用hello.go
   import "go_study_basic/001basic_datatype/hello"
   ```
3. 如果是不同的私有项目，以后再谈

#### 注意点：

- 文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
- 文件夹名与包名没有直接关系，并非需要一致。
- 同一个文件夹下的文件只能有一个包名，否则编译报错。
- 所有的包名都应该使用小写字母。
- 如果对一个包进行更改或重新编译，所有引用了这个包的客户端程序都必须全部重新编译。
- 如果你导入了一个包却没有使用它，则会在构建程序时引发错误

## 命名规范

### 文件名

文件命名一律采用小写，不用驼峰式，尽量见名思义，看见文件名就可以知道这个文件下的大概内容。

除了测试文件可以用下划线连接，其他最好都是小写。

### 包名

一律小写，使用短命名，尽量不要和标准库冲突。

### 变量

采用驼峰式，遇到特有名词可以全大写或全小写。

### 常量

和变量相同，如果模块复杂，可按功能统一定义在包下的一个文件里。

### 接口

单个函数 接口名以`er`为后缀

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

2个函数 接口名综合两个函数

```go
type WriteFlusher interface {
    Write([]byte) (int, error)
    Flush() error
}
```

三个以上函数的接口名类似于结构体名

```go
type Car interface {
    Start() 
    Stop()
    Drive()
}
```

### 结构体

结构体名应该是名词或名词短语，如`Account`,`Book`

如果该数据结构需要序列化，如json， 则首字母大写， 包括里面的字段。

### 方法

方法名应该是动词或动词短语，采用驼峰式。将功能及必要的参数体现在名字中， 不要嫌长， 如updateById，getUserInfo.

如果是结构体方法，那么 `Receiver` 的名称应该缩写，一般使用一个或者两个字符作为 `Receiver` 的名称。

```go
func (f foo) method() {
    ...
}
```

如果 `Receiver` 是指针， 那么统一使用p。

```go
func (p *foo) method() {
    ...
}
```

### 注释

每个包都应该有一个包注释，位于 package 之前

```go
// balabala.
package net
```

每个以大写字母开头（即可以导出）的方法应该有注释，且以该函数名开头。如：

```go
// Get 会响应对应路由转发过来的 get 请求
func (c *Controller) Get() {
    ...
}
```

### README

每个文件夹下都应该有一个README文件，该文件是对当前目录下所有文件的一个概述，和主要方法描述。并给出一些相应的链接地址，包含代码所在地、引用文档所在地、API文档所在地,实际情况可酌情考虑
