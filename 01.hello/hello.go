// 같은 폴더에 위치한 .go 파일은 모두같은 패키지에 포함되고 폴더명 = 패키지명이다.

// 패키지 선언: 이 코드가 어떤 패키지에 속하는지 알려주는 기능 (main package에 속해있음 을 알려준다.)
// package main은 m패키지 프로그램의 시작점을 포함하는 특별한 패키지
package main

//특정 패키지에서 제공하는 기능을 사용하고 싶을때 호출(import)
import "fmt"

func main() {
	fmt.Println(("Hello World!"))
}
