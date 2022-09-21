package main

import (
	"fmt"
	"sync"
)

type in struct {
	val      int32
	oddChan  chan int32
	evenChan chan int32
}

var serverChan = make(chan in, 100)

func Server() {
	for k := range serverChan {
		if k.val%2 == 0 {
			k.oddChan <- k.val
		} else {
			k.evenChan <- k.val
		}
	}
}
func main() {
	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	oddChan := make(chan int32)
	evenChan := make(chan int32)
	for idx := 0; idx < len(arr); idx++ {
		i := idx
		serverChan <- in{arr[i], oddChan, evenChan}
	}

	odds, evens := []int32{}, []int32{}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	go func() {
		for newOdd := range oddChan {
			odds = append(odds, newOdd)
			wg.Done()
		}
	}()
	go func() {
		for newEven := range evenChan {
			evens = append(evens, newEven)
			wg.Done()
		}
	}()
	go Server()
	wg.Wait()

	for _, resultItem := range odds {
		fmt.Printf("%d", resultItem)
		fmt.Printf("\n")
	}

	for i, resultItem := range evens {
		fmt.Printf("%d", resultItem)

		if i != len(evens)-1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
}
