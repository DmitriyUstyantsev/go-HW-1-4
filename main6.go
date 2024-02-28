package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)
	lineNumber := 1

	for {
		fmt.Print("Введите ваше сообщение (напишите 'exit' для выхода): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("%d %s %s\n", lineNumber, currentTime, input)
		_, err := file.WriteString(formattedMessage)

		if err != nil {
			fmt.Println(err)
			return
		}

		lineNumber++
	}

	fmt.Println("Запись прошла успешно.")
}
