package main

import "fmt"


// Go는 main 패키지를 컴파일한 후, main() 함수를 찾아 실행한다.
// Go 프로그램을 실행하기 위해서는 go run 명령어를 사용하거나 go build 명령어를 사용하여 컴파일 후 실행한다.
func main()  {
	fmt.Println("Hello World!")
}