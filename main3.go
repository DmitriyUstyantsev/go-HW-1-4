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

	fmt.Println("Program for adding URLs to the list")
	fmt.Println("Press Esc to exit the application")

	// List to store URLs
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
			// Adding a new URL to the storage list
			fmt.Println("Enter a new entry in the format <url description tags>")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Enter the correct arguments in the url description tags format")
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
			fmt.Println("URL added successfully")

		case 'l':
			// Display list of added URLs
			fmt.Println("Number of added URLs:", len(urlList))
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
			// Deleting a URL from the storage list
			fmt.Println("Enter the name of the link to delete")
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
			// If Esc is pressed, exit the application
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
