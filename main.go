package main

import (
	"fmt"
	"sync"
)

func Worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		fmt.Printf("worker %d processando %d\n", id, n)
		results <- n * n
	}
}

func RunPool(in []int, n int) []int {
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}
	go func() {
		for _, v := range in {
			jobs <- v
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var out []int
	for r := range results {
		out = append(out, r)
	}
	return out
}

func main() {
	out := RunPool([]int{2, 3, 1, 12, 77, 10, 131, 120}, 3)
	fmt.Println(out)
}
