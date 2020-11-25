package main

import (
	"context"
	"fmt"
	"time"
)

//func worker(ctx context.Context, ch chan int) {
//	fmt.Println("Starting worker")
//	time.Sleep(5 * time.Second)
//	ch <- 5
//	fmt.Println("Ending worker")
//
//}
func worker(ctx context.Context, upChannel chan int) {
	fmt.Println("Starting worker")
	downChannel := make(chan int)
	defer close(downChannel)

	go timer(ctx, downChannel)
	select {
	case result := <-downChannel:
		upChannel <- result
		fmt.Println("worker succeeded")

	case <-ctx.Done():
		fmt.Println("ERRRR ....... worker canceled by timeout")
	}

}

func timer(ctx context.Context, ch chan int) {
	fmt.Println("START: Slow operation")

	for i := 0; i < 100000; i++ {
		j := i + i*45
		if j%397 == 0 {
			fmt.Print("j")
		}
	}
	ch <- 100
	fmt.Println("END: Slow operation")
}
func main() {

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)

	messages := make(chan int)
	defer close(messages)
	fmt.Println("Main begun")

	go worker(ctx, messages)
	select {
	case msg := <-messages:
		fmt.Println("Main received ", msg)
	case <-ctx.Done():
		fmt.Println("Main Request Timed out.... context cancelled")
	}
	//time.Sleep(5 * time.Second)

}
