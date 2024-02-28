package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	var wg sync.WaitGroup

	// Программа для вычисления квадрата
	wg.Add(1)
	go func(in <-chan int, out chan<- int) {
		defer wg.Done()
		for x := range in {
			out <- x * x
		}
		close(out)
	}(in, out)

	// Программа для расчета продукта
	wg.Add(1)
	go func(out <-chan int) {
		defer wg.Done()
		for x := range out {
			fmt.Println("Решение:", x*2)
		}
	}(out)

	// Ввод данных от пользователя
	for {
		var input string
		fmt.Scan(&input)
		if input == "stop" {
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Не верный ввод. Введите число или напишите stop")
			continue
		}
		in <- num
	}

	close(in)
	wg.Wait()
}
