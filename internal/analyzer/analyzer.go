package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"

	"golang.org/x/tools/go/analysis"

	"github.com/ladev74/linter/internal/analyzer/rules"
)

const (
	slogPath = "log/slog"
	zapPath  = "github.com/go-uber/zap"
)

type Config struct {
	Rules rules.Config `yaml:"rules"`
}

const (
	analyzerName = "linter"
	analyzerDoc  = "this analyzer reports linting errors"
)

func New(cfg *Config) *analysis.Analyzer {
	analyzer := &analysis.Analyzer{
		Name: analyzerName,
		Doc:  analyzerDoc,
		Run: func(pass *analysis.Pass) (any, error) {
			return run(pass, cfg)
		},
	}

	return analyzer
}

func run(pass *analysis.Pass, cfg *Config) (interface{}, error) {
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

			errMsg := rules.CheckRules(cfg.Rules, msg)
			if errMsg != "" {
				pass.Reportf(lit.Pos(), "%s", errMsg)
			}

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
				return isSupportedLoggerPath(path)
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
			return isSupportedLoggerPath(path)
		}

		if named.Obj() != nil && named.Obj().Name() == "Logger" {
			return true
		}
	}

	return false
}

func isSupportedLoggerPath(path string) bool {
	return path == slogPath || path == zapPath
}
