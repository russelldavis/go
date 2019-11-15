// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rkSyntax

import (
	"cmd/compile/internal/syntax"
	"io"
)

// Parse parses a single RK source file from src and returns the corresponding
// syntax tree. If there are errors, Parse will return the first error found,
// and a possibly partially constructed syntax tree, or nil.
//
// If errh != nil, it is called with each error encountered, and Parse will
// process as much source as possible. In this case, the returned syntax tree
// is only nil if no correct package clause was found.
// If errh is nil, Parse will terminate immediately upon encountering the first
// error, and the returned syntax tree is nil.
//
// If pragh != nil, it is called with each pragma encountered.
//
func ParseRk(base *PosBase, src io.Reader, errh syntax.ErrorHandler, pragh syntax.PragmaHandler, mode syntax.Mode) (_ *syntax.File, first error) {
	defer func() {
		if p := recover(); p != nil {
			if err, ok := p.(syntax.Error); ok {
				first = err
				return
			}
			panic(p)
		}
	}()

	var p parser
	p.init(base, src, errh, pragh, mode)
	p.next()
	return p.fileOrNil(), p.first
}
