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

// Tesla represents a magnetic field strength in tesla
type MagneticFieldStrength float64

const (
	Yottatesla MagneticFieldStrength = 1e24
	Zettatesla MagneticFieldStrength = 1e21
	Exatesla   MagneticFieldStrength = 1e18
	Petatesla  MagneticFieldStrength = 1e15
	Teratesla  MagneticFieldStrength = 1e12
	Gigatesla  MagneticFieldStrength = 1e9
	Megatesla  MagneticFieldStrength = 1e6
	Kilotesla  MagneticFieldStrength = 1e3
	Hectotesla MagneticFieldStrength = 1e2
	Decatesla  MagneticFieldStrength = 1e1
	Tesla      MagneticFieldStrength = 1.0
	Decitesla  MagneticFieldStrength = 1e-1
	Centitesla MagneticFieldStrength = 1e-2
	Millitesla MagneticFieldStrength = 1e-3
	Microtesla MagneticFieldStrength = 1e-6
	Nanotesla  MagneticFieldStrength = 1e-9
	Picotesla  MagneticFieldStrength = 1e-12
	Femtotesla MagneticFieldStrength = 1e-15
	Attotesla  MagneticFieldStrength = 1e-18
	Zeptotesla MagneticFieldStrength = 1e-21
	Yoctotesla MagneticFieldStrength = 1e-24
)

// Unit converts the MagneticFieldStrength to a *Unit
func (t MagneticFieldStrength) Unit() *Unit {
	return New(float64(t), Dimensions{
		CurrentDim: -1,
		MassDim:    1,
		TimeDim:    -2,
	})
}

// MagneticFieldStrength allows MagneticFieldStrength to implement a MagneticFieldStrengther interface
func (t MagneticFieldStrength) MagneticFieldStrength() MagneticFieldStrength {
	return t
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (t *MagneticFieldStrength) From(u Uniter) error {
	if !DimensionsMatch(u, Tesla) {
		*t = MagneticFieldStrength(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*t = MagneticFieldStrength(u.Unit().Value())
	return nil
}

func (t MagneticFieldStrength) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", t, float64(t))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(t))
		fmt.Fprint(fs, " T")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g T)", c, t, float64(t))
		return
	}
}
