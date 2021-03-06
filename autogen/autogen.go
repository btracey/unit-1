package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Unit struct {
	Name          string
	Receiver      string
	Offset        int    // From normal (for example, mass base unit is kg, not kg)
	PrintString   string // print string for the unit (kg for mass)
	ExtraConstant []Constant
	Suffix        string
	Singular      string
	TypeComment   string // Text to comment the type
	Dimensions    []Dimension
	ErForm        string //For Xxxer interface
}

type Dimension struct {
	Name  string
	Power int
}

const (
	CurrentName           string = "CurrentDim"
	LengthName            string = "LengthDim"
	LuminousIntensityName string = "LuminousIntensityDim"
	MassName              string = "MassDim"
	SubstanceAmountName   string = "SubstanceAmountDim"
	TemperatureName       string = "TemperatureDim"
	TimeName              string = "TimeDim"
	AngleName             string = "AngleDim"
)

type Constant struct {
	Name  string
	Value string
}

type Prefix struct {
	Name  string
	Power int
}

var Prefixes = []Prefix{
	{
		Name:  "Yotta",
		Power: 24,
	},
	{
		Name:  "Zetta",
		Power: 21,
	},
	{
		Name:  "Exa",
		Power: 18,
	},
	{
		Name:  "Peta",
		Power: 15,
	},
	{
		Name:  "Tera",
		Power: 12,
	},
	{
		Name:  "Giga",
		Power: 9,
	},
	{
		Name:  "Mega",
		Power: 6,
	},
	{
		Name:  "Kilo",
		Power: 3,
	},
	{
		Name:  "Hecto",
		Power: 2,
	},
	{
		Name:  "Deca",
		Power: 1,
	},
	{
		Name:  "",
		Power: 0,
	},
	{
		Name:  "Deci",
		Power: -1,
	},
	{
		Name:  "Centi",
		Power: -2,
	},
	{
		Name:  "Milli",
		Power: -3,
	},
	{
		Name:  "Micro",
		Power: -6,
	},
	{
		Name:  "Nano",
		Power: -9,
	},
	{
		Name:  "Pico",
		Power: -12,
	},
	{
		Name:  "Femto",
		Power: -15,
	},
	{
		Name:  "Atto",
		Power: -18,
	},
	{
		Name:  "Zepto",
		Power: -21,
	},
	{
		Name:  "Yocto",
		Power: -24,
	},
}

var Units = []Unit{
	{
		Name:        "AbsorbedDose",
		Receiver:    "g",
		PrintString: "Gy",
		Suffix:      "gray",
		Singular:    "Gray",
		TypeComment: "AbsorbedDose represents an absorbed dose of ionizing radiation in Grays",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
	},
	{
		Name:        "Acceleration",
		Receiver:    "a",
		PrintString: "m/s^2",
		TypeComment: "Acceleration represents an acceleration in meters per second squared",
		Dimensions: []Dimension{
			{
				Name:  LengthName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
	},
	{
		Name:        "Angle",
		Receiver:    "a",
		PrintString: "rad",
		Suffix:      "radian",
		Singular:    "Radian",
		TypeComment: "Angle represents an angle in radians",
		Dimensions:  []Dimension{},
		ExtraConstant: []Constant{
			{
				Name:  "Degree",
				Value: "180/math.Pi",
			},
		},
	},
	{
		Name:        "Capacitance",
		Receiver:    "ca",
		PrintString: "F",
		Suffix:      "farad",
		Singular:    "Farad",
		TypeComment: "Capacitance represents an electric capacitance in farads",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  CurrentName,
				Power: 2,
			},
			{
				Name:  LengthName,
				Power: -2,
			},
			{
				Name:  MassName,
				Power: -1,
			},
			{
				Name:  TimeName,
				Power: 4,
			},
		},
	},
	{
		Name:        "CatalyticActivity",
		Receiver:    "ca",
		PrintString: "kat",
		Suffix:      "katal",
		Singular:    "Katal",
		TypeComment: "CatalyticActivity represents catalytic activity in katals",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  SubstanceAmountName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -1,
			},
		},
	},
	{
		Name:        "Charge",
		Receiver:    "ch", // because c is used for the rune
		PrintString: "C",
		Suffix:      "coulomb", // TODO: Check spelling
		Singular:    "Coulomb",
		TypeComment: "Change represents an electric charge in Coulombs",
		Dimensions: []Dimension{ // N/m^2 = (kg-m/s^2)/(m^2) = kg/(m-s^2)
			{
				Name:  CurrentName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: 1,
			},
		},
	},
	{
		Name:        "Conductance",
		Receiver:    "co",
		PrintString: "S",
		Suffix:      "siemens",
		Singular:    "Siemens",
		TypeComment: "Conductance represents an electrical conductance in Siemens",
		Dimensions: []Dimension{ // TODO: Check dimension
			{
				Name:  CurrentName,
				Power: 2,
			},
			{
				Name:  MassName,
				Power: -1,
			},
			{
				Name:  LengthName,
				Power: -2,
			},
			{
				Name:  TimeName,
				Power: 3,
			},
		},
	},
	{
		Name:        "Current",
		Receiver:    "a",
		PrintString: "A",
		Suffix:      "ampere",
		Singular:    "Ampere",
		TypeComment: "Current represents an electric current in Amperes",
		Dimensions: []Dimension{ // TODO: Check dimension
			{
				Name:  CurrentName,
				Power: 1,
			},
		},
	},
	{
		Name:        "Density",
		Receiver:    "d",
		PrintString: "kg/m^3",
		TypeComment: "Density represents a density in kilograms per meters cubed",
		Dimensions: []Dimension{
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  LengthName,
				Power: -3,
			},
		},
	},
	{
		Name:        "Dimless",
		Receiver:    "d",
		TypeComment: "Dimless represents a dimensionless constant",
		Dimensions:  []Dimension{},
	},
	{
		Name:        "Energy",
		Receiver:    "e",
		PrintString: "J",
		Suffix:      "joule",
		Singular:    "Joule",
		TypeComment: "Energy represents an amount of energy in joules",
		Dimensions: []Dimension{
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
	},
	{
		Name:        "Force",
		Receiver:    "f",
		PrintString: "N",
		Suffix:      "newton",
		Singular:    "Newton",
		TypeComment: "Force represents a force in Newtons",
		Dimensions: []Dimension{
			{
				Name:  LengthName,
				Power: 1,
			},
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
		ExtraConstant: []Constant{
			{
				Name:  "StandardGravity",
				Value: "9.80665", // TODO: Check this
			},
		},
	},
	{
		Name:        "Frequency",
		Receiver:    "f",
		PrintString: "Hz",
		Suffix:      "hertz",
		Singular:    "Hertz",
		TypeComment: "Frequency represents a frequency in hertz",
		Dimensions: []Dimension{
			{
				Name:  TimeName,
				Power: -1,
			},
		},
	},
	{
		Name:        "Illuminance",
		Receiver:    "il",
		PrintString: "lx",
		Suffix:      "lux",
		Singular:    "Lux",
		TypeComment: "Illuminance represents illuminance in lux",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  LuminousIntensityName,
				Power: 1,
			},
			{
				Name:  LengthName,
				Power: -2,
			},
		},
	},
	{
		Name:        "Inductance",
		Receiver:    "in",
		PrintString: "H",
		Suffix:      "henry",
		Singular:    "Henry",
		TypeComment: "Inductance represents an electrical inductance in henrys",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  CurrentName,
				Power: -2,
			},
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
	},
	{
		Name:        "Mass",
		Receiver:    "m",
		Offset:      -3,
		PrintString: "kg",
		Suffix:      "gram",
		Singular:    "Gram",
		TypeComment: "Mass represents a mass in kilograms",
		Dimensions: []Dimension{
			{
				Name:  MassName,
				Power: 1,
			},
		},
	},
	{
		Name:        "Length",
		Receiver:    "l",
		PrintString: "m",
		Suffix:      "meter",
		Singular:    "Meter",
		TypeComment: "Length represents a length in meters",
		Dimensions: []Dimension{
			{
				Name:  LengthName,
				Power: 1,
			},
		},
	},
	{
		Name:        "LuminousFlux",
		Receiver:    "l",
		PrintString: "Ω",
		Suffix:      "lumen",
		Singular:    "Lumen",
		TypeComment: "LuminousFlux represents a luminous flux in lumens",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  LuminousIntensityName,
				Power: 1,
			},
		},
	},
	{
		Name:        "LuminousIntensity",
		Receiver:    "l",
		PrintString: "cd",
		Suffix:      "candela",
		Singular:    "Candela",
		TypeComment: "LuminousIntensity represents a luminous intensity in candela",
		Dimensions: []Dimension{
			{
				Name:  LuminousIntensityName,
				Power: 1,
			},
		},
	},
	{
		Name:        "MagneticFlux",
		Receiver:    "m",
		PrintString: "Wb",
		Suffix:      "weber",
		Singular:    "Weber",
		TypeComment: "MagneticFlux represents a magnetic flux in webers",
		Dimensions: []Dimension{ // TODO: Check dimension
			{
				Name:  CurrentName,
				Power: -1,
			},
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
	},
	{
		Name:        "Power",
		Receiver:    "p",
		PrintString: "W",
		Suffix:      "watt",
		Singular:    "Watt",
		TypeComment: "Power represents a power in Watts",
		Dimensions: []Dimension{ // N/m^2 = (kg-m/s^2)/(m^2) = kg/(m-s^2)
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  TimeName,
				Power: -3,
			},
		},
	},
	{
		Name:        "Pressure",
		Receiver:    "p",
		PrintString: "Pa",
		Suffix:      "pascal",
		Singular:    "Pascal",
		TypeComment: "Pressure represents a pressure in Pascals",
		Dimensions: []Dimension{ // N/m^2 = (kg-m/s^2)/(m^2) = kg/(m-s^2)
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  LengthName,
				Power: -1,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
		ExtraConstant: []Constant{
			{
				Name:  "Bar",
				Value: "1e5",
			},
			{
				Name:  "Atmosphere",
				Value: "1.01e5", // TODO: Check this value
			},
		},
	},
	{
		Name:        "Resistance",
		Receiver:    "r",
		PrintString: "Ω",
		Suffix:      "ohm",
		Singular:    "Ohm",
		TypeComment: "Resistance represents a resistance is ohms",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  CurrentName,
				Power: -2,
			},
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -3,
			},
		},
	},
	{
		Name:        "Radioactivity",
		Receiver:    "r",
		PrintString: "Bq",
		Suffix:      "becquerel",
		Singular:    "Becquerel",
		TypeComment: "Radioactivity represents decays per unit time in Becquerels",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  TimeName,
				Power: -1,
			},
		},
	},
	{
		Name:        "EquivalentDose",
		Receiver:    "e",
		PrintString: "Sv",
		Suffix:      "sievert",
		Singular:    "Sievert",
		TypeComment: "EquivalentDose represents an equivalent dose of ionizing radiaton in sieverts",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
	},
	{
		Name:        "SolidAngle",
		Receiver:    "s",
		PrintString: "sr",
		Suffix:      "steradian",
		Singular:    "Steradian",
		TypeComment: "SolidAngle represents a solid angle in steradians",
		Dimensions:  []Dimension{},
	},
	{
		Name:        "SubstanceAmount",
		Receiver:    "s",
		PrintString: "mol",
		Suffix:      "mol",
		Singular:    "Mol",
		TypeComment: "SubstanceAmount represents a substance amount in mol",
		Dimensions: []Dimension{
			{
				Name:  SubstanceAmountName,
				Power: 1,
			},
		},
	},
	{
		Name:        "Temperature",
		Receiver:    "t",
		PrintString: "K",
		Suffix:      "kelvin",
		Singular:    "Kelvin",
		TypeComment: "Temperature represents a temperature in Kelvin",
		Dimensions: []Dimension{ // N/m^2 = (kg-m/s^2)/(m^2) = kg/(m-s^2)
			{
				Name:  TemperatureName,
				Power: 1,
			},
		},
	},
	{
		Name:        "MagneticFieldStrength",
		Receiver:    "t",
		PrintString: "T",
		Suffix:      "tesla",
		Singular:    "Tesla",
		TypeComment: "Tesla represents a magnetic field strength in tesla",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  CurrentName,
				Power: -1,
			},
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -2,
			},
		},
	},
	{
		Name:        "Time",
		Receiver:    "t",
		PrintString: "s",
		Suffix:      "second",
		Singular:    "Second",
		TypeComment: "Time represents a time in seconds",
		ExtraConstant: []Constant{
			{
				Name:  "Hour",
				Value: "3600",
			},
			{
				Name:  "Minute",
				Value: "60",
			},
		},
		Dimensions: []Dimension{
			{
				Name:  TimeName,
				Power: 1,
			},
		},
		ErForm: "Timer",
	},
	{
		Name:        "Velocity",
		Receiver:    "v",
		PrintString: "m/s",
		TypeComment: "Velocity represents a velocity in meters per second",
		Dimensions: []Dimension{
			{
				Name:  LengthName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -1,
			},
		},
		// Speed of light?
	},
	{
		Name:        "Voltage",
		Receiver:    "e",
		PrintString: "V",
		Suffix:      "volt",
		Singular:    "Volt",
		TypeComment: "ElecPoten represents an electric potential in Volts",
		Dimensions: []Dimension{ // TODO: Check Dimension
			{
				Name:  CurrentName,
				Power: -1,
			},
			{
				Name:  LengthName,
				Power: 2,
			},
			{
				Name:  MassName,
				Power: 1,
			},
			{
				Name:  TimeName,
				Power: -3,
			},
		},
	},
}

var gopath string
var unitPkgPath string

func init() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("no gopath")
	}

	unitPkgPath = filepath.Join(gopath, "src", "github.com", "gonum", "unit")
}

// Generate generates a file for each of the units
func main() {
	for _, unit := range Units {
		generate(unit)
	}
}

const headerTemplate = `// This file is autogenerated by github.com/gonum/unit/autogen
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

// {{.TypeComment}}
type {{.Name}} float64
`

var header = template.Must(template.New("header").Parse(headerTemplate))

const constTemplate = `
const(
	{{$unit := .Unit}}
	{{range $unit.ExtraConstant}} {{.Name}} {{$unit.Name}} = {{.Value}}
	{{end}}
	{{$prefixes := .Prefixes}}
	{{range $prefixes}} {{if .Name}} {{.Name}}{{$unit.Suffix}} {{else}} {{$unit.Singular}} {{end}} {{$unit.Name}} = {{if .Power}} 1e{{.Power}} {{else}} 1.0 {{end}}
	{{end}}
)
`

var prefix = template.Must(template.New("prefix").Parse(constTemplate))

const methodTemplate = `
// Unit converts the {{.Name}} to a *Unit
func ({{.Receiver}} {{.Name}}) Unit() *Unit{
	return New(float64({{.Receiver}}), Dimensions{
		{{range .Dimensions}} {{.Name}}: {{.Power}},
		{{end}}
		})
}

// {{.Name}} allows {{.Name}} to implement a {{if .ErForm}}{{.ErForm}}{{else}}{{.Name}}er{{end}} interface
func ({{.Receiver}} {{.Name}}) {{.Name}}() {{.Name}} {
	return {{.Receiver}}
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func ({{.Receiver}} *{{.Name}}) From(u Uniter) error{
	if !DimensionsMatch(u, {{.Singular}}){
		*{{.Receiver}} = {{.Name}}(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*{{.Receiver}} = {{.Name}}(u.Unit().Value())
	return nil
}
`

var methods = template.Must(template.New("methods").Parse(methodTemplate))

const formatTemplate = `
func ({{.Receiver}} {{.Name}}) Format(fs fmt.State, c rune){
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", {{.Receiver}}, float64({{.Receiver}}))
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
		fmt.Fprintf(fs, "%*.*"+string(c), w, p, float64({{.Receiver}}))
		{{if .PrintString}}fmt.Fprint(fs, " {{.PrintString}}"){{end}}
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g{{if .PrintString}} {{.PrintString}}{{end}})", c, {{.Receiver}}, float64({{.Receiver}}))
	return
}
}
`

var form = template.Must(template.New("format").Parse(formatTemplate))

func generate(unit Unit) {
	lowerName := strings.ToLower(unit.Name)
	filename := filepath.Join(unitPkgPath, lowerName+".go")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Need to define new prefixes because text/template can't do math.
	// Need to do math because kilogram = 1 not 10^3

	prefixes := make([]Prefix, len(Prefixes))
	for i, p := range Prefixes {
		prefixes[i].Name = p.Name
		prefixes[i].Power = p.Power + unit.Offset
	}

	data := struct {
		Prefixes []Prefix
		Unit     Unit
	}{
		prefixes,
		unit,
	}

	buf := bytes.NewBuffer(make([]byte, 0))

	err = header.Execute(buf, unit)
	if err != nil {
		log.Fatal(err)
	}
	if data.Unit.Suffix != "" {
		err = prefix.Execute(buf, data)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = methods.Execute(buf, unit)
	if err != nil {
		log.Fatal(err)
	}

	err = form.Execute(buf, unit)
	if err != nil {
		log.Fatal(err)
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		f.Write(buf.Bytes()) // This is here to debug bad format
		log.Fatalf("error formatting: %s", err)
	}

	f.Write(b)
}
