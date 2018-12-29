package main

import "fmt"

func main() {
	// 示例1。
	{
		// MyString 是 string 类型的别名类型
		type MyString = string    // （别名类型，相同类型）
		// 内置的 byte = uint8 、 rune = int32

		type MyString2 string    //新类型  类型在定义 、 string 是 MyString2 的潜在类型

		str := "BCD"
		myStr1 := MyString(str)
		myStr2 := MyString("A" + str)
		fmt.Printf("%T(%q) == %T(%q): %v\n",
			str, str, myStr1, myStr1, str == myStr1)
		fmt.Printf("%T(%q) > %T(%q): %v\n",
			str, str, myStr2, myStr2, str > myStr2)
		fmt.Printf("Type %T is the same as type %T.\n", myStr1, str)

		strs := []string{"E", "F", "G"}
		myStrs := []MyString(strs)
		fmt.Printf("A value of type []MyString: %T(%q)\n",
			myStrs, myStrs)
		fmt.Printf("Type %T is the same as type %T.\n", myStrs, strs)
		fmt.Println()
	}
	// 示例2。
	{
		type MyString string
		str := "BCD"
		myStr1 := MyString(str)
		myStr2 := MyString("A" + str)
		_ = myStr2
		//fmt.Printf("%T(%q) == %T(%q): %v\n",
		//	str, str, myStr1, myStr1, str == myStr1) // 这里的判等不合法，会引发编译错误。
		//fmt.Printf("%T(%q) > %T(%q): %v\n",
		//	str, str, myStr2, myStr2, str > myStr2) // 这里的比较不合法，会引发编译错误。
		fmt.Printf("Type %T is different from type %T.\n", myStr1, str)

		strs := []string{"E", "F", "G"}
		var myStrs []MyString
		//myStrs := []MyString(strs) // 这里的类型转换不合法，会引发编译错误。
		//fmt.Printf("A value of type []MyString: %T(%q)\n",
		//	myStrs, myStrs)
		fmt.Printf("Type %T is different from type %T.\n", myStrs, strs)
		fmt.Println()
	}
	// 示例3。
	{
		type MyString1 = string
		type MyString2 string
		str := "BCD"
		myStr1 := MyString1(str)
		myStr2 := MyString2(str)
		myStr1 = MyString1(myStr2)
		myStr2 = MyString2(myStr1)

		myStr1 = str
		//myStr2 = str // 这里的赋值不合法，会引发编译错误。
		//myStr1 = myStr2 // 这里的赋值不合法，会引发编译错误。
		//myStr2 = myStr1 // 这里的赋值不合法，会引发编译错误。
	}
}