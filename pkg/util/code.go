package util

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
)

func Code(l *slog.Logger, title, code, lexer string) {
	border := strings.Repeat("-", 80)
	l.Info(fmt.Sprintf("\n%s\n%s\n%s", border, title, border))
	if err := quick.Highlight(os.Stdout, code, lexer, "terminal", "monokai"); err != nil {
		l.Debug(err.Error())
		fmt.Println(code)
	}
}
