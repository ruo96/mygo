package main

import "fmt"

func main1() {
	x := 0

	// if x > 10        // Error: missing condition in if statement
	// {
	// }

	/**不支持三元操作符(三目运算符) "a > b ? a : b"。*/
	if n := "abc"; x > 0 { // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
		fmt.Println(n[2])
		fmt.Printf("%c", n[2])
	} else if x < 0 { // 注意 else if 和 else 左大括号位置。
		fmt.Println(n[1])
		fmt.Printf("%c", n[1])
	} else {
		fmt.Println(n[0])
		fmt.Printf("%c", n[0])
	}
}

func main2() {
	/* 定义局部变量 */
	var a int = 10
	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func main3() {
	/* 局部变量定义 */
	var a int = 100
	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func main4() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}
