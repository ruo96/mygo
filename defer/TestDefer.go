package main

import "fmt"

func main1() {
	deferFuncParameter1()
}

func deferFuncParameter1() {
	var aInt = 1

	defer fmt.Println(aInt)

	aInt = 2
	return
}

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferFuncParameter() {
	var aArray = [3]int{1, 2, 3}

	defer printArray(&aArray)

	aArray[0] = 10
	return
}
func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}


func main() {
	//deferFuncParameter()
	deferFuncReturn()
}

