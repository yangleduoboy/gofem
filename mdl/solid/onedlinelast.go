// Copyright 2016 The Gofem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package solid

import "github.com/cpmech/gosl/fun/dbf"

// OnedLinElast implements a linear elastic model for 1D elements
type OnedLinElast struct {
	E   float64 // Young's modulus
	G   float64 // shear modulus
	A   float64 // cross-sectional area
	I22 float64 // moment of inertia of cross section about y2-axis
	I11 float64 // moment of inertia of cross section about y1-axis
	Jtt float64 // torsional constant
	Rho float64 // density
}

// add model to factory
func init() {
	allocators["oned-elast"] = func() Model { return new(OnedLinElast) }
}

// Free frees memory
func (o *OnedLinElast) Free() {
}

// GetRho returns density
func (o *OnedLinElast) GetRho() float64 {
	return o.Rho
}

// GetA returns cross-sectional area
func (o *OnedLinElast) GetA() float64 {
	return o.A
}

// Init initialises model
func (o *OnedLinElast) Init(ndim int, pstress bool, prms dbf.Params) (err error) {
	prms.Connect(&o.E, "E", "oned-elast model")
	prms.Connect(&o.G, "G", "oned-elast model")
	prms.Connect(&o.A, "A", "oned-elast model")
	prms.Connect(&o.I22, "I22", "oned-elast model")
	prms.Connect(&o.I11, "I11", "oned-elast model")
	prms.Connect(&o.Jtt, "Jtt", "oned-elast model")
	prms.Connect(&o.Rho, "rho", "oned-elast model")
	return
}

// InitIntVars: unused
func (o *OnedLinElast) InitIntVars(σ []float64) (s *State, err error) {
	return
}

// GetPrms gets (an example) of parameters
func (o OnedLinElast) GetPrms() dbf.Params {
	return []*dbf.P{
		&dbf.P{N: "E", V: 2.0000e+08},
		&dbf.P{N: "G", V: 7.5758e+07},
		&dbf.P{N: "A", V: 1.0000e-02},
		&dbf.P{N: "I22", V: 8.3333e-06},
		&dbf.P{N: "I11", V: 8.3333e-06},
		&dbf.P{N: "Jtt", V: 1.4063e-05},
		&dbf.P{N: "rho", V: 7.8500e+00},
	}
}

// InitIntVars initialises internal (secondary) variables
func (o OnedLinElast) InitIntVars1D() (s *OnedState, err error) {
	s = NewOnedState(0, 0)
	return
}

// Update updates stresses for given strains
func (o OnedLinElast) Update(s *OnedState, ε, Δε, aux float64) (err error) {
	s.Sig += o.E * Δε
	return
}

// CalcD computes D = dσ_new/dε_new consistent with StressUpdate
func (o OnedLinElast) CalcD(s *OnedState, firstIt bool) (float64, float64, error) {
	return o.E, 0, nil
}
