package main

import (
	"fmt"
	"generics/alg"
)

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

func main() {
	intSliceExec()
	fmt.Println("")
	stringSliceExec()
}
