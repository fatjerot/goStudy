// go基础学习测试
package main

import (
	"fmt"
	first "go_study_basic/001first"
)

func init() {
	fmt.Printf("init\n")
}

func main() {
	fmt.Printf("###### main程序 start！\n")
	fmt.Printf("===================第一个golang程序===================\n")
	first.Hello()
}
