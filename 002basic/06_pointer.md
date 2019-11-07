# 指针

指针是存储一个变量的内存地址的变量。

## 声明

```go
package main

import (  
    "fmt"
)

func main() {  
    b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("address of b is", a)
}
```

## 空值

