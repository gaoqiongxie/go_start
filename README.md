##[Go](http://c.biancheng.net/golang/syntax/)
主要目标：兼具 Python 等动态语言的开发速度和 C/C++等编译型语言的性能与安全性
用途：访问底层操作系统，提供了强大的网络编程和并发编程支持，可以进行网络编程、系统编程、并发编程、分布式编程
优势：旨在不损失应用程序性能的情况下降低代码的复杂性，具有“部署简单、并发性好、语言设计良好、执行性能好”等优势

###Go 是编译型语言 
GO 使用编译器来编译代码。
编译器将源代码编译成二进制（或字节码）格式；
在编译代码时，编译器检查错误、优化性能并输出可在不同平台上运行的二进制文件。
要创建并运行 Go 程序，程序员必须执行如下步骤:
1.使用文本编辑器创建 Go 程序；
2.保存文件；
3.编译程序；
4.运行编译得到的可执行文件。

Go 快速编译，高效执行，易于开发
国际化，使用UTF-8编码，跨平台

Go 没有类和继承的概念，通过接口的概念来实现多态性 

####语言特性
##### 语法简单
- 将“++”、“--”从运算符降级为语句
- 保留指针，但默认阻止指针运算，
- 将切片和字典作为内置类型，从运行时的层面进行优化

##### 并发模型
- 关键字 Goroutine，类协程处理并发单元，又在运行时层面做了更深度的优化处理
- 搭配 channel，实现 CSP 模型，将并发单元间的数据耦合拆解开来，各司其职

##### 内存分配
- tcmalloc 
	使用cache为当前执行线程提供无锁分配，
	多个 central 在不同线程间平衡内存单元复用，
	在更高层次里，heap 管理大块内存，用以切分成不同等级的复用内存块
	快速分配和二级内存平衡机制
- 除偶尔因性能问题而被迫采用对象池和自主内存管理外，基本无须参与内存管理操作

##### 垃圾回收
- 从并发清理，到降低 STW 时间，引入三色标记和写屏障等

##### 静态链接
- 只须编译后的一个可执行文件，无须附加任何东西就能部署将运行时、依赖库直接打包到可执行文件内部，简化了部署和发布操作，无须事先安装运行环境和下载诸多第三方库
- buildmode 未完工
	
##### 标准库
- 提供了清晰的构建模块和公共接口，包含 I/O 操作、文本处理、图像、密码学、网络和分布式应用程序等
- 支持许多标准化的文件格式和编解码协议
- 编译器，通过词法器扫描源码，使用语法树获得源码逻辑分支等

##### 工具链
比较完善：编译、格式化、错误检查、帮助文档、第三方包下载、更新
内置完整测试框架，其中包括单元测试、性能测试、代码覆盖率、数据竞争，
以及用来调优的 pprof
还可通过环境变量输出运行时监控信息，尤其是垃圾回收和并发调度跟踪

#### 为并发而生
Go 是在多核和网络化的时代背景下诞生的原生支持并发的编程语言
并发基于 goroutine ，类似于线程，多个 goroutine 中，Go 语言使用 channel 进行通信，通道是一种内置的数据结构，可以让用户在不同的 goroutine 之间同步发送具有类型的消息，这让编程模型更倾向于在 goroutine 之间发送消息，而不是让多个 goroutine 争夺同一个数据的使用权，即*Go 语言通过通道可以实现多个 goroutine 之间的内存共享*
程序可以将并发的环节设计为生产者模式和消费者模式，将数据放入通道
	
### 哪些项目使用 Go 语言开发
- Docker
- go 语言
- Kubernetes
- etcd
- beego
- martini
- codis
- delve

### 安装工具包

### Go 语言基本语法
####Go 语言的基本类型：
- bool
- string
- int、int8、int16、int32、int64
- uint、uint8、uint16、uint32、uint64、uintptr
- byte // uint8 的别名
- rune // int32 的别名 代表一个 Unicode 码
- float32、float64
- complex64、complex128


当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，切片、函数、指针为 nil 等。所有的内存在 Go 中都是经过初始化的。

骆驼命名法

####变量声明的形式
- 标准格式
	var 变量名 变量类型
- 批量格式
~~~
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
~~~

- 简短格式
	+ 名字 := 表达式
~~~
func main() {
   x:=100
   a,s:=1, "abc"
}
~~~
	+ 注意
		- 定义变量，同时显示初始化
		- 不能提供数据类型
		- 只能用在函数内部


####变量声明的初始化
- 标准格式
	var 变量名 类型 = 表达式
~~~
var hp int = 100
~~~
- 编译器推导类型的格式
~~~
var hp = 100
~~~
- 短变量声明并初始化
~~~
hp := 100
~~~
	+ 注意
		- 由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误
		- 在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中
	
####多个变量同时赋值
~~~
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
~~~

#### 匿名变量
匿名变量的特点是一个下画线**“_”**，“_”本身就是一个特殊的标识符，被称为空白标识符
~~~
/** 匿名变量*/
package main 
import(
	"fmt"
)

func GetData() (int, int) {
    return 100, 200
}

func main(){
	a, _ := GetData()
	_, b := GetData()

	fmt.Println(a, b)
}
~~~
*匿名变量不占用命名空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。*

####变量作用域
~~~
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
~~~

###Go语言类型的本质
这个类型的本质是什么？保持传递的一致性很重要，这个背后的原则是，不要只关注某个方法是如何处理这个值，而是要关注这个值的本质是什么
- 如果要撞见一个新值，该类型的方法就使用值接收者，即按照值传递
- 如果要修改当前值，就使用指针接收者，即按指针传递（指针共享数据）

####内置类型
是由语言提供的一组类型，如：数值类型、字符串类型和布尔类型，这些类型本质是原始的类型，因此，对这些值进行增加或者删除的时候，会创建一个新值

####引用类型
如：切片、映射、通道、接口和函数类型，声明这些类型的变量时，创建的变量被称作标头（header）值，从技术细节来说，字符串也是一种引用类型。
每个引用类型创建的标头值是包含一个指向底层数据结构的指针。每个引用类型还包含一组独特的字段，用于管理底层数据结构。因为标头值是为复制而设计的，所以永远不需要共享一个引用类型的值。标头值里包含一个指针，因此通过复制来传递一个引用类型的值的副本，本质上就是在共享底层数据结构。

####结构类型
可以用来描述一组数据值，这组值的本质即可以是原始的，也可以是非原始的。如果决定在某些东西需要删除或者添加某个结构类型的值时该结构类型的值不应该被更改，那么需要遵守之前提到的内置类型和引用类型的规范。

###Go语言类型（整数类型）
- 有符号 int8、int16、int32 和 int64
- 无符号 uint8、uint16、uint32 和 uint64

- 大多数情况下，只需要 int 一种整型即可，它可以用于循环计数器、数组、切片索引，以及任何通用目的的整型运算符，通常 int 类型的处理速度也是最快的。

- Unicode 字符 rune 类型是和 int32 等价的类型
- byte 也是 uint8 类型的等价类型，byte 类型一般用于强调数值是一个原始的数据而不是一个小的整数。
- 无符号的整数类型 uintptr，没有指定具体的 bit 大小但是足以容纳指针，底层编程时才需要，特别是 Go语言和 C语言函数库或操作系统接口相交互的地方
- 不管它们的具体大小，int、uint 和 uintptr 是不同类型的兄弟类型，其中 int 和 int32 也是不同的类型，即使 int 的大小也是 32bit，在需要将 int 当作 int32 类型的地方需要一个显式的类型转换操作，反之亦然。
- 有符号整数采用 2 的补码形式表示，也就是最高 bit 位用来表示符号位，一个 n-bit 的有符号数的值域是从 $-2^{n-1}$ 到 $2^{n-1}-1$，int8 类型整数的值域是从 -128 到 127
- 无符号整数的所有 bit 位都用于表示非负数，值域是 0 到 $2^n-1$，uint8 类型整数的值域是从 0 到 255

####哪些情况下使用 int 和 uint
- 实际使用中，切片或 map 的元素数量等都可以用 int 来表示
- 在二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用 int 和 uint

###Go语言浮点类型
float32 和 float64
```
	//float32 的有效 bit 位只有 23 个 ，其他的 bit 位用于指数和符号
	//当整数大于 23bit 能表达的范围时，float32 的表示将出现误差
	var f float32 = 16777216
	fmt.Println(f == f+1)

	//很小或者很大的数最好用科学计数法书写
	const Avogadro = 6.02214129e23 //阿伏伽德罗常数
	const Planck = 6.62606957e-34//普朗克常数

	//打印 %g %e %f
	for x := 0; x < 8; x++{
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	//
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //0 -0 +Inf -Inf NaN

	//在浮点数中，NaN、正无穷大和负无穷大都不是唯一的，每个都有非常多种的 bit 模式表示。
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) //false false false
```
###Go语言复数
complex64（32位实数和虚数）
complex128（64位实数和虚数）
~~~
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x*y) //(-5+10i)
	//real 和 imag 函数分别返回复数的实部和虚部
	fmt.Println(real(x*y))
	fmt.Println(imag(x*y))

	//复数使用 re+imI 来表示 即实数+虚数 I代表根号负1
	fmt.Println(1i * 1i)

	//如果re 和 im 的类型君威float32 那么complex64的复数C可以通过以下方式获得
	//c = complex(re, im) 
~~~

###bool类型
```
//布尔值并不会隐式转换为数字值 0 或 1，反之亦然。必须使用一个显式的 if 语句辅助转换：
// 如果b为真，btoi返回1；如果为假，btoi返回0
func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}

//数字到布尔型的逆转换则
func itob(i int) bool { return i != 0 }

//Go 语言中不允许将整型强制转换为布尔型
var n bool
fmt.Println(int(n) * 2)//cannot convert n (type bool) to type int

//布尔型无法参与数值运算，也无法与其他类型进行转换。

```

###字符串
####go支持的字面值
1.解释字符串
	+ \n
	+ \r 
	+ \t
	+ \u
	+ \\

2.非解释字符串
该类字符串使用反引号‘`’，支持换行

#####获取字符串某个字节的地址的行为是非法的

#####拼接 “+”
由于编译器行尾自动补全分号的缘故，加号 + 必须放在第一行。

#####字符串实现基于 UTF-8 编码
通过 runne 类型，可以方便的对每个 UTF-8 字符串进行访问，也支持按传统的ASCII码方式进行逐字符访问


#####定义多行字符串
在源码中嵌入多行字符串时，必须使用‘`’，反引号间的换行被称作字符串中的换行，所有的转移字符均无效。
一般用于内嵌源码和内嵌数据等


#####计算长度
内建函数len()，可以用来获取切片、字符串、通道等的长度，返回值的类型为 int，表示字符串的ASCII字符个数或字节长度
UTF-8 每个中文占3个字节 
fmt.Println(len("忍者")) //6

uneCountInString() 函数，统计Uncode字符数量   
fmt.Println(utf8.RuneCountInString("忍者")) //2

#####格式化输出
表：字符串格式化时常用动词及功能

动词 | 功能
-|-
%v | 按值的本来值输出
%+v | 在 %v 基础上，对结构体字段名和值进行展开
%#v | 输出 Go 语言语法格式的值
%T | 输出 Go 语言语法格式的类型和值
%% | 输出 % 本体
%b | 整型以二进制方式显示
%o | 整型以八进制方式显示
%d | 整型以十进制方式显示
%x | 整型以十六进制方式显示
%X | 整型以十六进制、字母大写方式显示
%U | Unicode 字符
%f | 浮点数
%p | 指针，十六进制方式显示

```
package main 
import (
	"unicode/utf8"
	"encoding/base64"
	"fmt"
	"strings"
	"bytes"
)

func main(){
	//字符串拼接符
	//由于编译器行尾自动补全分号的缘故，加号 + 必须放在第一行。
	//"+"
	str := "Beginning of the string " + 
	"second part of the string"
	fmt.Println(str)

	//也可以是"+="
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”

	//字符串长度
	//ASCII 字符串长度使用 len() 函数
	//这里的差异是由于 Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节，因此使用 len() 获得两个中文文字对应的 6 个字节。
	fmt.Println(len("genji is a ninja"))//16
	fmt.Println(len("忍者"))//6

	//Unicode 字符串长度使用 utf8.RuneCountInString() 函数
	fmt.Println(utf8.RuneCountInString("忍者"))//2
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))//11


	//遍历字符串
	//ASCII 字符串遍历直接使用下标
	theme := "狙击 start"
	for i := 0; i < len(theme); i++{
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}

	//Unicode 字符串遍历用 for range
	for _, s := range theme {
	    fmt.Printf("Unicode: %c  %d\n", s, s)
	}

	//字符串截取
	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ", ") //12
	pos := strings.Index(tracer[comma:], "死神")//3
	fmt.Println(comma, pos, tracer[comma+pos:])//12 3 死神bye bye

	//修改字符串
	//Go 语言的字符串无法直接修改每一个字符元素，只能通过重新构造新的字符串并赋值给原来的字符串变量实现
	angel := "Heros never die"
	angleBytes := []byte(angel)//修改字符串时，可以将字符串转换为 []byte 进行修改
	for i := 5; i <= 10; i++ {
	    angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))//[]byte 和 string 可以通过强制类型转换互转
	//线程安全，无须加锁，内存共享

	//字符串拼接
	// + 或 bytes.Buffer 类似 StringBuilder
	hammer := "吃我一锤"
	sickle := "死吧"
	// 声明字节缓冲
	var stringBuilder bytes.Buffer
	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())


	//格式化输出 fmt.Sprintf(格式化样式, 参数列表…)

	//Base64
	//Base64 编码是常见的对 8 比特字节码的编码方式之一。Base64 可以使用 64 个可打印字符来表示二进制数据，电子邮件就是使用这种编码。
	// 需要处理的字符串
    message := "Away from keyboard. https://golang.org/"
    // 编码消息
    encodedMessage := base64.StdEncoding.EncodeToString([]byte (message))
    // 输出编码完成的消息
    fmt.Println(encodedMessage)
    // 解码消息
    data, err := base64.StdEncoding.DecodeString(encodedMessage)
    // 出错处理
    if err != nil {
        fmt.Println(err)
    } else {
        // 打印解码完成的数据
        fmt.Println(string(data))
    }


}
```

###字符类型 byte 和 rune
+ unit8，实际上是byte 代表 ASCII 码的一个字符
+ rune 实际上是int32，代表 UTF-8 字符，处理中文、日文或者其他复和字符 
```
package main 
import (
	"fmt"
	"unicode"
)

func main(){
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	//函数
	//判断是否为字母：
	var ch4 = '\u0041'
	fmt.Println(unicode.IsLetter(ch4))
	//判断是否为数字：
	fmt.Println(unicode.IsDigit(ch4))
	//判断是否为空白符号：
	fmt.Println(unicode.IsSpace(ch4))
}
```
###数据类型转换

###指针
```
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
```

###变量逃逸分析



###关键字和标识符
#### 25 个关键字或保留字
break	default 	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var

#### 36 个预定义标识符
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr


###字符串和数字之间的相互转换
+ 整数转换成字符串，两种方法，一是fmt.Sprintf() 一个是strconv.Itoa()
+ FormatInt 和 FormatUint 可以按不同的进位制格式化数字
+ fmt.Printf 里的谓词 %b、%d、%o 和 %x 往往比 Format 函数方便
+ strconv 包内的 Atoi 函数或 ParseInt 函数用于解释表示整数的字符串，ParseUint 用于无符号整数
```
```
### 语言组合和方法集
结构体类型 struct
+ struct 可以嵌入任意其他类型的字段
+ struct 可以嵌套自身的指针类型的字段

##语言容器
###语言数组
