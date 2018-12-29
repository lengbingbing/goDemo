package main

import "fmt"

func main() {
	// 示例1。

	//make(T) 返回一个类型为 T 的初始值，


	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))   //计算切片的长度
	fmt.Printf("The capacity of s1: %d\n", cap(s1)) //计算切片的容量
	fmt.Printf("The value of s1: %d\n", s1)

	s10 := new([5]int)
	//new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0
	// 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
	fmt.Printf("The value of s10: %d\n", s10)



	s2 := make([]int, 5, 8)  //切片的长度和容量分别是 5、8
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Println()

	//S2 的容量长度 ==  切片底层数组的长度（底层数组长度不可变）
	//S2 的长度 == 可以看到窗口的大小 （只能看到 5个 [0-4] )


	// 示例2。
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println()

	//s3 切片的生成基于（数组、或切片）
	// 示例3。
	s5 := s4[:cap(s4)]
	fmt.Printf("The length of s5: %d\n", len(s5))
	fmt.Printf("The capacity of s5: %d\n", cap(s5))
	fmt.Printf("The value of s5: %d\n", s5)


}


//数组和切片的区别
//1.数组类型的长度是固定的、切片的值是可变的
// [1]string 和 [2]string 就是2个不同的类型， 长度是类型的一部分

//2.切片的长度随着、元素数量的增加而增加、不会随着减少而减少
// 切片的底层是底层数组、 切片是对底层数组的连续引用


// 引用类型、 切片、通道、函数
// 值类型  数组、结构体、基础类型

//（数组的长度和切片相等）