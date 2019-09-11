//gp loops 循环
package main 
import "fmt"

func main() {
	// forTest()
	// primeTest()
	// gotoTest()
   multiTest()
}

func forTest(){
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5} 

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}   
}

//使用循环嵌套来输出 2 到 100 间的素数：
func primeTest(){
	/* 定义局部变量 */
   var i, j int

   for i=2; i < 100; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break; // 如果发现因子，则不是素数
         }
      }
      if(j > (i/j)) {
         fmt.Printf("%d  是素数\n", i);
      }
   }  
}


/*
	Go 语言的 goto 语句可以无条件地转移到过程中指定的行
	goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
	在结构化程序设计中一般不主张使用 goto 语句
*/
func gotoTest(){
   /* 定义局部变量 */
   var a int = 10

   /* 循环 */
   LOOP: for a < 20 {
      if a == 15 {
         /* 跳过迭代 */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++     
   }  
}

//九九乘法表
func multiTest(){
   for y := 1; y <= 9; y++{
      for x := 1; x <= y; x++{
         fmt.Printf("%d*%d=%d  ", x, y, x*y)
      }
      fmt.Println()
   }
}