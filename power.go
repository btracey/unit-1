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

// Power represents a power in Watts
type Power float64

const (
	Yottawatt Power = 1e24
	Zettawatt Power = 1e21
	Exawatt   Power = 1e18
	Petawatt  Power = 1e15
	Terawatt  Power = 1e12
	Gigawatt  Power = 1e9
	Megawatt  Power = 1e6
	Kilowatt  Power = 1e3
	Hectowatt Power = 1e2
	Decawatt  Power = 1e1
	Watt      Power = 1.0
	Deciwatt  Power = 1e-1
	Centiwatt Power = 1e-2
	Milliwatt Power = 1e-3
	Microwatt Power = 1e-6
	Nanowatt  Power = 1e-9
	Picowatt  Power = 1e-12
	Femtowatt Power = 1e-15
	Attowatt  Power = 1e-18
	Zeptowatt Power = 1e-21
	Yoctowatt Power = 1e-24
)

// Unit converts the Power to a *Unit
func (p Power) Unit() *Unit {
	return New(float64(p), Dimensions{
		CurrentDim: 1,
	})
}

// Power allows Power to implement a Powerer interface
func (p Power) Power() Power {
	return p
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (p *Power) From(u Uniter) error {
	if !DimensionsMatch(u, Watt) {
		*p = Power(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*p = Power(u.Unit().Value())
	return nil
}

func (p Power) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", p, float64(p))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(p))
		fmt.Fprint(fs, " W")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g W)", c, p, float64(p))
		return
	}
}