package analyzer

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func checkRules(pass *analysis.Pass, lit *ast.BasicLit, msg string) {
	msg = strings.TrimSpace(msg)

	if !startsWithLower(msg) {
		pass.Reportf(lit.Pos(), "log message must start with lowercase letter")
	}

	if !isEnglish(msg) {
		pass.Reportf(lit.Pos(), "log message must be in English")
	}
}

func startsWithLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return unicode.IsLower(r)
		}
	}
	return true
}

func isEnglish(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}
