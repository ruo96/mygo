package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[6:9]
	// 自动扩容到4 原来的两倍  从2开始，4，8，16，32
	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))
}

/*golang slice data[:6:8] 两个冒号的理解

常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为4（6-9）

另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8

a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x  */

func main1() {
	str := "Hello world"
	s := []byte(str) //中文字符需要用[]rune(str)
	fmt.Println("%v", s)
	s[6] = 'G'
	fmt.Println(s)
	s = s[:8]
	fmt.Println(s)
	s = append(s, '!')
	fmt.Println(str)
	str = string(s)
	fmt.Println(str)
}
