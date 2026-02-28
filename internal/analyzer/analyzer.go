package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "linter",
	Doc:  "this analyzer reports linting errors",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			if !isSupportedLogger(pass, sel) {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			lit, ok := call.Args[0].(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				return true
			}

			msg, _ := strconv.Unquote(lit.Value)

			checkRules(pass, lit, msg)

			return true
		})
	}

	return nil, nil
}

func isSupportedLogger(pass *analysis.Pass, sel *ast.SelectorExpr) bool {
	if ident, ok := sel.X.(*ast.Ident); ok {
		if obj := pass.TypesInfo.Uses[ident]; obj != nil {
			if pkgName, ok := obj.(*types.PkgName); ok {
				path := pkgName.Imported().Path()
				if path == "log/slog" {
					return true
				}
			}
		}
	}

	t := pass.TypesInfo.TypeOf(sel.X)
	if t == nil {
		return false
	}

	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	if named, ok := t.(*types.Named); ok {
		if named.Obj() != nil && named.Obj().Pkg() != nil {
			path := named.Obj().Pkg().Path()
			if path == "github.com/go-uber/zap" {
				return true
			}
		}
	}

	return false
}
