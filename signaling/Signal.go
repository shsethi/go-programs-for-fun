package main


import (
	"fmt"
	"time"
)

//Function that signals after publishing the message
func PublishAndSignal(msg string) ( chan struct{}) {
	ch  := make(chan struct{})
	go func() {
		fmt.Println("Publishing message..... ")
		time.Sleep(5 * time.Second)
		fmt.Println(msg)
		close(ch)
	}()
	return ch
}

// Generator returns a channel that produces the numbers 1, 2, 3,…
// To stop the underlying goroutine, close the channel.
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}
func main() {
	wait := PublishAndSignal("Hello")
	<- wait
	fmt.Println("Published, now time to exit.")

	// Channel closing can also be used to kill another go routine
	fmt.Println("Generating numbers from infinite loop in a  go routine")
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println("Closing channel kills go routine")
	close(number)
}

