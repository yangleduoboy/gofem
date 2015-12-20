// Copyright 2015 Dorival Pedroso and Raul Durand. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package fld implements models for fluid density
package fld

import (
	"math"

	"github.com/cpmech/gosl/fun"
	"github.com/cpmech/gosl/plt"
	"github.com/cpmech/gosl/utl"
)

// Model implements a model to compute pressure (p) and intrinsic density (R) of a fluid
// along a column with gravity (g). The model is:
//   R(p) = R0 + C・(p - p0)   thus   dR/dp = C
type Model struct {

	// material data
	R0 float64 // intrinsic density corresponding to p0
	P0 float64 // pressure corresponding to R0
	C  float64 // compressibility coefficient; e.g. R0/Kbulk or M/(R・θ)

	// additional data
	Grav float64 // gravity acceleration (positive constant)
	H    float64 // elevation where (R0,p0) is known
}

// Init initialises this structure
func (o *Model) Init(prms fun.Prms, H, grav float64) {
	for _, p := range prms {
		switch p.N {
		case "R0":
			o.R0 = p.V
		case "P0":
			o.P0 = p.V
		case "C":
			o.C = p.V
		}
	}
	o.Grav = grav
	o.H = H
}

// GetPrms gets (an example of) parameters
//  Input:
//   example -- returns example of parameters
//   dryair  -- example of dry air parameters is returned;
//              othewise returns example of water parameters
func (o Model) GetPrms(example, dryair bool) fun.Prms {
	if example {
		if dryair {
			return fun.Prms{ // dry air
				&fun.Prm{N: "R0", V: 0.0012}, // [Mg/m³]
				&fun.Prm{N: "P0", V: 0.0},    // [kPa]
				&fun.Prm{N: "C", V: 1.17e-5}, // [Mg/(m³・kPa)]
			}
		}
		return fun.Prms{ // water
			&fun.Prm{N: "R0", V: 1.0},    // [Mg/m³]
			&fun.Prm{N: "P0", V: 0.0},    // [kPa]
			&fun.Prm{N: "C", V: 4.53e-7}, // [Mg/(m³・kPa)]
		}
	}
	return fun.Prms{
		&fun.Prm{N: "R0", V: o.R0},
		&fun.Prm{N: "P0", V: o.P0},
		&fun.Prm{N: "C", V: o.C},
	}
}

// Calc computes pressure and density
func (o Model) Calc(z float64) (p, R float64) {
	p = o.P0 + (o.R0/o.C)*(math.Exp(o.C*o.Grav*(o.H-z))-1.0)
	R = o.R0 + o.C*(p-o.P0)
	return
}

// Plot plots pressure and density along height of column
func (o Model) Plot(dirout, fnkey, subscript string, np int) {

	Z := utl.LinSpace(0, o.H, np)
	P := make([]float64, np)
	R := make([]float64, np)
	for i, z := range Z {
		P[i], R[i] = o.Calc(z)
	}

	pMaxLin := o.R0 * o.Grav * o.H

	plt.SetForEps(1.2, 400)
	plt.Subplot(2, 1, 1)
	plt.Plot(P, Z, "'b-', clip_on=0")
	plt.Plot([]float64{o.P0, pMaxLin}, []float64{o.H, 0}, "'k--', color='gray'")
	plt.Gll("$p_{"+subscript+"}$", "$z$", "")

	plt.Subplot(2, 1, 2)
	plt.Plot(R, Z, "'r-', clip_on=0")
	plt.Plot([]float64{o.R0, o.R0 + o.C*pMaxLin}, []float64{o.H, 0}, "'k--', color='gray'")
	plt.Gll("$\\rho_{"+subscript+"}$", "$z$", "")
	plt.TicksNoOffset()

	plt.SaveD(dirout, fnkey+".eps")
}