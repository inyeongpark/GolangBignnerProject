package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a1 int
	var b1 int

	//  Scan 함수를 통해 기본값을 받은 a와 b에 값을 입력력받는다.
	//  Scan 함수를 통해 읽힌 데이터는 표준입력 스트림에 저장된다.
	//  표준입력스트림은 FIFO 구조를 가지며 저장된 맨앞 데이터를 읽오 오고 그 데이터가 선언된 변수의 자료형고 맞는지 비교후 틀리다면 nil을 뱉는다. 함수를 호출
	n1, err1 := fmt.Scan(&a1, &b1)

	if err1 != nil {
		fmt.Println(n1, err1)
	} else {
		fmt.Println("정상", n1, a1, b1)
		fmt.Println("모두 정상적으로 입력받는 다면 err1값은 ", err1, "이다")
	}

	// 표준 입력 스트림에서 한줄을 읽어오는 Reader 객체를 제공
	stdin := bufio.NewReader(os.Stdin)
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		// error 발생시 줄바꿈 문자 까지 읽는다 => 표준입력 스트림에 남아있던 글자를 모두읽어서 비운다.
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
