package main

import (
	"fmt"
	"os"
)

func main() {
	// Создать новый файл
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Запишите некоторые данные в файл (выдаст ошибку, так как мы пытаемся выполнить запись в файл, открытый в режиме только для чтения)
	_, err = file.WriteString("Это проверка")
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	} else {
		fmt.Println("Запись в файл успешна")
	}
	// Смена свойств на только чтение
	err = os.Chmod("test.txt", 0444) // 0444 атрибут только чтение
	fmt.Println(err)
	return
}
