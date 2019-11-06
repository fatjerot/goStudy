// go基础学习测试
package main

import (
	"fmt"
	"go_study_basic/002basic/test01"
	"go_study_basic/002basic/test02"
)

func init() {
	fmt.Printf("init\n")
}

func main() {
	fmt.Printf("###### main程序 start！\n")
	fmt.Printf("===================第一个golang程序===================\n")
	test01.Hello()
	fmt.Printf("\n===================变量===================\n")
	test02.TestVar()
}
