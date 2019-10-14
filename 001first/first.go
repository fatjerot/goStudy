package first

import (
	"fmt"
	"go_study_basic/001first/testinit"
)

// 第一个go程序
func Hello() {
	testinit.Testinit()
	fmt.Printf("hello golang !\n")
}

// 测试init函数执行顺序
func init() {
	fmt.Printf("init inner\n")
}
