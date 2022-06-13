// 같은 폴더에 위치한 .go 파일은 모두같은 패키지에 포함되고 폴더명 = 패키지명이다.

// 패키지 선언: 이 코드가 어떤 패키지에 속하는지 알려주는 기능 (main package에 속해있음 을 알려준다.)
// package main은 m패키지 프로그램의 시작점을 포함하는 특별한 패키지
package main

//특정 패키지에서 제공하는 기능을 사용하고 싶을때 호출(import)
// fmt는 go의 기본언어 패키지
import "fmt"

func main() {
	// 변수선언의 기본형태는 아래와같다
	// var(변수선언키워드) hello_text(변수명) string(변수타입) = "Hello World!"(변수값)
	var hello_text string = "Hello World!"
	var b int

	// 선언 대입문
	hello_text2 := "Hello World2!!"

	fmt.Println((hello_text))
	fmt.Println(b)
	fmt.Println(hello_text2)
}
