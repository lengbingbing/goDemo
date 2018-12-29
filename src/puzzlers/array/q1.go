package main

import "fmt"

func main() {

	//var identifier [len]type
	// var 数组名称  [长度]数组类型
	//var arr1 [5]int




	var arr1 = new([5]int)
	for i:=0;i< len(arr1);i++{

		arr1[i] = i * 2
	}
	arr2 := arr1
	arr2[2] = 100


	for i:=0;i<len(arr1);i++{

		fmt.Printf("array %d is %d\n",i,arr1[i])
	}
	for i:=0;i<len(arr2);i++{

		fmt.Printf("array %d is %d\n",i,arr2[i])
	}
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}

