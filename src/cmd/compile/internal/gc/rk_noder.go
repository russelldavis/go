package gc

import "cmd/compile/internal/syntax"

type TypedNames struct {
	// Each array will have the same length.
	// It's structured this way, rather than an array of Name/Type pairs,
	// so we can set the Names array directly to Node.List on OAS2 nodes
	// (see variter).
	Names []*Node
	Types []*Node
}

func (p *noder) declTypedNames(typedNames []*syntax.TypedName) TypedNames {
	var res = TypedNames{}
	for _, typedName := range typedNames {
		res.Types = append(res.Types, p.typeExprOrNil(typedName.Type))
		res.Names = append(res.Names, p.declName(typedName.Name))
	}
	return res
}
