package main

import (
	ascii "ascii/Convert"
	"bufio"
	"fmt"
	"os"
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
	var bannerlines []string

	for scanner.Scan() {
		bannerlines = append(bannerlines, scanner.Text())
	}

	result := ascii.AsciiArt(input, bannerlines)
	fmt.Print(result)
}
