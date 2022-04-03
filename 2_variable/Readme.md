# 개요
* 변수선언

# 내용
* var키워드로 변수를 선언할 수 있습니다.
```go
package main

import "fmt"

func main() {
	var [변수이름]
}
```

* go는 변수 타입을 검사합니다. 그러므로 변수 선언시 타입을 정의하지 않으면 컴파일 오류가 발생합니다.
```go
package main

import "fmt"

func main() {
	var a; <-- 오류

	fmt.Println(a)
}
```

* 타입은 변수 이름 뒤에 입력합니다. 값을 설정하지 않으면 0으로 초기화 됩니다.
```go
package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
}
```

* 선언과 동시에 값을 초기화 하면 타입을 설정하지 않아도 됩니다.
```go
package main

import "fmt"

func main() {
	var a int
	var b string = "hello"

	fmt.Println(a)
	fmt.Println(b)
}
```