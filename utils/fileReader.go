package utils

import (
	"bufio"
	"fmt"
	"os"
)

func FileReader(filename string) ([]string, error) {
	var lines []string
	f, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
