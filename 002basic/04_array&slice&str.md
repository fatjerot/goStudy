## 数组

Go 语言里面的数组其实很不常用，这是因为数组是定长的静态的，一旦定义好长度就无法更改，而且不同长度的数组属于不同的类型，之间不能相互转换相互赋值，用起来多有不方便之处。

#### 声明

```go
package main

import (  
    "fmt"
)

func main() {  
    var a [3]int //int array with length 3
    fmt.Println(a)

    b := [3]int{12, 78, 50} // shorthand declaration to create array
    fmt.Println(b)

    c := [2]int{12} 
    fmt.Println(c) // [12 0]

    b = c // wrong
}
```

#### 访问赋值

在 Go 中数组是值类型而不是引用类型。这意味着当数组变量被赋值时，将会获得原数组的拷贝。新数组中元素的改变不会影响原数组中元素的值。

```go
package main

import "fmt"

func main() {
    var a = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    var b [9]int
    b = a
    a[0] = 12345
    fmt.Println(a)
    fmt.Println(b)
}

--------------------------
[12345 2 3 4 5 6 7 8 9]
[1 2 3 4 5 6 7 8 9]
```

如果将数组作为参数传递给函数，仍然是值传递，在函数中对（作为参数传入的）数组的修改不会造成原数组的改变。

```go
package main

import "fmt"

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func main() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}

--------------------------
before passing to function  [5 6 7 8 8]  
inside function  [55 6 7 8 8]  
after passing to function  [5 6 7 8 8]  
```

#### 遍历

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
```

```go
package main

import "fmt"

func main() {
    var a = [5]int{1,2,3,4,5}
    for index, _ := range a {
        fmt.Println(index, a[index])
    }
    for index, value := range a {
        fmt.Println(index, value)
    }
}

------------
0 1
1 2
2 3
3 4
4 5
0 1
1 2
2 3
3 4
4 5
```

#### 多维数组

```go
package main

import (  
    "fmt"
)

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {  
    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
}
```

a数组末尾的逗号是必须的，这是因为词法分析器会根据一些简单的规则自动插入分号。如果你想了解更多，请阅读：https://golang.org/doc/effective_go.html#semicolons。

#### 下标越界

Go 会在编译后的代码中插入下标越界检查的逻辑，所以数组的下标访问效率是要打折扣的，比不得 C 语言的数组访问性能。

```go
package main

import "fmt"

func main() {
    var a = [5]int{1,2,3,4,5}
    var b = 101
    a[b] = 255
    fmt.Println(a)
}

------------
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
    /Users/qianwp/go/src/github.com/pyloque/practice/main.go:8 +0x3d
exit status 2
```

## slice

#### 原理

切片是动态的数组，是可以扩充内容增加长度的数组。当长度不变时，它用起来就和普通数组一样。当长度不同时，它们也属于相同的类型，之间可以相互赋值。这就决定了数组的应用领域都广泛地被切片取代了。

内部结构非常类似于 ArrayList。当数组容量不够需要扩容时，就会换新的数组，还需要将老数组的内容拷贝到新数组。ArrayList 内部有两个非常重要的属性 capacity 和 length。capacity 表示内部数组的总长度，length 表示当前已经使用的数组的长度。length 永远不能超过 capacity。

![](https://pic3.zhimg.com/80/v2-a6a2110a710cb7b842ba7954b6efdb7a_hd.jpg)

上图中一个切片变量包含三个域，分别是底层数组的指针、切片的长度 length 和切片的容量 capacity。切片支持 append 操作可以将新的内容追加到底层数组，也就是填充上面的灰色格子。如果格子满了，切片就需要扩容，底层的数组就会更换。

#### 创建

##### 普通方式[满容]

- 数组：`[3]int{1,2,3}`
- 切片：`[]int{1,2,3}`

```go
package main

import "fmt"

func main() {
 var s []int = []int{1,2,3,4,5}  // 满容的
 fmt.Println(s, len(s), cap(s))
}

---------
[1 2 3 4 5] 5 5
```

##### make[可指定]

使用 make 函数创建的切片内容默认是「零值切片」，也就是内部数组的元素都是零值。

`make(类型,长度,容量)`

```go
package main

import "fmt"

func main() {
 var s1 []int = make([]int, 5, 8)  
 var s2 []int = make([]int, 8) // 满容切片
 fmt.Println(s1)
 fmt.Println(s2)
}

-------------
[0 0 0 0 0]
[0 0 0 0 0 0 0 0]
```

##### nil&空 切片

```go
package main

import "fmt"

func main() {
 var s1 []int                  // 「nil切片」
 var s2 []int = []int{}        // 「空切片」
 var s3 []int = make([]int, 0) // 「空切片」
 fmt.Println(s1, s2, s3)
 fmt.Println(len(s1), len(s2), len(s3))
 fmt.Println(cap(s1), cap(s2), cap(s3))
}

-----------
[] [] []
0 0 0
0 0 0
```

##### 通过切片创建新的切片

`slice[i:j:k]` --> `{slice[i],..,slice[j+1]}  cap:k-i`

-  i 表示从老切片的第i下标开始切【第i+1个】
-  j 表示切到老切片的第j个【第j个】
-  k 表示新切片的容量为(k-i)，如果没有给定k，则表示切到底层数组的最尾部，k最大不能超过老切片容量【k-(i-0)】

```go
myNum := []int{1, 2, 3, 4, 5} // [1,2,3,4,5] cap:5
newNum1 := slice[:]           // [1,2,3,4,5] cap:5
newNum2 := slice[1:]          // [2,3,4,5]   cap:4
newNum3 := slice[:3]          // [1,2,3]     cap:5
newNum3 := slice[2:3]         // [3,4]       cap:3
newNum3 := slice[2:3:5]       // [3]         cap:3
```

#### 扩容

当 cap 小于 1024 的时候，是成倍的增长，超过的时候，每次增长 25%，而这种内存增长不仅仅数据拷贝（从旧的地址拷贝到新的地址）需要消耗额外的性能，旧地址内存的释放对 gc 也会造成额外的负担，所以如果能够知道数据的长度的情况下，尽量使用 make([]int, len, cap) 预分配内存

```go
myNum := []int{10, 20, 30, 40}
newNum := append(myNum, 50)
```

#### 内存重用

```go
si1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
si2 := si1[:7]
Convey("不重新分配内存", func() {
    header1 := (*reflect.SliceHeader)(unsafe.Pointer(&si1))
    header2 := (*reflect.SliceHeader)(unsafe.Pointer(&si2))
    fmt.Println(header1.Data)
    fmt.Println(header2.Data)
    So(header1.Data, ShouldEqual, header2.Data)
})

Convey("往切片里面 append 一个值", func() {
    si2 = append(si2, 10)
    Convey("改变了原 slice 的值", func() {
        header1 := (*reflect.SliceHeader)(unsafe.Pointer(&si1))
        header2 := (*reflect.SliceHeader)(unsafe.Pointer(&si2))
        fmt.Println(header1.Data)
        fmt.Println(header2.Data)
        So(header1.Data, ShouldEqual, header2.Data)
        So(si1[7], ShouldEqual, 10)
    })
})
```

si2 是 si1 的一个切片，从第一段代码可以看到切片并不重新分配内存，si2 和 si1 的 Data 指针指向同一片地址，而第二段代码可以看出，当我们往 si2 里面 append 一个新的值的时候，我们发现仍然没有内存分配，而且这个操作使得 si1 的值也发生了改变，因为两者本就是指向同一片 Data 区域，利用这个特性，我们只需要让 si1 = si1[:0] 就可以不断地清空 si1 的内容，实现内存的复用

你可以使用 copy(si2, si1) 实现深拷贝

## 字符串
