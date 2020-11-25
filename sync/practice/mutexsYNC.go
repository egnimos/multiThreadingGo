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
	//for i := 1; i <= 21; i++ {
	 count()
	//StartThreadsA()
	//RunAndWait()
	//}
	//time.Sleep(1 * time.Millisecond)

	//wait group test
	//wg := sync.WaitGroup{}
	//wg.Wait()
	//fmt.Println("done!!")
	//count()
}

//wait group
func count() {
	wg := sync.WaitGroup{}
	x := 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&x, &wg)
	}
	wg.Wait()
	fmt.Printf("%d\n", x)
}

func increment(x *int, wg *sync.WaitGroup) {
	for i :=0; i < 100; i++ {

		*x += 1

	}
	wg.Done()
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