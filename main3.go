package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Name     string
	Datetime time.Time
	Tags     string
	Link     string
}

func main() {
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления URL-адресов в список")
	fmt.Println("Нажмите Esc, чтобы выйти из приложения")

	// Список для хранения URL-адресов
	var urlList []Item

OuterLoop:
	for {
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Добавление нового URL-адреса в список
			fmt.Println("Введите новую запись в формате <теги описания url>")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильные аргументы в формате тегов описания URL-адреса")
				continue OuterLoop
			}
			tags := strings.Join(args[2:], ",")
			newItem := Item{
				Name:     args[1],
				Datetime: time.Now(),
				Tags:     tags,
				Link:     args[0],
			}
			urlList = append(urlList, newItem)
			fmt.Println("URL-адрес успешно добавлен")

		case 'l':
			// Отобразить список добавленных URL-адресов
			fmt.Println("Количество добавленных URL-адресов:", len(urlList))
			for _, item := range urlList {
				fmt.Println("Name:", item.Name)
				fmt.Println("URL:", item.Link)
				fmt.Println("Tags:", item.Tags)
				fmt.Println("Date:", item.Datetime.Format("2006-01-02 15:04:05"))
				fmt.Println("-------------")
			}

		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление URL-адреса из списка
			fmt.Println("Введите название ссылки для удаления")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			nameToDelete := strings.TrimSpace(text)
			var newURLList []Item
			for _, item := range urlList {
				if item.Name != nameToDelete {
					newURLList = append(newURLList, item)
				}
			}
			urlList = newURLList
			fmt.Printf("Link with name %s has been deleted\n", nameToDelete)

		default:
			// Если нажата клавиша Esc, выход из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
