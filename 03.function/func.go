package main

import "fmt"

//  int a int b는 함수가 외부로부터 입력받는 값 매개변수
//
func Add(a int, b int) int {
	return a + b
}

func main() {
	//  3과 6은 인수로 합수를 호출 할때 임력하는 값 argment
	c := Add(3, 6)
	fmt.Println(c)
}
