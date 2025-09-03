package main

import (
	"fmt"
	"unsafe"
)

// 示例1：不同字段顺序对结构体大小的影响
type Struct1 struct {
	a bool  // 1字节，对齐值1
	b int64 // 8字节，对齐值8
	c int8  // 1字节，对齐值1
}

type Struct2 struct {
	a bool  // 1字节
	c int8  // 1字节
	b int64 // 8字节
}

func main() {
	// 打印结构体大小
	fmt.Println("Struct1 size:", unsafe.Sizeof(Struct1{})) // 16字节
	fmt.Println("Struct2 size:", unsafe.Sizeof(Struct2{})) // 16字节

	// 打印字段偏移量（验证对齐）
	s1 := Struct1{}
	fmt.Printf("Struct1.a offset: %d\n", unsafe.Offsetof(s1.a)) // 0
	fmt.Printf("Struct1.b offset: %d\n", unsafe.Offsetof(s1.b)) // 8（跳过7字节填充）
	fmt.Printf("Struct1.c offset: %d\n", unsafe.Offsetof(s1.c)) // 16（超出结构体大小，实际在16位置）
}
