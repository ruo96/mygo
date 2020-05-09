package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (s Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

/*func (stu *Stduent) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}*/

func main2() {
	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
