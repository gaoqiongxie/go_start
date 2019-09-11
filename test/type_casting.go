//语言类型转换
//用于将一种数据类型的变量转换为另一种类型的变量
//type_name(expression)
package main

import "fmt"

//整型转化为浮点型，并计算结果，将结果赋值给浮点型变量
func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
}