package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mgill25/monkey-go/lexer"
	"github.com/mgill25/monkey-go/token"
)

const PROMPT = "#> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		// Scan all the tokens in this line via the lexer
		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
			fmt.Printf("%+v\n", tok)
		}
	}
}
