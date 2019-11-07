package test02

import "fmt"

func TestConst() {
	fmt.Println("###### 字符串常量")
	const hello = "hello world!"
	fmt.Printf("type %T value %v\n", hello, hello)

	fmt.Println("###### 数值常量")
	const a = 5
	var intVar int = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "float64Var", float64Var, "complex64Var", complex64Var)
}
