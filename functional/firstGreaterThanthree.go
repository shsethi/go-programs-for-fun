package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {

	fmt.Println("Result")
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := getEvenNumberGreaterThanThree(a)
	fmt.Printf("first even number in list and greater than 3 is  -> %d   \n", result)

	//var b = []int{1, 2}
	//result := getEvenNumberGreaterThanThree(b)
	//fmt.Printf("first even number in list and greater than 3 is  -> %d   \n", result)

	//result = functional(b)
	//fmt.Printf("first even number in list and greater than 3 is  -> %d   \n", result)

}

func getEvenNumberGreaterThanThree(input []int) int {
	var result int
	for i := range input {
		if i > 3 && i % 2 == 0 {
			result = i
			break
		}
	}
	return result
}

func functional(input []int) int {

	l := filter(isGreaterThan(3), input)
	l2 := filter(isEven, l)

	result := l2[0] // !!! index out of range [0] with length 0

	return result
}

func functional2(input []int) (int, error) {

	l := filter(isGreaterThan(3), input)
	l2 := filter(isEven, l)

	if len(l2) == 0 {
		return 0, errors.New("No valid result")
	}
	return l2[0], nil
}

//pure function
func isEven(i int) bool {
	return i%2 == 0
}

//function that returns function
func isGreaterThan(num int) func (int) bool {
	return func (i int) bool {
		return i > num
	}
}
//functor
func filter(predicate func(int) bool, input []int) (output []int) {
	for _, i := range input {
		if predicate(i) {
			output = append(output, i)
		}
	}
	return
}




