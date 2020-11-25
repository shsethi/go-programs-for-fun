package main

import "fmt"

type Pie struct {
	cherry int
}

func getPies() []Pie {
	var input = []int{1, 2}

	var pies = make([]Pie, len(input))
	for i, val := range input {
		pies[i] = Pie{cherry: val}
	}
	return pies
}

func main() {
	pies := getPies()

	greatPie := getGreaterThanThree2(pies)
	fmt.Print(greatPie.cherry)
}

func getGreaterThanThree2(input []Pie) Pie {
	for _, pie := range input {
		i := pie.cherry
		if i > 3 && i%2 == 0 {
			return pie
		}
	}
	return Pie{}
}
