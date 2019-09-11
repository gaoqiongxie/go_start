//go Slice 语言切片
//slice是对数组一个连续片段的引用
//动态数组
package main 
import "fmt"

func main(){
	// initSliceTest()

	// nilSliceTest()

	// cutOutSliceTest()

	// appendAndCopyTest()

	// delTest()

	rangeTest()
	multiTest()
}

//定义&初始化切片
func initSliceTest(){
	//make([]T, length, capacity) 
	var numbers = make([]int,3,5)

    printSlice(numbers) // len=3 cap=5 slice=[0, 0, 0]

}

func printSlice(x []int){
	// len() 长度 
	// cap() 可以测量切片最长可达到多少
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

//空切片
func  nilSliceTest() {
	//声明一个未指定大小的数组来定义切片
	var numbers []int

	printSlice(numbers)

	if(numbers == nil){
	  fmt.Println("切片是空的")
	}
}

//截取
func cutOutSliceTest(){
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}   
	printSlice(numbers) // len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers) //numbers == [0 1 2 3 4 5 6 7 8]

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers[:]) //numbers == [0 1 2 3 4 5 6 7 8]

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4]) // numbers[1:4] == [1 2 3]

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3]) // numbers[:3] == [0 1 2]

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:]) //numbers[4:] == [4 5 6 7 8]

	numbers1 := make([]int,0,5)
	printSlice(numbers1) // len=0 cap=5 slice=[]

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2) // len=2 cap=9 slice=[0 1]

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)//len=3 cap=9 slice=[2 3 4]

}

func appendAndCopyTest(){
	//声明int切片
	var numbers []int
	printSlice(numbers) //len=0 cap=0 slice=[]

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)//len=1 cap=1 slice=[0]

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)//len=2 cap=2 slice=[0 1]

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers) //len=5 cap=6 slice=[0 1 2 3 4]
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1) //len=5 cap=12 slice=[0 1 2 3 4]

}

func delTest(){
	var a = []int{1, 2, 3}
	fmt.Println(a)
	a = a[:len(a)-1] // 删除尾部1个元素	
	fmt.Println(a)
	a = a[:len(a)-2] // 删除尾部N个元素
	fmt.Println(a)

	//或
	// a = []int{1, 2, 3}
	// a = a[1:] // 删除开头1个元素	
	// a = a[N:] // 删除开头N个元素
	

	// //不移动指针，将后面的数据向开头移动
	// a = []int{1, 2, 3}
	// a = append(a[:0], a[1:]...) // 删除开头1个元素
	// a = append(a[:0], a[N:]...) // 删除开头N个元素
	

	// //或copy
	// a = []int{1, 2, 3}
	// a = a[:copy(a, a[1:])] // 删除开头1个元素
	
	// a = a[:copy(a, a[N:])] // 删除开头N个元素

	// //对于删除中间的元素，需要对剩余的元素进行一次整体挪动，
	// //同样可以用 append 或 copy 原地完成
	// a = []int{1, 2, 3, ...}
	// a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	// a = append(a[:i], a[i+N:]...) // 删除中间N个元素
	// a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	// a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置       
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}

//循环
func rangeTest(){
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
	    fmt.Printf("Index: %d Value: %d\n", index, value)
	    //显示值和地址
	    fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
    		value, &value, &slice[index])
	}

	for _, value := range slice {
    	fmt.Printf("Value: %d\n", value)
	}

	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice); index++ {
	    fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}

//循环
func multiTest(){
	// 创建一个整型切片的切片
	slice := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice[0] = append(slice[0], 20)
	fmt.Println(slice)

}