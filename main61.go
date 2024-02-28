package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filePath := "log.txt"
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не существует")
		} else {
			log.Fatal(err)
		}
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileSize := fileInfo.Size()
	if fileSize == 0 {
		fmt.Println("Файл пустой")
		return
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("%d %s %s\n", lineNumber, currentTime, line)
		lineNumber++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
