package main

import (
	"fmt"
	"runtime"
)

func main() {
	get_CPU_Cores()
	getThreads()
}

func get_CPU_Cores() {
	fmt.Println("Total Number of CPU Cores =", runtime.NumCPU())
}

func getThreads() {
	fmt.Println("Total Number of threads on CPU")
}
