package test02

import "fmt"

// TestVar 测试变量
func TestVar() {
	fmt.Printf("###### 变量的声明\n")

	var a1 bool
	var a2 int64
	var a3 uint8
	var a4 complex64
	var (
		b1 string
		b2 []float32
		b3 func() bool
		b4 struct {
			x int
		}
	)

	fmt.Printf("bool:  %v\n", a1)
	fmt.Printf("int64：  %v\n", a2)
	fmt.Printf("uint8：  %v\n", a3)
	fmt.Printf("complex64：  %v\n", a4)
	fmt.Printf("string：  %v\n", b1)
	fmt.Printf("float32：  %v\n", b2)
	fmt.Printf("func() bool：  %v\n", b3)
	fmt.Printf("struct {x int}：  %v\n", b4)

	fmt.Printf("###### 短变量的声明和初始化\n")
	c1 := 100
	c2, c3 := 1, "abc"
	fmt.Printf("简短格式，有值：  c1=%v , c2=%v , c3=%v\n", c1, c2, c3)


	fmt.Printf("###### 空白标识符\n")
	_,num1,str1 := numbers()
	fmt.Println(num1, str1)
}

// 空白标识符
func numbers()(int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}