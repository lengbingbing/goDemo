package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
}


// 1.变量的重新声明、必须和原来的类型一致
// 2.重生明只能是同一个代码块中
// 3.只能是短代码声明才能发生
// 4.被预知的变量必须是多个,并且有一个是新的变量

