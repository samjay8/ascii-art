package ascii

import (
	"strings"
)

func AsciiArt(input string, bannerlines []string) string {
	if input == "" {
		return ""
	}
	if input == "\\n" {
		return "\n"
	}

	textsplit := strings.Split(input, `\n`)
	var result string

	for _, char := range textsplit {
		if char == "" {
			result += "\n"
			continue
		}

		for row := 0; row < 8; row++ {
			var rowString string
			for col := 0; col < len(char); col++ {
				position := int(char[col]-32) * 9
				rowString += bannerlines[position+row]
			}
			result += rowString + "\n"
		}
	}
	return result
}
