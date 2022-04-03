# 개요
* print함수

# 내용
* print함수는 fmt모듈을 사용합니다. print은 개행이 없고 println은 출력 후 개행이 존재합니다.
```go
package main

import "fmt"

func main() {
	var a int = 3
	var b string = "hello"

	fmt.Println(a)
	fmt.Println(b)
}
```

* format기능은 printf를 사용하면 됩니다. 범용적으로 %v 포맷을 사용합니다.
```go
package main

import "fmt"

func main() {
	var a int = 3
	var b string = "hello"

	fmt.Printf("a : %v, b:%v\n", a, b)
}
```