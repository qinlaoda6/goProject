package function

import "fmt"

// TwoNumAdd
// @description 2个数相加
// @param num1
// @param num2
// @return int
func TwoNumAdd(num1 int, num2 int) int {
	num3 := num1 + num2
	return num3
}

// Test
// @description 递归调用
// @param n
func Test(n int) {
	if n > 2 {
		n--
		Test(n)
	}
	fmt.Println("n=", n)
}
func Test2(n int) interface{} {
	if n > 2 {
		n--
		Test(n)
	}
	fmt.Println("n=", n)
	return n
}
