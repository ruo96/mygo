package function

import "fmt"

/**
  • 无需声明原型。
  • 支持不定 变参。
  • 支持多返回值。
  • 支持命名返回参数。
  • 支持匿名函数和闭包。
  • 函数也是一种类型，一个函数可以赋值给变量。

  • 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
  • 不支持 重载 (overload)
  • 不支持 默认参数 (default parameter)。
 */

/*func test(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func main() {
	i,s := test(1,2,"hello, %d")
	fmt.Println(i,s)
}*/

func test(fn func() int) int {
	return fn()
}
// 定义函数类型。
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func mainf1() {
	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	fmt.Println(s1, s2)
}

