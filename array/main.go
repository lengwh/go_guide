package main

import (
	"fmt"
	"unsafe"
)

/*
数组：一组数据的结合
特点：
长度：定长数组的长度是固定的，无法动态调整，变长数组长度不固定
数组的元素类型必须一致
数组是值类型，拷贝是值拷贝

数组的内存布局是连续的
切片：
切片是对数组的一个引用
切片是引用类型，拷贝是引用拷贝
切片的内存布局是不连续的
切片分为三部分：指针、长度、容量


*/

func main1() {
	var arr [10]int
	arrStr := [5]string{"a", "b", "c", "d", "e"}
	arrStr[0] = "hello"
	fmt.Println(arrStr)
	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, value := range arrStr {
		fmt.Println(index, value)
	}

	var matrix [3][4]int
	matrix[0][0] = 1
	fmt.Println(matrix)

	str := "hello"
	fmt.Printf("str size: %d \n", unsafe.Sizeof(str))
	fmt.Printf("arrStr size: %d bytes\n", unsafe.Sizeof(arrStr))
}

/*
*
切片创建方式
1. 通过数组创建切片
2. 通过make函数创建切片
3. 通过字面量创建切片
*/
func main01() {
	var s []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s1[1])
	s2 := make([]int, 10, 20)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	//拷贝
	s3 := make([]int, 5, 6)
	copy(s3, s)
	fmt.Println(len(s3), cap(s3))
	fmt.Println(s3)

	s3 = append(s3, 100, 200, 300)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))
	/**
	nil和空切片
	空切片：长度和容量都是0，但是不等于nil，指针不为nil
	nil切片：长度和容量都是0，指针为nil
	*/
	slice1 := make([]int, 0)
	fmt.Println(slice1)

	var slice2 []int
	fmt.Println(slice2)
	var slice3 = []int{}
	fmt.Println(slice3)
	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)
	fmt.Println(slice3 == nil)

}

/*
*
切片扩容机制
*/
func main3() {

}

/*
*
内存对齐
*/
func main() {
	var a uint8
	var b uint32
	var c uint16
	var d float64
	var e string
	var f bool
	var g int
	fmt.Println(unsafe.Sizeof(a)) // 1
	fmt.Println(unsafe.Sizeof(b)) // 4
	fmt.Println(unsafe.Sizeof(c)) // 2
	fmt.Println(unsafe.Sizeof(d)) // 8
	fmt.Println(unsafe.Sizeof(e)) // 16
	fmt.Println(unsafe.Sizeof(f)) // 1
	fmt.Println(unsafe.Sizeof(g)) // 8
}
