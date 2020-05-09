package function

import "fmt"

func add3(a, b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2

	return
}
/**
Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
 */
func mainr1() {
	var a, b int = 1, 2
	c := add3(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
}
