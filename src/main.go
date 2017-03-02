package main

import (
	"runtime"
	"fmt"
	"syscall"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc/(1024*1024))
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)


}
