package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	exit := make(chan bool)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Нажмите Ctrl+C для остановки программы

	go func() {
		for {
			select {
			case <-sig:
				fmt.Println("exit the program")
				os.Exit(0)
			case <-exit:
				return
			}
		}
	}()

	i := 1
	for {
		fmt.Printf("%d ", i*i)
		i++
	}
}
