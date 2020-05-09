package main

import (
	"math"
	"os"
	"strings"
)
import "fmt"

const (
	n0 = 2
	n1 = iota
	n2
	_  = 200
	n3 = iota
	n4
)

const (
	x, y = 1, 2
	a, b = iota + 1, iota + 2
	c, d
	e, f1
)

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main3() {
	/*var arr1 [5]int
	// 数据传参是必须要传地址的
	printArr(&arr1)
	fmt.Println(arr1)*/

	//这样就会改变数组的值
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

func main2() {
	fmt.Println(str[3], str[4])

	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func main1() {
	buf := make([]byte, 1024)
	f, _ := os.Open("e:\\file\\test.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	fmt.Println(n0, n1, n2, n3, n4)
	fmt.Println(a, b, c, d, e, f1)

	f2 := math.MaxFloat32
	fmt.Println(f2)

	f3 := math.MaxFloat64
	fmt.Println(f3)

	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1, s2)

	index := strings.Index(s1, "ll")
	fmt.Println(index)

	traversalString()
	changeString()
	sqrtDemo()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
