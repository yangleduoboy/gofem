// Copyright 2016 The Gofem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package solid

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/fun/dbf"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
)

func Test_hyperelast01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("hyperelast01")

	var m HyperElast1
	m.Init(2, false, []*dbf.P{
		&dbf.P{N: "kap", V: 0.05},
		&dbf.P{N: "kapb", V: 20.0},
		&dbf.P{N: "G0", V: 10000},
		&dbf.P{N: "pr", V: 2.0},
		&dbf.P{N: "pt", V: 10.0},
	})
	io.Pforan("m = %+v\n", m)
	/*
		pr := m.pr
		pt := m.pt

		np := 21
		Ev := utl.LinSpace(0, -0.2, np)
		P := make([]float64, np)
		Q := make([]float64, np)
		X := make([]float64, np)
	*/

	/*
		for j, ed := range []float64{0, 0.05, 0.1, 0.15, 0.2} {
			for i, ev := range Ev {
				P[i], Q[i] = m.Calc_pq(ev, ed)
				X[i] = math.Log(1.0 + (P[i]+pt)/pr)
			}
			slope := (Ev[0] - Ev[np-1]) / (X[np-1] - X[0])
			xm := (X[0] + X[np-1]) / 2.0
			ym := (Ev[0]+Ev[np-1])/2.0 - float64(j)*0.01

			plt.Subplot(3, 2, 1)
			plt.Plot(P, Ev, io.Sf("label='$\\\\varepsilon_d=%g$'", ed))
			plt.PlotOne(P[0], Ev[0], "'ro', clip_on=0")
			plt.Gll("$p$", "$\\varepsilon_v$", "")

			plt.Subplot(3, 2, 3)
			plt.Plot(X, Ev, "")
			plt.PlotOne(X[0], Ev[0], "'ro', clip_on=0")
			plt.Text(xm, ym, io.Sf("slope=%g", slope), "")
			plt.Gll("$x=\\log{[1+(p+p_t)/p_r]}$", "$\\varepsilon_v$", "")

			plt.Subplot(3, 2, 5)
			plt.Plot(Q, Ev, "")
			plt.PlotOne(Q[0], Ev[0], "'ro', clip_on=0")
			plt.Gll("$q$", "$\\varepsilon_v$", "")
		}

		Ed := utl.LinSpace(0, -0.2, np)

		for j, ev := range []float64{0, -0.05, -0.1, -0.15, -0.2} {
			for i, ed := range Ed {
				P[i], Q[i] = m.Calc_pq(ev, ed)
				X[i] = math.Log(1.0 + (P[i]+pt)/pr)
			}
			slope := (Ed[0] - Ed[np-1]) / (Q[np-1] - Q[0])
			xm := (Q[0] + Q[np-1]) / 2.0
			ym := (Ed[0]+Ed[np-1])/2.0 - float64(j)*0.01

			plt.Subplot(3, 2, 2)
			plt.Plot(P, Ed, io.Sf("label='$\\\\varepsilon_v=%g$'", ev))
			plt.PlotOne(P[0], Ed[0], "'ro', clip_on=0")
			plt.Gll("$p$", "$\\varepsilon_d$", "")

			plt.Subplot(3, 2, 4)
			plt.Plot(X, Ed, "")
			plt.PlotOne(X[0], Ed[0], "'ro', clip_on=0")
			plt.Gll("$x=\\log{[1+(p+p_t)/p_r]}$", "$\\varepsilon_d$", "")

			plt.Subplot(3, 2, 6)
			plt.Plot(Q, Ed, "")
			plt.PlotOne(Q[0], Ed[0], "'ro', clip_on=0")
			plt.Text(xm, ym, io.Sf("slope=%g", slope), "")
			plt.Gll("$q$", "$\\varepsilon_d$", "")
		}
	*/

	//plt.Show()
}

func Test_hyperelast02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("hyperelast02 (linear)")

	E, ν := 1500.0, 0.25
	K := Calc_K_from_Enu(E, ν)
	G := Calc_G_from_Enu(E, ν)
	io.Pforan("K = %v\n", K)
	io.Pforan("G = %v\n", G)

	var m HyperElast1
	m.Init(2, false, []*dbf.P{
		&dbf.P{N: "K0", V: K},
		&dbf.P{N: "G0", V: G},
		&dbf.P{N: "le", V: 1},
	})
	io.Pforan("m = %+v\n", m)

	ε := []float64{-0.001, -0.002, -0.003}
	σ := make([]float64, 3)
	m.L_update(σ, ε)
	io.Pfblue2("ε = %v\n", ε)
	io.Pfcyan("σ = %v\n", σ)

	D := la.MatAlloc(3, 3)
	m.L_CalcD(D, ε)
	la.PrintMat("D", D, "%14.6f", false)

	tol := 1e-12
	verb := io.Verbose
	chk.DerivVecVec(tst, "D", tol, D, ε, 1e-1, verb, func(f, x []float64) error {
		m.L_update(f, x)
		return nil
	})
}

func Test_hyperelast03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("hyperelast03 (nonlinear)")

	var m HyperElast1
	m.Init(2, false, []*dbf.P{
		&dbf.P{N: "kap", V: 0.05},
		&dbf.P{N: "kapb", V: 20.0},
		&dbf.P{N: "G0", V: 1500},
		&dbf.P{N: "pr", V: 2.2},
		&dbf.P{N: "pt", V: 11.0},
	})
	io.Pforan("m = %+v\n", m)

	ε := []float64{-0.001, -0.002, -0.003}
	σ := make([]float64, 3)
	m.L_update(σ, ε)
	io.Pfblue2("ε = %v\n", ε)
	io.Pfcyan("σ = %v\n", σ)

	D := la.MatAlloc(3, 3)
	m.L_CalcD(D, ε)
	la.PrintMat("D", D, "%14.6f", false)

	tol := 1e-8
	verb := io.Verbose
	chk.DerivVecVec(tst, "D", tol, D, ε, 1e-1, verb, func(f, x []float64) error {
		m.L_update(f, x)
		return nil
	})
}
