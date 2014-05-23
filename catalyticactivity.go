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

// CatalyticActivity represents catalytic activity in katals
type CatalyticActivity float64

const (
	Yottakatal CatalyticActivity = 1e24
	Zettakatal CatalyticActivity = 1e21
	Exakatal   CatalyticActivity = 1e18
	Petakatal  CatalyticActivity = 1e15
	Terakatal  CatalyticActivity = 1e12
	Gigakatal  CatalyticActivity = 1e9
	Megakatal  CatalyticActivity = 1e6
	Kilokatal  CatalyticActivity = 1e3
	Hectokatal CatalyticActivity = 1e2
	Decakatal  CatalyticActivity = 1e1
	Katal      CatalyticActivity = 1.0
	Decikatal  CatalyticActivity = 1e-1
	Centikatal CatalyticActivity = 1e-2
	Millikatal CatalyticActivity = 1e-3
	Microkatal CatalyticActivity = 1e-6
	Nanokatal  CatalyticActivity = 1e-9
	Picokatal  CatalyticActivity = 1e-12
	Femtokatal CatalyticActivity = 1e-15
	Attokatal  CatalyticActivity = 1e-18
	Zeptokatal CatalyticActivity = 1e-21
	Yoctokatal CatalyticActivity = 1e-24
)

// Unit converts the CatalyticActivity to a *Unit
func (ca CatalyticActivity) Unit() *Unit {
	return New(float64(ca), Dimensions{
		SubstanceAmountDim: 1,
		TimeDim:            -1,
	})
}

// CatalyticActivity allows CatalyticActivity to implement a CatalyticActivityer interface
func (ca CatalyticActivity) CatalyticActivity() CatalyticActivity {
	return ca
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (ca *CatalyticActivity) From(u Uniter) error {
	if !DimensionsMatch(u, Katal) {
		*ca = CatalyticActivity(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*ca = CatalyticActivity(u.Unit().Value())
	return nil
}

func (ca CatalyticActivity) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", ca, float64(ca))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64(ca))
		fmt.Fprint(fs, " kat")
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g kat)", c, ca, float64(ca))
		return
	}
}
