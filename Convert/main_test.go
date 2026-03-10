package ascii

import (
	"bufio"
	"os"
	"testing"
)

func TestAscii(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	file, err := os.Open("../standard.txt")
	if err != nil {
		t.Fatal("Error opening file", err)
	}
	defer file.Close()

	var bannerlines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bannerlines = append(bannerlines, scanner.Text())
	}

	tests := []testCase{
		{"hello", "Hello", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"},
		{"newline", `Hello\nThere`, " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n _______   _                           \n|__   __| | |                          \n   | |    | |__     ___   _ __    ___  \n   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n   | |    | | | | |  __/ | |    |  __/ \n   |_|    |_| |_|  \\___| |_|     \\___| \n                                       \n                                       \n"},
		{"empty", "", "\n"},
		{"numbers", "123", "                    \n _   ____    _____  \n/ | |___ \\  |___ /  \n| |   __) |   |_ \\  \n| |  / __/   ___) | \n|_| |_____| |____/  \n                    \n                    \n"},
		{"spaces", "Hello World", " _    _          _   _          __          __                 _       _  \n| |  | |        | | | |               \\ \\        / /                | |     | | \n| |__| |   ___  | | | |   ___          \\ \\  /\\  / /    ___    _ __  | |   __| | \n|  __  |  / _ \\ | | | |  / _ \\          \\ \\/  \\/ /    / _ \\  | '__| | |  / _` | \n| |  | | |  __/ | | | | | (_) |          \\  /\\  /    | (_) | | |    | | | (_| | \n|_|  |_|  \\___| |_| |_|  \\___/            \\/  \\/      \\___/  |_|    |_|  \\__,_| \n                                                                                \n                                                                                \n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AsciiArt(tc.input, bannerlines)
			if got != tc.expected {
				t.Errorf("AsciiArt(%q) = %q, expected = %q", tc.input, got, tc.expected)
			}
		})
	}
}
