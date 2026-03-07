package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error, check the file input")
		return
	}

	input := os.Args[1]
	file, err := os.Open("standard.txt")

	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	textsplit := strings.Split(input, `\n`)

	for _, char := range textsplit {
		if char == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ {
			var rowString string
			for col := 0; col < len(char); col++ {
				startline := int(char[col]-32) * 9
				rowString += lines[startline+row]
			}
			fmt.Println(rowString)
		}
	}

}
