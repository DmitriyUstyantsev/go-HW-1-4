package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Твоя строка задана"
	frequencyMap := make(map[rune]int)

	// Посчитать  буквы
	for _, char := range sentence {
		if char != ' ' { // Исключить пробелы
			frequencyMap[char]++
		}
	}

	// Посчитать процент повторения
	totalChars := float64(len(strings.ReplaceAll(sentence, " ", "")))
	for key, value := range frequencyMap {
		percentage := (float64(value) / totalChars) * 100
		fmt.Printf("Letter %c occurs %d times, %.2f%%\n", key, value, percentage)
	}
}
