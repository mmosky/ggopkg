package constraint

import "golang.org/x/exp/constraints"

type Integer = constraints.Integer

type Float = constraints.Float

type Number interface {
	Integer | Float
}

type Ordered = constraints.Ordered

type Signed = constraints.Signed

type Unsigned = constraints.Unsigned
