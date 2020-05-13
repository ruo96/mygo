package main

import "fmt"

func main() {

	s4 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s4, s2, s3)

	err := fmt.Errorf("这是一个错误")
	fmt.Println(err.Error())

	/**
	占位符 	说明
	%v 	值的默认格式表示
	%+v 	类似%v，但输出结构体时会添加字段名
	%#v 	值的Go语法表示
	%T 	打印值的类型
	%% 	百分号
	 */
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	/**
	占位符 	说明
	%b 	表示为二进制
	%c 	该值对应的unicode码值
	%d 	表示为十进制
	%o 	表示为八进制
	%x 	表示为十六进制，使用a-f
	%X 	表示为十六进制，使用A-F
	%U 	表示为Unicode格式：U+1234，等价于”U+%04X”
	%q 	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	 */
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)

	/**
	占位符 	说明
	%b 	无小数部分、二进制指数的科学计数法，如-123456p-78
	%e 	科学计数法，如-1234.456e+78
	%E 	科学计数法，如-1234.456E+78
	%f 	有小数部分但无指数部分，如123.456
	%F 	等价于%f
	%g 	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	%G 	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	 */
	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	/**
	字符串和[]byte
	占位符 	说明
	%s 	直接输出字符串或者[]byte
	%q 	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	%x 	每个字节用两字符十六进制数表示（使用a-f
	%X 	每个字节用两字符十六进制数表示（使用A-F）
	 */
	s := "枯藤"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	/**
	指针
	占位符 	说明
	%p 	表示为十六进制，并加上前导的0x
	 */
	a := 18
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)

	/**
	宽度标识符

	宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。
	精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下
	占位符 	说明
	%f 	默认宽度，默认精度
	%9f 	宽度9，默认精度
	%.2f 	默认宽度，精度2
	%9.2f 	宽度9，精度2
	%9.f 	宽度9，精度0
	 */
	n1 := 88.88
	fmt.Printf("默认宽度，默认精度 [%f]\n", n1)
	fmt.Printf("宽度9，默认精度[%9f]\n", n1)
	fmt.Printf("默认宽度，精度2[%.2f]\n", n1)
	fmt.Printf("宽度9，精度2[%9.2f]\n", n1)
	fmt.Printf("宽度9，精度0[%9.f]\n", n1)

	/**
	其他falg
	占位符 	说明
	’+’ 	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
	’ ‘ 	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
	’-’ 	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
	’#’ 	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
	‘0’ 	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
	 */

	s1 := "枯藤"
	fmt.Printf("[%s]\n", s1)
	fmt.Printf("[%5s]\n", s1)
	fmt.Printf("[%-5s]\n", s1)
	fmt.Printf("[%5.7s]\n", s1)
	fmt.Printf("[%-5.7s]\n", s1)
	fmt.Printf("[%5.2s]\n", s1)
	fmt.Printf("[%05s]\n", s1)
	
}
