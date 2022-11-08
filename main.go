package main

import (
	"fmt"
	"runtime"
)

func main() {
	getCPU()
	getThreads()
}

func getCPU() {
	fmt.Println("Total Number of CPUs = ", runtime.NumCPU())
}

func getThreads() {
	fmt.Println("Total Number of threads on CPU")
}
