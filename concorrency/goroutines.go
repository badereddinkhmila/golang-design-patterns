package concorrency

func searchMax(list []int) int {
	max := 0
	for _, v := range list {
		if v >= max {
			max = v
		}
	}
	return max
}

func searchMaxChan(list []int, c chan int) {
	max := 0
	for _, v := range list {
		if v >= max {
			max = v
		}
	}
	c <- max
}
