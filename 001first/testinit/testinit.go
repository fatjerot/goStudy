// 测试init方法

package testinit

import "fmt"

func init() {
	fmt.Printf("###### 先从里往外执行init方法，然后再执行main方法。\n")
	fmt.Printf("init inner inner\n")
}

// Testinit 必须有被外部引用的方法才会初始化
func Testinit() {
	fmt.Printf("testinit\n")
}
