/*
меню из 2 пунктов

записать значение в файл
считать его из мап
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {

    // Открытие файла с флагами для чтения и записи; создание, если файл не существует
	file, err := openFileToWrite("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
	data := make(map[string]string)



	writeToFile(file, "key1: value1")
	writeToFile(file, "key2: value2")
	writeToFile(file, "key3: value3")

	readFromFile("example.txt", &data) // передаем имя, потому что открываем отдельно 
	fmt.Println(data)

	writeToFile(file, "key4: value4")


	readFromFile("example.txt", &data) // передаем имя, потому что открываем отдельно 
	fmt.Println(data)

	defer file.Close()
}

func openFileToWrite(fileName string) (*os.File, error){
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	//defer file.Close()
	return file, err
} 

func writeToFile(file* os.File, str string){
	_, err := file.WriteString(str + "\n")
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func readFromFile(fileName string, m* map[string]string) {
	// читаем файл и записываем содержимое в мап
    file, errOpen := os.Open(fileName)
	if errOpen != nil {
		fmt.Println("Error opening file for reading:", errOpen)
		return
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		
		strs := strings.Split(scanner.Text(), ": ") // парсим и в мап
		(*m)[strs[0]] = strs[1]
		
    }

    // Проверка на наличие ошибок во время чтения
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}