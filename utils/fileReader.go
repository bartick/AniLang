package utils

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/bartick/AniLang/token"
)

func FileReader(filename string) ([]string, error) {
	var lines []string
	f, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	// lexedToken := token.NewLexer(f)
	// for {
	// 	pos, tok, lit := lexedToken.Lex()
	// 	if tok == token.EOF {
	// 		break
	// 	}

	// 	fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.column, tok, lit)
	// }
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
