package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "log.txt"

	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) || fileInfo.Size() == 0 {
		fmt.Println("Файл не существует либо пуст.")
		return
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println(string(data))
}
