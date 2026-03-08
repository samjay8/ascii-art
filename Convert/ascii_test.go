package ascii

import (
	"bufio"
	"os"
	"testing"
)

func TestAsciiSimple(t *testing.T) {
	input := "Hello"
	files, err := os.Open("standard.txt")

	if err != nil {
		t.Fatal("Error opening file", err)
	}

	defer files.Close()
	var bannerlines []string

	scanner := bufio.NewScanner(files)
	for scanner.Scan() {
		bannerlines = append(bannerlines, scanner.Text())
	}

	got := AsciiArt(input, bannerlines)

	expected := ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  `

	if got != expected {
		t.Errorf("AsciiArt(%q) = %q, expected = %q", input, got, expected)
	}

}
