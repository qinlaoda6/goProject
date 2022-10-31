package demo02

import (
	"fmt"
	"time"
)

func PrintfHelloWorld() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d Hello World !\n", i)
		time.Sleep(time.Millisecond * 200)
	}
}
