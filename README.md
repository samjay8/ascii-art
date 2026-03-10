# ascii-art

## Description
This is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. 

## How it works!
This project gets its acii character from some banner files with a specific graphical template representation. The files are formatted in a way that is not necessary to change them. 

### Banner Format
+ Each character has a height of 8 lines.
+ Characters are separated by a new line `\n`.

### How it match!
So, the ascii value for space is 32, then the calculation is each rune character converted to integer - 32 * 9.

## Installation Requirement
```
Go version needed
git clone <repository url>
cd repo
The banner files needed (standard.txt, shadow.txt, thinkertoy.txt)
```

## Usage
```
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "{Hello There}" | cat -e
go run . "Hello\n\nThere" | cat -e
```

## Supported Input
The project should handle an input with numbers, letters, spaces, special characters and `\n`

## Author
**Samuel Ojetunde**

---
