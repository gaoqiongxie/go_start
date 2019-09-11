//map 无序 key-value
//var map_variable map[key_data_type]value_data_type
//map_variable := make(map[key_data_type]value_data_type)
package main

import (
    "fmt"
    "sync"
    "bufio"
    "os"
)

func main() {
    // test1()

    // deleteTest()

    // mapErrorTest()

    // sysnMapTest()

    dup1Test()
}

//
func test1(){
	var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap = make(map[string]string)

    /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "巴黎"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"

    /*使用键输出地图值 */ 
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    /*查看元素在集合中是否存在 */
    capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("American 的首都是", capital)
    } else {
        fmt.Println("American 的首都不存在")
    }
}

func deleteTest(){
	/* 创建map */
    countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

    fmt.Println("原始地图")

    /* 打印地图 */
    for country := range countryCapitalMap {
            fmt.Println(country, "首都是", countryCapitalMap [ country ])
    }

    /*删除元素*/ 
    delete(countryCapitalMap, "France")
    fmt.Println("法国条目被删除")

    fmt.Println("删除元素后地图")

    /*打印地图*/
    for country := range countryCapitalMap {
            fmt.Println(country, "首都是", countryCapitalMap [ country ])
    }

    //清空 map 的所有元素，唯一办法就是重新 make 一个新的mao
}

//map 在并发情况下，只读是线程安全的，同时读写线程线程不安全
//以下会报错：fatal error: concurrent map read and map write
func mapErrorTest(){
    // 创建一个int到int的映射
    m := make(map[int]int)

    // 开启一段并发代码
    go func(){
        //不停地对map进行写入
        for{
            m[1] = 1
        }
    }()

    // 开启一段并发代码
    go func(){
        //不停地读取
        for{
            _ = m[1]
        }
    }()

    // 无限循环，让并发程序在后台执行
    for{

    }


}

//并发安全的 sync.Map
func sysnMapTest() {
    var scene sync.Map 
    // 将键值对保存到sync.Map
    scene.Store("greece", 97)
    scene.Store("london", 100)
    scene.Store("egypt", 200)
    // 从sync.Map中根据键取值
    fmt.Println(scene.Load("london"))
    // 根据键删除对应的键值对
    scene.Delete("london")
    // 遍历所有sync.Map中的键值对
    scene.Range(func(k, v interface{}) bool {
        fmt.Println("iterate:", k, v)
        return true
    })
}

//找出重复行
func dup1Test(){
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    //注意：忽略 input.Err() 中可能的错误
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }

}