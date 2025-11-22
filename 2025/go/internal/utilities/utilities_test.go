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
