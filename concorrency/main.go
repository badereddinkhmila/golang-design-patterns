package concorrency

import "fmt"

func Main() {
	fmt.Println("<------------ Goroutines ----------->")
	list := []int{88, 78, 13, 45, 66, 45, 90, 33}
	c := make(chan int)

	max1 := searchMax(list)
	go searchMaxChan(list, c)
	max2 := <-c

	fmt.Printf("Max 1: %d et Max 2: %d", max1, max2)
	fmt.Println("<------------ Goroutines ----------->")
}
