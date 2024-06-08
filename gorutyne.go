package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan<- int) {
	for {
		num := rand.Intn(10) // Генеруємо випадкове число від 0 до 99
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
	rand.Seed(time.Now().UnixNano())

	numCh := make(chan int)
	avgCh := make(chan float64)

	go generateNumbers(numCh)
	go calculateAverage(numCh, avgCh)
	go printAverage(avgCh)

	// Запобігаємо завершенню програми
	select {}
}
