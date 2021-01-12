package main

import (
	"fmt"
	"sync"
)

type RR2 struct {
	ips   []string
	index int
	mutex   sync.Mutex
}

// uses lock for incerementing index safely
func (r *RR2) GetIPs() string {
	r.mutex.Lock()
	ip := r.ips[r.index]
	r.index = (r.index + 1)  % len(r.ips)
	//fmt.Println(ip)  // print should be always before unlock for correct logging
	r.mutex.Unlock()
	return ip

}

func main() {
	args := []string{"100.1", "100.2", "100.3", "100.4", "100.5"}

	lb := RR2{ips: args}

	const n = 500
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

//func main() {
//	args := []string{"100.1", "100.2", "100.3", "100.4", "100.5"}
//
//	lb := RR2{ips: args}
//
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		reqId := i
//		go func(wgInner *sync.WaitGroup) {
//			fmt.Println(lb.GetIPs(reqId))
//			wgInner.Done()
//		}(&wg)
//	}
//	wg.Wait()
//
//}
