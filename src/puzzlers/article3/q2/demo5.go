package main

import (
	"flag"

	"puzzlers/article3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)

	//lib5 为限定符
}


// go install puzzlers/article3/lib  生成归档文件 （pkg 目录中）


// (1) import 是根据目录名导入
// (2) 限定符，是根据 packge 名称导入
// (3) 手写字母大写 、可以供包外程序调用
//（4） internal 模块化私有




