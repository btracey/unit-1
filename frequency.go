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

// Frequency represents a frequency in hertz
type Frequency float64

const (
	Yottahertz Frequency = 1e24
	Zettahertz Frequency = 1e21
	Exahertz   Frequency = 1e18
	Petahertz  Frequency = 1e15
	Terahertz  Frequency = 1e12
	Gigahertz  Frequency = 1e9
	Megahertz  Frequency = 1e6
	Kilohertz  Frequency = 1e3
	Hectohertz Frequency = 1e2
	Decahertz  Frequency = 1e1
	Hertz      Frequency = 1.0
	Decihertz  Frequency = 1e-1
	Centihertz Frequency = 1e-2
	Millihertz Frequency = 1e-3
	Microhertz Frequency = 1e-6
	Nanohertz  Frequency = 1e-9
	Picohertz  Frequency = 1e-12
	Femtohertz Frequency = 1e-15
	Attohertz  Frequency = 1e-18
	Zeptohertz Frequency = 1e-21
	Yoctohertz Frequency = 1e-24
)

// Unit converts the Frequency to a *Unit
func (f Frequency) Unit() *Unit {
	return New(float64(f), Dimensions{
		TimeDim: -1,
	})
}

// Frequency allows Frequency to implement a Frequencyer interface
func (f Frequency) Frequency() Frequency {
	return f
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (f *Frequency) From(u Uniter) error {
	if !DimensionsMatch(u, Hertz) {
		*f = Frequency(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*f = Frequency(u.Unit().Value())
	return nil
}

func (f Frequency) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", f, float64(f))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(f))
		fmt.Fprint(fs, " Hz")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g Hz)", c, f, float64(f))
		return
	}
}
