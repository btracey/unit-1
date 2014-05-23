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

// SolidAngle represents a solid angle in steradians
type SolidAngle float64

const (
	Yottasteradian SolidAngle = 1e24
	Zettasteradian SolidAngle = 1e21
	Exasteradian   SolidAngle = 1e18
	Petasteradian  SolidAngle = 1e15
	Terasteradian  SolidAngle = 1e12
	Gigasteradian  SolidAngle = 1e9
	Megasteradian  SolidAngle = 1e6
	Kilosteradian  SolidAngle = 1e3
	Hectosteradian SolidAngle = 1e2
	Decasteradian  SolidAngle = 1e1
	Steradian      SolidAngle = 1.0
	Decisteradian  SolidAngle = 1e-1
	Centisteradian SolidAngle = 1e-2
	Millisteradian SolidAngle = 1e-3
	Microsteradian SolidAngle = 1e-6
	Nanosteradian  SolidAngle = 1e-9
	Picosteradian  SolidAngle = 1e-12
	Femtosteradian SolidAngle = 1e-15
	Attosteradian  SolidAngle = 1e-18
	Zeptosteradian SolidAngle = 1e-21
	Yoctosteradian SolidAngle = 1e-24
)

// Unit converts the SolidAngle to a *Unit
func (s SolidAngle) Unit() *Unit {
	return New(float64(s), Dimensions{})
}

// SolidAngle allows SolidAngle to implement a SolidAngleer interface
func (s SolidAngle) SolidAngle() SolidAngle {
	return s
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (s *SolidAngle) From(u Uniter) error {
	if !DimensionsMatch(u, Steradian) {
		*s = SolidAngle(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*s = SolidAngle(u.Unit().Value())
	return nil
}

func (s SolidAngle) Format(fs fmt.State, c rune) {
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
		fmt.Fprint(fs, " sr")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g sr)", c, s, float64(s))
		return
	}
}
