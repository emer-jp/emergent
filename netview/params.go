// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netview

import (
	"github.com/emer/etable/minmax"
	"github.com/goki/gi/giv"
)

// Params holds parameters controlling how the view is rendered
type Params struct {
	MaxRecs   int              `min:"1" desc:"maximum number of records to store to enable rewinding through prior states"`
	UnitSize  float32          `min:"0.1" max:"1" step:"0.1" def:"0.9" desc:"size of a single unit, where 1 = full width and no space.. .9 default"`
	LayNmSize float32          `min:"0.01" max:".1" step:"0.01" def:"0.05" desc:"size of the layer name labels -- entire network view is unit sized"`
	ColorMap  giv.ColorMapName `desc:"name of color map to use"`
	ZeroAlpha float32          `min:"0" max:"1" step:"0.1" def:"0.4" desc:"opacity (0-1) of zero values -- greater magnitude values become increasingly opaque on either side of this minimum"`
	NetView   *NetView         `copy:"-" json:"-" xml:"-" view:"-" desc:"our netview, for update method"`
}

func (nv *Params) Defaults() {
	if nv.MaxRecs == 0 {
		nv.MaxRecs = 100
	}
	if nv.UnitSize == 0 {
		nv.UnitSize = .9
	}
	if nv.LayNmSize == 0 {
		nv.LayNmSize = .05
	}
	if nv.ZeroAlpha == 0 {
		nv.ZeroAlpha = 0.4
	}
	if nv.ColorMap == "" {
		nv.ColorMap = giv.ColorMapName("ColdHot")
	}
}

// Update satisfies the gi.Updater interface and will trigger display update on edits
func (nv *Params) Update() {
	if nv.NetView != nil {
		nv.NetView.Config()
		nv.NetView.Update()
	}
}

// VarParams holds parameters for display of each variable
type VarParams struct {
	Var     string         `desc:"name of the variable"`
	ZeroCtr bool           `desc:"keep Min - Max centered around 0, and use negative heights for units -- else use full min-max range for height (no negative heights)"`
	Range   minmax.Range32 `view:"inline" desc:"range to display"`
	MinMax  minmax.F32     `view:"inline" desc:"if not using fixed range, this is the actual range of data"`
}

// Defaults sets default values if otherwise not set
func (vp *VarParams) Defaults() {
	if vp.Range.Max == 0 && vp.Range.Min == 0 {
		vp.ZeroCtr = true
		vp.Range.SetMin(-1)
		vp.Range.SetMax(1)
	}
}
