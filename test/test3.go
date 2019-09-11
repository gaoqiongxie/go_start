/** 变量的作用域*/
package main 

import(
	"fmt"
)

func main() {
    // x := "hello!"
    // for i := 0; i < len(x); i++ {
    //     x := x[i]
    //     if x != '!' {
    //         x := x + 'A' - 'a'
    //         fmt.Printf("%c", x) // "HELLO" (每次迭代一个字母)
    //     }
    // }

    // for _, x := range x{
    // 	x := x + 'A' - 'a'
    // 	fmt.Printf("%c", x) // HELLO
    // }

    if x := f(); x == 0 {
	    fmt.Println(x)
	} else if y := g(x); x == y {
	    fmt.Println(x, y)
	} else {
	    fmt.Println(x, y)
	}
	fmt.Println(x, y) // 编译错误: x 和 y 未定义
}