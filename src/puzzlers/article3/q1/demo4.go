package main

import (
"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello(name)



}

//同一个目录下的源码文件声明成同一个代码包
//包名和目录名称可以不一致，按目录名生成可执行文件