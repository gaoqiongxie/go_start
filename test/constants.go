// go_constants 常量
package main

import (
    "fmt"
	// "unsafe"
)
// const(
//     a = iota
//     b = iota
//     c = iota
// )
// const (
//     a = iota   //0
//     b          //1
//     c          //2
//     d = "ha"   //独立值，iota += 1
//     e          //"ha"   iota += 1
//     f = 100    //iota +=1
//     g          //100  iota +=1
//     h = iota   //7,恢复计数
//     i          //8
// )
// func main() {
//     const LENGTH int = 10
// 	// const WIDTH int = 5   
// 	// var area int
// 	// const a, b, c = 1, false, "str" //多重赋值

// 	// area = LENGTH * WIDTH
// 	// fmt.Printf("面积为 : %d", area)
// 	// println()
// 	 fmt.Println(a,b,c,d,e,f,g,h,i)
// }

const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}