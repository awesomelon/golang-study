// 1. 부울린 타입
// bool

// 2 .문자열 타입
// string: string은 한번 생성되면 수정될 수 없는 Immutable 타입임

// 3. 정수형 타입
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr

// 4. Float 및 복소수 타입
// float32 float64 complex64 complex128

// 5. 기타 타입
// byte: uint8과 동일하며 바이트 코드에 사용
// rune: int32과 동일하며 유니코드 코드포인트에 사용한다

package main

import (
	"fmt"
)

func main() {
	// Back Quote로 둘러 싸인 문자열은 Raw String Literal이라고 부른다. 이 안에 있는 문자열은 별도로 해석되지 않고,
	// Raw String 그대로의 값을 갖는다.
	rawLiteral := `아리랑\n
아리랑\n
아라리요`

	// 이중인용부호로 둘러 싸인 문자열은 Interpreted String Literal이라고 부른다. 이 안에 있는 이스케이프 코드는 해석된다.
	interLiteral := "아리랑아리랑\n아라리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

	// Type Conversion
	// 하나의 데이터 타입에서 다른 데이터 타입으로 변환하는 것을 Type Conversion이라고 한다.
	// Go에서 Type Conversion은 명시적으로 지정해 주어야 한다.
	// T(v)와 같이 표현하며, T는 변환하고자 하는 타입, v는 변환될 값(value)을 나타낸다.
	// Go에서 타입간 변환은 명시적으로 지정해주어야 한다. 
	// 예를 들어 정수형 int에서 unit로 변환할 때, 암묵적으로 변환되지 않는다.
	
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}