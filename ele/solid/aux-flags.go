// Copyright 2016 The Gofem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be fndim in the LICENSE file.

package solid

import (
	"math"

	"github.com/cpmech/gosl/io"
)

func GetSolidFlags(axisym, pstress bool, extra string) (useB, debug bool, thickness float64) {

	// defaults
	useB = false
	debug = false
	thickness = 1.0

	// flag: use B matrix
	if s_useB, found := io.Keycode(extra, "useB"); found {
		useB = io.Atob(s_useB)
	}

	// fix useB flag in case of axisymmetric simulation
	if axisym {
		useB = true
	}

	// flag: thickess
	if s_thick, found := io.Keycode(extra, "thick"); found {
		thickness = io.Atof(s_thick)
	}

	// flag: debug
	if s_debug, found := io.Keycode(extra, "debug"); found {
		debug = io.Atob(s_debug)
	}
	return
}

func GetContactFaceFlags(extra string) (Macaulay bool, BetRamp, Kappa float64) {

	// defaults
	Macaulay = false
	BetRamp = math.Ln2 / 0.01
	Kappa = 1.0

	// use macaulay function ?
	if s_mac, found := io.Keycode(extra, "mac"); found {
		Macaulay = io.Atob(s_mac)
	}

	// coefficient for smooth ramp function
	if s_bet, found := io.Keycode(extra, "bet"); found {
		BetRamp = io.Atof(s_bet)
	}

	// κ coefficient
	if s_kap, found := io.Keycode(extra, "kap"); found {
		Kappa = io.Atof(s_kap)
	}
	return
}
