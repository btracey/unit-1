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

// Dimless represents a dimensionless constant
type Dimless float64

const (
	Yottaone Dimless = 1e24
	Zettaone Dimless = 1e21
	Exaone   Dimless = 1e18
	Petaone  Dimless = 1e15
	Teraone  Dimless = 1e12
	Gigaone  Dimless = 1e9
	Megaone  Dimless = 1e6
	Kiloone  Dimless = 1e3
	Hectoone Dimless = 1e2
	Decaone  Dimless = 1e1
	One      Dimless = 1.0
	Decione  Dimless = 1e-1
	Centione Dimless = 1e-2
	Millione Dimless = 1e-3
	Microone Dimless = 1e-6
	Nanoone  Dimless = 1e-9
	Picoone  Dimless = 1e-12
	Femtoone Dimless = 1e-15
	Attoone  Dimless = 1e-18
	Zeptoone Dimless = 1e-21
	Yoctoone Dimless = 1e-24
)

// Unit converts the Dimless to a *Unit
func (d Dimless) Unit() *Unit {
	return New(float64(d), Dimensions{})
}

// Dimless allows Dimless to implement a Dimlesser interface
func (d Dimless) Dimless() Dimless {
	return d
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (d *Dimless) From(u Uniter) error {
	if !DimensionsMatch(u, One) {
		*d = Dimless(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*d = Dimless(u.Unit().Value())
	return nil
}

func (d Dimless) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", d, float64(d))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(d))
		fmt.Fprint(fs, " ")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g )", c, d, float64(d))
		return
	}
}
