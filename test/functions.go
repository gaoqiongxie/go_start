//go-functions 语言函数
package main

import (
	"fmt"
	"math"
)

func main() {
	/* 定义局部变量 */
	// var a int = 100
	// var b int = 200
	// var ret int

	// /* 调用函数并返回最大值 */
	// ret = max(a, b)

	// fmt.Printf( "最大值是 : %d\n", ret )

	// a, b := swap("Google", "Runoob")
 //    fmt.Println(a, b)

	// swapTest()	
	// getSquareRootTest()

	// getSequenceTest()

	getAreaTest()
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result 
}

// func swap(x, y string) (string, string) {
//    return y, x
// }

func swap(x *int, y *int) {
	var temp int
	temp = *x    /* 保持 x 地址上的值 */
	*x = *y      /* 将 y 值赋给 x */
	*y = temp    /* 将 temp 值赋给 y */
}

func swapTest(){
	var a int = 100
	var b int= 200

	fmt.Printf("交换前，a 的值 : %d\n", a )
	fmt.Printf("交换前，b 的值 : %d\n", b )

	/* 调用 swap() 函数
	* &a 指向 a 指针，a 变量的地址
	* &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	fmt.Printf("交换后，a 的值 : %d\n", a )
	fmt.Printf("交换后，b 的值 : %d\n", b )
}

//语言函数作为实参
//函数作为另外一个函数的实参
func getSquareRootTest(){
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
	  return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

}

//语言函数闭包  即匿名函数，可在动态编程中使用
func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i  
	}
}

func getSequenceTest(){
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()  

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()  
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

}

//方法 就是一个包含了接受者的函数
/* 定义结构体 */
type Circle struct {
  radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}

func getAreaTest(){
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}
