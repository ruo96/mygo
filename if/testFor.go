package main

import "fmt"

func length(s string) int {
	fmt.Println("call length.")
	return len(s)
}

func main8() {
	s := "abcd"

	for i, n := 0, length(s); i < n; i++ {     // 避免多次调用 length 函数。
		fmt.Println(i, s[i])
	}
}

func main9() {

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}
	fmt.Printf("after for a 的值为: %d\n", a)
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
}
