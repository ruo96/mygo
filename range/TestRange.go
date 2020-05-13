package main

import "fmt"

func bianlishuzu() {

	shuzu := [5]int{1,2,3,4,5}
	for index, value := range shuzu{
		fmt.Printf("数组的index：%v,  值：%v \n", index, value)
	}
}

func bianlimap(){
	mymap := map[int]string{
		1:"w", 2:"r", 3:"h",

	}
	for key,value := range mymap{
		fmt.Printf("Map的key：%v,  值：%v \n", key, value)
	}
}

/**
能够正常结束。循环内改变切片的长度，不影响循环次数，循环次数在循环开始前就已经确定了
 */
func dongtaibianli() {
	v := []int{1, 2, 3}
	for i:= range v {
		v = append(v, i)
	}
}

func main(){
	//bianlishuzu()
	//bianlimap()
	dongtaibianli()
}
