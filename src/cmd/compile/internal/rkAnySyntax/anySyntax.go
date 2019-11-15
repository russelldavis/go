package rkAnySyntax

import (
	"cmd/compile/internal/rkSyntax"
	"cmd/compile/internal/syntax"
	"io"
	"path/filepath"
)

// Delegates to ParseGo or ParseRk based on the filename
func ParseAny(filename string, base *syntax.PosBase, src io.Reader, errh syntax.ErrorHandler, pragh syntax.PragmaHandler, mode syntax.Mode) (_ *syntax.File, first error) {
	ext := filepath.Ext(filename)
	if ext == ".rk" {
		return rkSyntax.ParseRk(base, src, errh, pragh, mode)
	} else {
		return syntax.ParseGo(base, src, errh, pragh, mode)
	}
}
