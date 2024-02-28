package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fileName := "log.txt"
	lineNumber := 1

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Для выхода напишите 'Ex': ")
		input, _ := reader.ReadString('\n')

		if input == "Ex\n" {
			fmt.Println("Завершение...")
			break
		}

		entry := fmt.Sprintf("%d %s %s", lineNumber, time.Now().Format("2006-01-02 15:04:05"), input)
		if err := ioutil.WriteFile(fileName, []byte(entry), os.ModeAppend); err != nil {
			fmt.Println("Не возможно прочитать файл:", err)
		}

		lineNumber++
	}
}
