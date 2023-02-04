package mylinter

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "mylinter is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "mylinter",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		// case *ast.Ident:
		// 	if n.Name == "gopher" {
		// 		pass.Reportf(n.Pos(), "identifier is gopher")
		// 	}
		case *ast.ForStmt:
			fmt.Println("found for statement")
			for _, s := range n.Body.List {
				switch s := s.(type) {
				case *ast.DeferStmt:
					pass.Reportf(s.Pos(), "defer should not be used in for loop.")
				}
			}
		}
	})

	return nil, nil
}
