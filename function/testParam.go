package function

import "fmt"

/* 定义相互交换值的函数 */
func swap(x, y *int) {
	var temp int

	temp = *x /* 保存 x 的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y*/

}

func swap2(x, y int) {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y   /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

}

/*func main() {
	var a, b int = 1, 2

	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址

	//swap(&a, &b)
	swap2(a, b)

	fmt.Println(a, b)
}*/

func myfunc(args...int) int{    //0个或多个参数
	var x int
	for _,i := range args{
		x += i
	}
	return x
}

func add(a int, args...int) int {    //1个或多个参数
	var x int
	for _,i := range args{
		x += i
	}
	return x + a
}

func add1(a int, b int, args...int) int {    //2个或多个参数
	var x int
	for _,i := range args{
		x += i
	}
	return x + a + b
}

func mainp1()  {
	x := myfunc(1,2,3)
	fmt.Println(x)

	y := add(4,5,6,7)
	fmt.Println(y)

	z := add1(8,9,10,11)
	fmt.Println(z)
}
