package main

import (
	"fmt"

	"github.com/trainchou/generics/alg"
)

type student struct {
	id   string
	name string
}

func intSliceExec() {
	fmt.Printf("int slice start")
	slice := alg.NewSlice()
	slice.Add(1)
	fmt.Println("current int slice:", slice)
	slice.Add(2)
	fmt.Println("current int slice:", slice)
	slice.Add(2)
	fmt.Println("current int slice:", slice)
	slice.Add(3)
	fmt.Println("current int slice:", slice)
	slice.Remove(2)
	fmt.Println("current int slice:", slice)
	slice.Remove(2)
	fmt.Println("current int slice:", slice)
	slice.Remove(3)
	fmt.Println("current int slice:", slice)
	fmt.Println("int slice end")
}

func stringSliceExec() {
	fmt.Printf("string slice start")
	slice := alg.NewSlice()
	slice.Add("hello")
	fmt.Println("current string slice:", slice)
	slice.Add("golang")
	fmt.Println("current string slice:", slice)
	slice.Add("golang")
	fmt.Println("current string slice:", slice)
	slice.Add("generic")
	fmt.Println("current string slice:", slice)
	slice.Remove("golang")
	fmt.Println("current string slice:", slice)
	slice.Remove("golang")
	fmt.Println("current string slice:", slice)
	slice.Remove("generic")
	fmt.Println("current string slice:", slice)
	fmt.Println("string slice end")
}

func structSliceExec() {
	fmt.Printf("struct slice start")
	xiaoMing := student{"1001", "xiao ming"}
	xiaoLei := student{"1002", "xiao lei"}
	xiaoFang := student{"1003", "xiao fang"}
	slice := alg.NewSlice()
	slice.Add(xiaoMing)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoFang)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoFang)
	fmt.Println("current struct slice:", slice)
	fmt.Println("struct slice end")
}

func main() {
	intSliceExec()
	fmt.Println("")
	stringSliceExec()
	fmt.Println("")
	structSliceExec()
	fmt.Println("")
}
