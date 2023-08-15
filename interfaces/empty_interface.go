package interfaces

import "fmt"

func Double(elm any) {
	switch elm := elm.(type) {
	case int:
		fmt.Println(elm * 2)
	case float64:
		fmt.Println(elm * 2)
	case string:
		fmt.Println(elm + elm)
	default:
		fmt.Println("Not supported!!")
	}
}
