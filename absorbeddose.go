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

// AbsorbedDose represents an absorbed dose of ionizing radiation in Grays
type AbsorbedDose float64

const (
	Yottagray AbsorbedDose = 1e24
	Zettagray AbsorbedDose = 1e21
	Exagray   AbsorbedDose = 1e18
	Petagray  AbsorbedDose = 1e15
	Teragray  AbsorbedDose = 1e12
	Gigagray  AbsorbedDose = 1e9
	Megagray  AbsorbedDose = 1e6
	Kilogray  AbsorbedDose = 1e3
	Hectogray AbsorbedDose = 1e2
	Decagray  AbsorbedDose = 1e1
	Gray      AbsorbedDose = 1.0
	Decigray  AbsorbedDose = 1e-1
	Centigray AbsorbedDose = 1e-2
	Milligray AbsorbedDose = 1e-3
	Microgray AbsorbedDose = 1e-6
	Nanogray  AbsorbedDose = 1e-9
	Picogray  AbsorbedDose = 1e-12
	Femtogray AbsorbedDose = 1e-15
	Attogray  AbsorbedDose = 1e-18
	Zeptogray AbsorbedDose = 1e-21
	Yoctogray AbsorbedDose = 1e-24
)

// Unit converts the AbsorbedDose to a *Unit
func (g AbsorbedDose) Unit() *Unit {
	return New(float64(g), Dimensions{
		LengthDim: 2,
		TimeDim:   -2,
	})
}

// AbsorbedDose allows AbsorbedDose to implement a AbsorbedDoseer interface
func (g AbsorbedDose) AbsorbedDose() AbsorbedDose {
	return g
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (g *AbsorbedDose) From(u Uniter) error {
	if !DimensionsMatch(u, Gray) {
		*g = AbsorbedDose(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*g = AbsorbedDose(u.Unit().Value())
	return nil
}

func (g AbsorbedDose) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", g, float64(g))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(g))
		fmt.Fprint(fs, " Gy")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g Gy)", c, g, float64(g))
		return
	}
}
