package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int //numbers
	var m int //target sum
	var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
	//var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
	c := make(chan []int)
	d := make(chan int)
	fmt.Fscanf(bufin, "%d %d\n", &n, &m)
	arr := make([]int, n)
	go ReadN(arr, n, bufin, c)
	go checker(<-c, n, m, d)
	fmt.Print(<-d)
}

func ReadN(arr []int, n int, bufin *bufio.Reader, c chan<- []int) {
	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &arr[i])
	}
	c <- arr
}

func checker(arr []int, n, m int, d chan<- int) {
	var start int
	var end int
	var sum int
	var result int

	for start < n {
		if sum > m || end == n {
			sum -= arr[start]
			start++
		} else {
			sum += arr[end]
			end++
		}
		if sum == m {
			result++
		}
	}
	d <- result
}
