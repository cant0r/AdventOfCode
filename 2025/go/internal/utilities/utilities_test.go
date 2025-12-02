package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPrintResult(t *testing.T) {
	testNumber := 1337
	stdout, stdin, err := os.Pipe()

	if err != nil {
		t.Fatal("Couldn't get file descriptors from the operating system.")
	}

	defer stdout.Close()
	defer stdin.Close()

	os.Stdout = stdin

	PrintResult(testNumber)
	capturedOutput, err := bufio.NewReader(stdout).ReadString('\n')

	if err != nil {
		t.Fatal("Couldn't read from stdout???")
	}

	if !strings.Contains(capturedOutput, strconv.Itoa(testNumber)) {
		t.Errorf("'%s' didn't contain %d", capturedOutput, testNumber)
	}
}

func TestMod(t *testing.T) {
	divisor := 5
	numbers := []int{-2, -1, 0, 1, 2}
	expectedRemainders := []int{3, 4, 0, 1, 2}

	for index, number := range numbers {
		if remainder := Mod(number, divisor); remainder != expectedRemainders[index] {
			t.Errorf("%d%%%d is expected to be %d but got %d", number, divisor, expectedRemainders[index], remainder)
		}
	}
}
