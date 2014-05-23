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

// Illuminance represents illuminance in lux
type Illuminance float64

const (
	Yottalux Illuminance = 1e24
	Zettalux Illuminance = 1e21
	Exalux   Illuminance = 1e18
	Petalux  Illuminance = 1e15
	Teralux  Illuminance = 1e12
	Gigalux  Illuminance = 1e9
	Megalux  Illuminance = 1e6
	Kilolux  Illuminance = 1e3
	Hectolux Illuminance = 1e2
	Decalux  Illuminance = 1e1
	Lux      Illuminance = 1.0
	Decilux  Illuminance = 1e-1
	Centilux Illuminance = 1e-2
	Millilux Illuminance = 1e-3
	Microlux Illuminance = 1e-6
	Nanolux  Illuminance = 1e-9
	Picolux  Illuminance = 1e-12
	Femtolux Illuminance = 1e-15
	Attolux  Illuminance = 1e-18
	Zeptolux Illuminance = 1e-21
	Yoctolux Illuminance = 1e-24
)

// Unit converts the Illuminance to a *Unit
func (il Illuminance) Unit() *Unit {
	return New(float64(il), Dimensions{
		LuminousIntensityDim: 1,
		LengthDim:            -2,
	})
}

// Illuminance allows Illuminance to implement a Illuminanceer interface
func (il Illuminance) Illuminance() Illuminance {
	return il
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (il *Illuminance) From(u Uniter) error {
	if !DimensionsMatch(u, Lux) {
		*il = Illuminance(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*il = Illuminance(u.Unit().Value())
	return nil
}

func (il Illuminance) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", il, float64(il))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(il))
		fmt.Fprint(fs, " lx")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g lx)", c, il, float64(il))
		return
	}
}