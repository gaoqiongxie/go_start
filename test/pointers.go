// go pointers 语言指针
package main 

import "fmt"

func main(){
	// test2()

	// test3()

	// arrayOfPointersTest()
	// pointer_to_pointerTest2()
	passing_pointers_to_functions_Test()
} 

func test1(){
	var a int = 10
	fmt.Printf("变量的地址：%x\n", &a)
}

/* 1.定义指针变量 2.为指针变量赋值 3.访问指针变量中指向地址的值*/
func test2(){
	var a int= 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}

/* Go 空指针*/
func test3(){
	var  ptr *int

    fmt.Printf("ptr 的值为 : %x\n", ptr )

    if ptr != nil {
    	fmt.Printf("ptr is nil")
    }

}

/* 指针数组*/
const MAX int = 3
func arrayOfPointersTest(){
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++{
		ptr[i] = &a[i]//整数地址赋值给指针数组
	}

	for i = 0; i < MAX; i++{
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

/*指向指针的指针*/
/*如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。*/
func pointer_to_pointerTest(){
	var a int
	var ptr *int
	var pptr **int  // 向指针的指针变量值需要使用两个 * 号

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

func pointer_to_pointerTest2(){
	var a int = 1
	var ptr1 *int = &a
	var ptr2 **int = &ptr1
	var ptr3 **(*int) = &ptr2 // 也可以写作：var ptr3 ***int = &ptr2
	// 依次类推
	fmt.Println("a:", a) //1
	fmt.Println("ptr1:", ptr1) //a地址
	fmt.Println("ptr2:", ptr2) //a地址
	fmt.Println("ptr3:", ptr3) //a地址
	fmt.Println("*ptr1:", *ptr1) //1
	fmt.Println("**ptr2:", **ptr2) //1
	fmt.Println("**(*ptr3):", **(*ptr3)) // 1 也可以写作：***ptr3
}

/* 向函数传递指针函数*/
func passing_pointers_to_functions_Test(){
	/* 定义局部变量 */
	var a int = 100
	var b int= 200

	fmt.Printf("交换前 a 的值 : %d\n", a )
	fmt.Printf("交换前 b 的值 : %d\n", b )

	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* &b 指向 b 变量的地址
	*/
	swap(&a, &b);

	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址的值 */
   *x = *y      /* 将 y 赋值给 x */
   *y = temp    /* 将 temp 赋值给 y */
}