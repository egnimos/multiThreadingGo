package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money = 10
	lock = sync.Mutex{}
	lockCond = sync.NewCond(&lock)
)

func stingy() {
	for i := 1; i<=1000; i++ {
		lock.Lock()
		money += 10
		fmt.Println("Stingy added the money", money)
		lockCond.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("STINGY IS DONE!!!")
}

func spendy() {
	for i := 1; i<=1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			lockCond.Wait()
		}
		money -= 20
		fmt.Println("SPENDY spended the money", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("SPENDY IS DONE!!!")
}

func main() {
	go stingy()
	go spendy()

	time.Sleep(3000 * time.Millisecond)
	fmt.Printf("so the account balance is %d \n", money)
}
