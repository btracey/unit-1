// This file is autogenerated by github.com/gonum/unit/autogen
// Changes should be made to the autogenerated template rather than this one

// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
)

// SubstanceAmount represents a substance amount in mol
type SubstanceAmount float64

const (
	Yottamol SubstanceAmount = 1e24
	Zettamol SubstanceAmount = 1e21
	Examol   SubstanceAmount = 1e18
	Petamol  SubstanceAmount = 1e15
	Teramol  SubstanceAmount = 1e12
	Gigamol  SubstanceAmount = 1e9
	Megamol  SubstanceAmount = 1e6
	Kilomol  SubstanceAmount = 1e3
	Hectomol SubstanceAmount = 1e2
	Decamol  SubstanceAmount = 1e1
	Mol      SubstanceAmount = 1.0
	Decimol  SubstanceAmount = 1e-1
	Centimol SubstanceAmount = 1e-2
	Millimol SubstanceAmount = 1e-3
	Micromol SubstanceAmount = 1e-6
	Nanomol  SubstanceAmount = 1e-9
	Picomol  SubstanceAmount = 1e-12
	Femtomol SubstanceAmount = 1e-15
	Attomol  SubstanceAmount = 1e-18
	Zeptomol SubstanceAmount = 1e-21
	Yoctomol SubstanceAmount = 1e-24
)

// Unit converts the SubstanceAmount to a *Unit
func (s SubstanceAmount) Unit() *Unit {
	return New(float64(s), Dimensions{
		SubstanceAmountDim: 1,
	})
}

// SubstanceAmount allows SubstanceAmount to implement a SubstanceAmounter interface
func (s SubstanceAmount) SubstanceAmount() SubstanceAmount {
	return s
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (s *SubstanceAmount) From(u Uniter) error {
	if !DimensionsMatch(u, Mol) {
		*s = SubstanceAmount(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*s = SubstanceAmount(u.Unit().Value())
	return nil
}

func (s SubstanceAmount) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", s, float64(s))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		if !pOk {
			p = -1
		}
		w, wOk := fs.Width()
		if !wOk {
			w = -1
		}
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(s))
		fmt.Fprint(fs, " mol")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g mol)", c, s, float64(s))
		return
	}
}
