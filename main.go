package main

import (
	"concurrency/md_1"
	"fmt"
)

func main() {
	//md_1.PrintRange(-10, 50, 5)
	md_1.MinMaxAvgGoroutines(-10000, 10000, 1000000)
	fmt.Println("*================================*")
	md_1.MinMaxAvgNoGoroutines(-10000, 10000, 1000000)
}
