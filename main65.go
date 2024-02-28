package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var out []string
	var backtrack func(S string, left, right int)

	backtrack = func(S string, left, right int) {
		if len(S) == 2*n {
			out = append(out, S)
			return
		}
		if left < n {
			backtrack(S+"(", left+1, right)
		}
		if right < left {
			backtrack(S+")", left, right+1)
		}
	}
	backtrack("", 0, 0)
	return out
}

func main() {
	var pairs int
	fmt.Print("Введите количество пар скобок: ")
	_, err := fmt.Scanf("%d", &pairs)
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}
	result := generateParenthesis(pairs)
	fmt.Println(result)
}
