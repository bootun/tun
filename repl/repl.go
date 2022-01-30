package repl

import (
	"bufio"
	"fmt"
	"github.com/bootun/tun/lexer"
	"github.com/bootun/tun/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for i := l.NextToken(); i.Type != token.EOF; i = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", i)
		}
	}
}
