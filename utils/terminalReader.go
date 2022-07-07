package utils

import (
	"bufio"
	"fmt"
	"os"
)

func TerminalReader() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		text = text[:len(text)-1]
		if text == ".exit" {
			break
		}
		fmt.Println(text)
	}
	os.Exit(0)
}
