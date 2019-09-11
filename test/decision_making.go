//go_decision_making 条件语句
package main

import (
	"fmt"
	"time"
)

func main() {
  // ifTest()
  // switchTest()
  // typeSwitchTest()
  // fallthroughTest()
  // selectTest()
	// selectTest2()
	gotoTest()
}

func ifTest(){
	 /* 定义局部变量 */
	var a int = 10

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
	   /* 如果条件为 true 则执行以下语句 */
	   fmt.Printf("a 小于 20\n" )
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func switchTest(){
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"  
	}

	switch {
		case grade == "A" :
			fmt.Printf("优秀!\n" )     
		case grade == "B", grade == "C" :
			fmt.Printf("良好\n" )      
		case grade == "D" :
			fmt.Printf("及格\n" )      
		case grade == "F":
			fmt.Printf("不及格\n" )
		default:
			fmt.Printf("差\n" );
	}
	fmt.Printf("你的等级是 %s\n", grade );
}

func typeSwitchTest(){
	var x interface{}

	switch i := x.(type) {
		case nil:   
			fmt.Printf(" x 的类型 :%T",i)                
		case int:   
		    fmt.Printf("x 是 int 型")                       
		case float64:
		    fmt.Printf("x 是 float64 型")           
		case func(int) float64:
		    fmt.Printf("x 是 func(int) 型")                      
		case bool, string:
		    fmt.Printf("x 是 bool 或 string 型" )       
		default:
		    fmt.Printf("未知型")     
	}   
}

// fallthrough 会强制执行后面的 case 语句，不会判断下一条 case 的表达式结果是否 true
func fallthroughTest(){
	switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }
}

/*
select 类似用于通信的 switch 语句，每个case 必须有一个通信操作，要么是发送要么是接收
select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。
一个默认的子句应该总是可运行的。
https://www.runoob.com/go/go-select-statement.html
*/
func  selectTest() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
		case i1 = <-c1:
			fmt.Printf("received ", i1, " from c1\n")
		case c2 <- i2:
			fmt.Printf("sent ", i2, " to c2\n")
		case i3, ok := (<-c3):  // same as: i3, ok := <-c3
			if ok {
			fmt.Printf("received ", i3, " from c3\n")
			} else {
			fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication\n")
	}    
}

func Chann(ch chan int, stopCh chan bool) {
    var i int
    i = 10
    for j := 0; j < 10; j++ {
        ch <- i
        time.Sleep(time.Second)
    }
    stopCh <- true
}
//select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。
func selectTest2(){
	ch := make(chan int)
    c := 0
    stopCh := make(chan bool)

    go Chann(ch, stopCh)

    for {
        select {
        case c = <-ch:
            fmt.Println("Recvice", c)
            fmt.Println("channel")
        case s := <-ch:
            fmt.Println("Receive", s)
        case _ = <-stopCh:
            goto end
        }
    }
    end:
}

func gotoTest(){
	if true{
		goto onExit
	}

	onExit:
		fmt.Println("true")
}

