package main

import (
	"fmt"
	"sync"
	"time"
)

//channel

var wg sync.WaitGroup

var channelFoo chan food = make(chan food, 2)

//food
type food struct {
	value int
}

//producer

func producer(c chan food, count int) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		time.Sleep(200 * time.Millisecond)

		f := food{value: i}
		fmt.Println("produced ", f)
		c <- f
	}
	fmt.Println("Producer ending   ")

}

//consumer
func consumer(c chan food, name string) {
	defer wg.Done()
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case f := <-c:
			fmt.Println(name, "  consumed   ", f)
		case <-time.After(2 * time.Second):
			fmt.Println("Consumer ending   ")
			return
		}
	}

}

func main() {

	fmt.Println("Start")

	go producer(channelFoo, 10)
	wg.Add(1)

	go consumer(channelFoo, "consumer 1")
	wg.Add(1)

	go consumer(channelFoo, "consumer 2")
	wg.Add(1)

	wg.Wait()
	close(channelFoo)
	fmt.Println("End")

}

