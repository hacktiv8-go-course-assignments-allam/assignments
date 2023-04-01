package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	jawaban1()
	jawaban2()
}

func jawaban1() {

	fmt.Println()
	fmt.Println("jawaban 1 running")
	interface1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	interface2 := []interface{}{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println(interface1, i+1)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println(interface2, i+1)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
	defer fmt.Println("jawaban 1 finish\n")
}

func jawaban2() {

	fmt.Println("jawaban 2 running")
	interface1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	interface2 := []interface{}{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup
	wg.Add(2)

	var mutex sync.Mutex

	go func() {
		for i := 0; i < 4; i++ {
			mutex.Lock()
			fmt.Println(interface1, i+1)
			mutex.Unlock()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 4; i++ {
			mutex.Lock()
			fmt.Println(interface2, i+1)
			mutex.Unlock()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
	defer fmt.Println("jawaban 2 finish\n")
}