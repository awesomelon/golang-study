package main

import "fmt"

func check(val int) {
	// fallthrough 키워드를 사용하면 다음 case문을 실행한다.
	switch {
	case val < 0:
		fmt.Println("val < 0")
		fallthrough
	case val == 0:
		fmt.Println("val == 0")
		fallthrough
	case val > 0:
		fmt.Println("val > 0")
		fallthrough
	default:
		fmt.Println("default")
	}
}

func main() {
	const k = 1

	// 반드시 조건 블럭 시작 브레이스({)를 if문과 같은 라인에 두어야 한다.
	// 한가지 주목할 점은 if문의 조건식은 반드시 boolean 타입이어야 한다.
	if k == 1 {
		fmt.Println("One")
	} else if k == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Other")
	}

	// if문의 조건식에 변수를 선언할 수 있다.
	// 이때 조건식의 변수는 if문의 블럭 내에서만 사용할 수 있다.
	if val := 2; val < k {
		fmt.Println("One")
	} else {
		fmt.Println("Other")
	}

	var name string
	var category = 1

	// Go는 다른 언어의 case문과 달리 break를 사용하지 않는다.
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}

	fmt.Println(name)

	switch interface{}(name).(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}

	check(3)

}
