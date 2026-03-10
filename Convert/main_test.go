package ascii

import (
	"strings"
	"testing"
)

func stubBanner() []string {
	lines := make([]string, 95*9)
	for c := 32; c < 127; c++ {
		base := (c - 32) * 9
		for row := 0; row < 8; row++ {
			lines[base+row] = string(rune(c))
		}
	}
	return lines
}

var b = stubBanner()

func rowCount(r string) int {
	return len(strings.Split(strings.TrimRight(r, "\n"), "\n"))
}

func TestSingleChar(t *testing.T) {
	if rowCount(AsciiArt("A", b)) != 8 {
		t.Errorf("want 8 rows, got %d", rowCount(AsciiArt("A", b)))
	}
}

func TestMultiChar(t *testing.T) {
	r := strings.Split(strings.TrimRight(AsciiArt("Hi", b), "\n"), "\n")
	if r[0] != "Hi" {
		t.Errorf("got %q, want \"Hi\"", r[0])
	}
}

func TestNewlineSplits(t *testing.T) {
	if rowCount(AsciiArt(`A\nB`, b)) != 16 {
		t.Errorf("want 16 rows, got %d", rowCount(AsciiArt(`A\nB`, b)))
	}
}

func TestEmptySegment(t *testing.T) {
	if AsciiArt(`\n`, b) != "\n\n" {
		t.Errorf("got %q, want \"\\n\\n\"", AsciiArt(`\n`, b))
	}
}

func TestEndsWithNewline(t *testing.T) {
	if !strings.HasSuffix(AsciiArt("Hello", b), "\n") {
		t.Error("result does not end with newline")
	}
}

func TestSpace(t *testing.T) {
	if rowCount(AsciiArt(" ", b)) != 8 {
		t.Errorf("want 8 rows, got %d", rowCount(AsciiArt(" ", b)))
	}
}
