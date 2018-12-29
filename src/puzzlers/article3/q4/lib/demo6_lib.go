package lib

import (
	"os"
	in "puzzlers/article3/q4/lib/internal"
	//包别名
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}