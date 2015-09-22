// Copyright 2015 Dorival Pedroso and Raul Durand. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inp

import (
	"os"
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/plt"
)

func Test_msh01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("msh01")

	msh, err := ReadMsh("data", "bh16.msh", 0)
	if err != nil {
		tst.Errorf("test failed:\n%v", err)
		return
	}
	io.Pforan("%v\n", msh)
	io.Pfcyan("lims = [%g, %g, %g, %g, %g, %g]\n", msh.Xmin, msh.Xmax, msh.Ymin, msh.Ymax, msh.Zmin, msh.Zmax)
	chk.Scalar(tst, "xmin", 1e-17, msh.Xmin, 10)
	chk.Scalar(tst, "xmax", 1e-17, msh.Xmax, 14)
	chk.Scalar(tst, "ymin", 1e-17, msh.Ymin, -1)
	chk.Scalar(tst, "ymax", 1e-17, msh.Ymax, 1)

	if chk.Verbose {
		msh.Draw2d(false)
		plt.SaveD("/tmp/gofem", "test_msh01.png")
	}
}

func Test_sim01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sim01")

	sim := ReadSim("data/bh16.sim", "", true, 0)
	if sim == nil {
		tst.Errorf("test failed:\n")
		return
	}
	if chk.Verbose {
		sim.GetInfo(os.Stdout)
		io.Pf("\n")
	}

	io.Pfyel("ndim    = %v\n", sim.Ndim)
	io.Pfyel("maxElev = %v\n", sim.MaxElev)
	io.Pfyel("grav    = %v\n", sim.Gravity.F(0, nil))
	io.Pfyel("Wrho0   = %v\n", sim.WaterRho0)
	io.Pfyel("Wbulk   = %v\n", sim.WaterBulk)
	io.Pfyel("Wlevel  = %v\n", sim.WaterLevel)
}

func Test_sim02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("sim01")

	sim := ReadSim("data/frees01.sim", "", true, 0)
	if sim == nil {
		tst.Errorf("test failed:\n")
		return
	}
	if chk.Verbose {
		sim.GetInfo(os.Stdout)
		io.Pf("\n")
	}

	io.Pfyel("ndim    = %v\n", sim.Ndim)
	io.Pfyel("maxElev = %v\n", sim.MaxElev)
	io.Pfyel("grav    = %v\n", sim.Gravity.F(0, nil))
	io.Pfyel("Wrho0   = %v\n", sim.WaterRho0)
	io.Pfyel("Wbulk   = %v\n", sim.WaterBulk)
	io.Pfyel("Wlevel  = %v\n", sim.WaterLevel)
}

func Test_mat01(tst *testing.T) {

	chk.PrintTitle("mat01")

	mdb1 := ReadMat("data", "bh.mat")
	if mdb1 == nil {
		tst.Errorf("test failed\n")
		return
	}
	io.Pforan("bh.mat just read:\n%v\n", mdb1)

	fn := "test_bh.mat"
	io.WriteFileSD("/tmp/gofem/inp", fn, mdb1.String())

	mdb2 := ReadMat("/tmp/gofem/inp/", fn)
	if mdb2 == nil {
		tst.Errorf("test failed\n")
		return
	}
	io.Pfblue2("\n%v\n", mdb2)
}

func Test_mat02(tst *testing.T) {

	chk.PrintTitle("mat02 (conversion)")

	convertsymbols := true
	MatfileOld2New("/tmp/gofem/inp", "new_layers.mat", "data/old_layers.mat", convertsymbols)

	mdb := ReadMat("/tmp/gofem/inp/", "new_layers.mat")
	if mdb == nil {
		tst.Errorf("test failed\n")
		return
	}
	io.Pfblue2("%v\n", mdb)
}

func Test_mat03(tst *testing.T) {

	chk.PrintTitle("mat03 (inverse conversion)")

	convertsymbols := true
	MatfileNew2Old("/tmp/gofem/inp", "converted_porous.mat", "data/porous.mat", convertsymbols)
}
