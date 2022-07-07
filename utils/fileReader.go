package utils

import (
	"bufio"
	"fmt"
	"os"
)

func fileReader(filename string) ([]string, error) {
	var lines []string
	f, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println(scanner.Text())
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
