package producer_consumer2

import (
	"fmt"
	"sync"
	"time"
)

//channel

var wgP sync.WaitGroup //wait group for producers
var wgC sync.WaitGroup //wait group for consumers

var channelFoo chan food = make(chan food, 2)

//food
type food struct {
	value int
}

//producer

func producer(c chan<- food, count int) {
	defer wgP.Done()

	for i := 0; i < count; i++ {
		time.Sleep(200 * time.Millisecond)

		f := food{value: i}
		fmt.Println("produced ", f)
		c <- f
	}
	fmt.Println("Producer ending   ")

}

//consumer
func consumer(c <-chan food, name string) {
	defer wgC.Done()
	for f := range c {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(name, "  consumed   ", f)
	}

}

//Why closing channel works when consuming by doing range over channel


func main() {

	fmt.Println("Start")

	go producer(channelFoo, 10)
	wgP.Add(1)

	go consumer(channelFoo, "consumer 1")
	wgC.Add(1)

	go consumer(channelFoo, "consumer 2")
	wgC.Add(1)

	//close channel when Producers are done
	wgP.Wait()
	// close channel autmatically causes range over channel to terminate
	close(channelFoo)

	wgC.Wait()
	fmt.Println("End")

}
