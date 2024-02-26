///package main

//import (
//	"fmt"
//	"path/filepath"
//)

//func main() {
// Задайте произвольный путь к файлу
//	filePath := "/path/to/your/file/somefile.txt"

// Получите имя файла и расширение
//	fileName := filepath.Base(filePath)
//	fileExtension := filepath.Ext(fileName)

// Напечатать имя файла без расширения
//	fmt.Println(fileName[:len(fileName)-len(fileExtension)])

// Напечатать расширение файла
//	fmt.Println(fileExtension)
//}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as a command-line argument")
		return
	}

	filePath := os.Args[1]
	fileName := filepath.Base(filePath)
	fileExt := filepath.Ext(filePath)
	fileExt = strings.TrimPrefix(fileExt, ".")

	// Output the filename and extension
	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
