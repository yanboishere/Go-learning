package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	/*var 关键宇声明了两个string类型的变量s和sep。
	变量可以在声明的时候初始化。
	如果变量没有明确地初始化，它将隐式地初始化为这个类型的空值。
	*/
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] /*sep + os.Args[i] 表示将 sep 和 os.Args[i〕追加到一起
		语句
		S += sep + os.Args[i]

		是一个赋值语句，将sep 和os.Args[i]追加到旧的s上面，并且重新赋给 s，相当于语句：
		s=s+sep + os.Args[i]

		+=是一个赋值操作符
		*/
		sep = " "
	}
	fmt.Println(s)
}
