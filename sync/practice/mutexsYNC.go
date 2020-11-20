package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock = sync.RWMutex{}
)

func releaseThread() {
	lock.Lock()
	lock.Lock()
	for i := 1; i <= 300; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	lock.RUnlock()
}

func StartThreadsA() {
	for i := 1; i <= 2; i++ {
		go releaseThread()
	}
	time.Sleep(1 * time.Second)
}

func main() {
	//for i := 1; i <= 2; i++ {
	StartThreadsA()
	//RunAndWait()
	//}
	time.Sleep(1 * time.Millisecond)
}

func RunAndWait() {
	go callBackFunction()
	time.Sleep(10 * time.Second)
}

func callBackFunction() {
	lock.Lock()
	lock.Lock()
	fmt.Println("Hello its working")
}