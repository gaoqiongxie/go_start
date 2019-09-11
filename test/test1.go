/** 变量赋值*/
package main 

import(
 "fmt"
)

func main() {
	var a int = 100
	var b int = 200
	// var t int
	// t = a
	// a = b
	// b = t
	// fmt.Println(a, b)


	// a = a ^ b
	// b = b ^ a
	// a = a ^ b
	// fmt.Println(a, b)


	b, a = a, b
	fmt.Println(a, b)
}