package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
}

// 1.变量声明、同一个代码块中、类型必须相同
// 2.可重命名变量  ， 不同的代码块（作用域中） 类型可以不通、覆盖
