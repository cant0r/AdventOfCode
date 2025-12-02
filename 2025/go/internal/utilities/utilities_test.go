package utilities

import (
	"bufio"
	"maps"
	"os"
	"slices"
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
		t.Errorf("'%s' didn't contain %d\n", capturedOutput, testNumber)
	}
}

func TestMod(t *testing.T) {
	divisor := 5
	numbers := []int{-2, -1, 0, 1, 2}
	expectedRemainders := []int{3, 4, 0, 1, 2}

	for index, number := range numbers {
		if remainder := Mod(number, divisor); remainder != expectedRemainders[index] {
			t.Errorf("%d%%%d is expected to be %d but got %d\n", number, divisor, expectedRemainders[index], remainder)
		}
	}
}

func TestGetUniqueCharacters(t *testing.T) {
	words := []string{"apple", "cat", "belabela", "aaddaabb"}
	uniqueCharacters := []string{"aple", "cat", "bela", "adb"}
	uniqueCharactersCounts := []int{4, 3, 4, 3}

	for index, word := range words {
		actualUniqueCharacters, actualUniqueCharactersCount := GetUniqueCharacters(word)
		if uniqueCharactersCounts[index] != actualUniqueCharactersCount {
			t.Errorf("Unique character count is exptected to be %d but got %d for %s\n", uniqueCharactersCounts[index], actualUniqueCharactersCount, word)
		}

		expectedUniqueCharacters := []rune(uniqueCharacters[index])

		i := 0
		for actualUniqueCharacter := range maps.Keys(actualUniqueCharacters) {
			if !slices.Contains(expectedUniqueCharacters, actualUniqueCharacter) {
				t.Errorf("Couldn't find %q in the expected unique characters in %q\n", actualUniqueCharacter, expectedUniqueCharacters)
			}
			i++
		}
	}
}
