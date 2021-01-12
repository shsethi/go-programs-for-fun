package main

import (
	"fmt"
	"sync"
)

type RR struct {
	ips   []string
	index int
}

//unsafe incrementing of index variable
func (r *RR) GetIPs() string {

	ip := r.ips[r.index]
	r.index = (r.index + 1) % len(r.ips)
	//fmt.Println(ip)
	return ip
}

/*
 go run -race loadbalncer.go
will find a data race becuase index is being incremented

*/
func main() {
	args := []string{"100.1", "100.2", "100.3", "100.4", "100.5"}

	lb := RR{ips: args}

	const n =500
	var res [n]string
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		index := i
		go func(wgInner *sync.WaitGroup) {
			res[index] = lb.GetIPs()
			wgInner.Done()
		}(&wg)
	}
	wg.Wait()

	m := make(map[string]int)
	for _, re := range res {
		m[re] = m[re] + 1

	}
	//fmt.Printf("total %d \n%v \n", len(res), res)
	fmt.Println("Count of each ip")
	for k, v := range m {
		fmt.Println(k, v)
	}

}
