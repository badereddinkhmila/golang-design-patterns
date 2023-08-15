package concorrency

import (
	"fmt"
	"sync"
)

var (
	mutex          sync.Mutex
	accountBalance int
)

func initBalance() {
	accountBalance = 1000
}

func Mutexes() {
	var wg sync.WaitGroup

	// n° of goroutines will be runned by the group
	wg.Add(2)

	initBalance()

	go withdraw(700, &wg)
	go deposit(165, &wg)

	fmt.Printf("The final balance %d", accountBalance)
}

func withdraw(value int, wg *sync.WaitGroup) {
	// Lock shared state
	mutex.Lock()
	accountBalance -= value
	// Release shared state
	mutex.Unlock()
	// goroutine done
	wg.Done()
}

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	accountBalance += value
	mutex.Unlock()
	wg.Done()
}
