package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateNumbers(ch chan<- int) {
	for {
		num := rand.Intn(10) // Генеруємо випадкове число від 0 до 9
		ch <- num
		time.Sleep(time.Second) // Затримка для наочності
	}
}

func calculateAverage(chIn <-chan int, chOut chan<- float64) {
	var sum, count int
	for num := range chIn {
		sum += num
		count++
		average := float64(sum) / float64(count)
		chOut <- average
	}
}

func printAverage(ch <-chan float64) {
	for avg := range ch {
		fmt.Printf("Середнє значення: %.2f\n", avg)
	}
}

func main() {
	numCh := make(chan int)
	avgCh := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		generateNumbers(numCh)
	}()
	go func() {
		defer wg.Done()
		calculateAverage(numCh, avgCh)
	}()
	go func() {
		defer wg.Done()
		printAverage(avgCh)
	}()

	// Запобігаємо завершенню програми
	wg.Wait()
}
