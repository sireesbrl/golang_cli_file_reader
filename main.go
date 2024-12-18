package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename string

	fmt.Println("Welcome to the GoLang File Reader!")
	fmt.Print("Enter the name/path of the file you want to read: ")
	fmt.Scan(&filename)

	readFile(filename)

}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
