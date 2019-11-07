package test02

import "fmt"

func TestCreateFromSlice()  {
	test := []int{1,2,3,4,5}
	test1 := test[1:]
	test2 := test[:3]
	test3 := test[2:4]
	test4 := test[2:3:5]

	fmt.Println("test:", test, "; cap:", cap(test))
	fmt.Println("test1:", test1, "; cap:", cap(test1))
	fmt.Println("test2:", test2, "; cap:", cap(test2))
	fmt.Println("test3:", test3, "; cap:", cap(test3))
	fmt.Println("test4:", test4, "; cap:", cap(test4))
}