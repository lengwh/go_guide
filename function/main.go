package main

import "fmt"

/*
*
go 语言支持多返回值
具名返回值
匿名函数
*/
func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b int) (x, y int) {
	x, y = b, a
	return
}

func main01() {
	//匿名函数
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	// 闭包
	counter := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	count := counter()
	fmt.Println(count())
	fmt.Println(count())

	//可变参数
	sum := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20, 30))
}

// 函数作为参数
func apply(a, b int, f func(int, int) int) int {
	return f(a, b)
}

// 函数作为返回值
func makeAdd() func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main02() {
	add := func(a, b int) int {
		return a + b
	}
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(apply(1, 2, add))
	fmt.Println(apply(1, 2, sub))
	f := makeAdd()
	fmt.Println(f(10, 20))
}

//defer 语句会延迟一个函数的执行直到上层函数返回, 无论是正常返回还是异常返回
//defer 语句会将延迟的函数调用压入一个栈中, 当上层函数返回时,
//按照后进先出的顺序依次执行这些延迟的函数调用
//defer 语句常用于资源释放、文件关闭、解锁等操作

func deferFunc() {
	fmt.Println("deferFunc")
}

func returnFunc() int {
	defer deferFunc()
	fmt.Println("returnFunc")
	return 0
}

func testDeferfunc() (a int) {
	defer func() {
		a += 100
	}()
	return 1
}

// defer 遇见panic
func panicDefer() {
	defer func() { fmt.Println("deferFunc before") }()
	defer func() { fmt.Println("deferFunc before2") }()
	panic("panic")
	defer func() { fmt.Println("deferFunc after") }()
}

/*
*

	panic recoer

4. 注意事项
recover 必须在 defer 语句中使用才有效。
如果 defer 语句在 panic 之后才定义，则 recover 无法捕获该 panic。
panic 会终止当前函数的执行，但会先执行所有已定义的 defer 语句。
多个 panic 只有最后一个会被捕获。
*/
func main03() {
	//fmt.Println(testDeferfunc()) //101
	//panicDefer()

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		} else {
			fmt.Println("No panic")
		}
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func funcA(index, value int) int {
	fmt.Println(index)
	return index
}

func main04() {
	defer funcA(1, funcA(3, 0))
	defer funcA(2, funcA(4, 0))
}

// 递归
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main06() {
	fmt.Println(factorial(5))
}

func processData(data int, callback func(int)) {
	fmt.Println(data)
	callback(data)
}

func handleCallback(data int) {
	fmt.Println("handleCallback", data)
}

func main() {
	processData(10, handleCallback)
}
