# 变量

## 变量的声明

go语言是静态类型语言，变量是一段或多段用来存储数据的内存。

所有的内存在go中都是初始化过的，系统自动赋予它该类型的0值，如：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等

声明变量，类型放在变量后面：

```go
// 标准格式
var name type

// 批量格式
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)

// 简短模式【不用var声明，是有条件的：】
// 3.声明后的变量不能用:=
// 2.声明的时候不能指定type类型
// 1.只能在函数内部
func main() {
   x:=100
   a,s:=1, "abc"
}

//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3
var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

type

- bool
- string
- int、int8、int16、int32、int64
- uint、uint8、uint16、uint32、uint64、uintptr
- byte // uint8 的别名
- rune // int32 的别名 代表一个 Unicode 码
- float32、float64
- complex64、complex128

## 变量的初始化

- 整型和浮点型变量的默认值为 0 和 0.0。
- 字符串变量的默认值为空字符串。
- 布尔型变量默认为 bool。
- 切片、函数、指针变量的默认为 nil。

#### 短变量声明并初始化

短变量声明并初始化的格式在开发中使用比较普遍

```go
conn, err := net.Dial("tcp","127.0.0.1:8080")
```

在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错

```go
conn, err := net.Dial("tcp", "127.0.0.1:8080")
conn2, err := net.Dial("tcp", "127.0.0.1:8080")
```

#### 编译器推导类型的格式

```go
var attack = 40
var defence = 20
var damageRate float32 = 0.17
var damage = float32(attack-defence) * damageRate
fmt.Println(damage)
```

- 第 1 和 2 行，右值为整型，attack 和 defence 变量的类型为 int。
- 第 3 行，表达式的右值中使用了 0.17。由于Go语言和C语言一样，编译器会尽量提高精确度，以避免计算中的精度损失。所以这里如果不指定 damageRate 变量的类型，Go语言编译器会将 damageRate 类型推导为 float64，我们这里不需要 float64 的精度，所以需要强制指定类型为 float32。
- 第 4 行，将 attack 和 defence 相减后的数值结果依然为整型，使用 float32() 将结果转换为 float32 类型，再与 float32 类型的 damageRate 相乘后，damage 类型也是 float32 类型。

#### 交换变量

```go
//两个变量的类型必须相同才能交换
a, b = b, a
```

#### 空白标识符

在函数返回值时的使用：

```go
package main

import "fmt"

func main() {
  _,numb,strs := numbers() //只获取函数返回值的后两个
  fmt.Println(numb,strs)
}

//一个可以返回多个值的函数
func numbers()(int,int,string){
  a , b , c := 1 , 2 , "str"
  return a,b,c
}
```

空白标识符 `_` 也被用于抛弃值，如值 5 在`_, b = 5, 7`中被抛弃。

#### 全局变量和局部变量，静态动态

如果全局变量的首字母大写，那么它就是公开的全局变量。如果全局变量的首字母小写，那么它就是内部的全局变量。

go没有c语言的静态变量


## 常量

#### 定义

关键字 const 修饰的名字为常量，不能被重新赋予任何值。

```go
package main

func main() {  
    const a = 55
    a = 89 //报错
}
```

常量的值在编译期间确定。因此不能将函数的返回值赋给常量，因为函数调用发生在运行期。

```go
package main

import (  
    "fmt"
    "math"
)

func main() {  
    fmt.Println("Hello, playground")
    var a = math.Sqrt(4)//allowed
    const b = math.Sqrt(4)//not allowed
}
```

#### 字符串常量

Go是强类型语言。在赋值时混合使用类型是不允许的。让我们通过以下代码说明这是什么意思。

```go
package main

func main() {  
    var defaultName = "Sam" //allowed
    type myString string
    var customName myString = "Sam" //allowed
    customName = defaultName //not allowed
}
```

#### 布尔常量

布尔常量与字符串常量（在概念上）没有区别。

#### 数值常量

包括整数，浮点数以及复数常量。

```go
package main

import (  
    "fmt"
)

func main() {  
    fmt.Println("Hello, playground")
    const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
}
```

a 的值是 5 并且 a 在语法上是泛化的（它既可以表示浮点数 `5.0`，也可以表示整数 `5`，甚至可以表示没有虚部的复数 `5 + 0i`）

#### 枚举类型

go语言并没有提供enum的定义，我们可以使用const来模拟枚举类型。

```go
package main

import  (
  "fmt"
)

type PolicyType int32

const (
    Policy_MIN      PolicyType = 0 
    Policy_MAX      PolicyType = 1 
    Policy_MID      PolicyType = 2 
    Policy_AVG      PolicyType = 3 
)

func (p PolicyType) String() string {
    switch (p) {
    case Policy_MIN: return "MIN"
    case Policy_MAX: return "MAX"
    case Policy_MID: return "MID"
    case Policy_AVG: return "AVG"
    default:         return "UNKNOWN"
    }
}

func foo(p PolicyType) {
    fmt.Printf("enum value: %v\n", p)
}

func main() {
    foo(Policy_MAX)
}
```