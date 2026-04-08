// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"gooxml"
	"gooxml/schema/soo/wml"
)

// NumberingLevel is the definition for numbering for a particular level within
// a NumberingDefinition.
type NumberingLevel struct {
	x *wml.CT_Lvl
}

// X returns the inner wrapped XML type.
func (n NumberingLevel) X() *wml.CT_Lvl {
	return n.x
}

// SetFormat sets the numbering format.
func (n NumberingLevel) SetFormat(f wml.ST_NumberFormat) {
	if n.x.NumFmt == nil {
		n.x.NumFmt = wml.NewCT_NumFmt()
	}
	n.x.NumFmt.ValAttr = f
}

// SetText sets the text to be used in bullet mode.
func (n NumberingLevel) SetText(t string) {
	if t == "" {
		n.x.LvlText = nil
	} else {
		n.x.LvlText = wml.NewCT_LevelText()
		n.x.LvlText.ValAttr = gooxml.String(t)
	}
}

// Properties returns the numbering level paragraph properties.
func (n NumberingLevel) Properties() ParagraphStyleProperties {
	if n.x.PPr == nil {
		n.x.PPr = wml.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{n.x.PPr}
}

// SetAlignment sets the paragraph alignment
func (n NumberingLevel) SetAlignment(j wml.ST_Jc) {
	if j == wml.ST_JcUnset {
		n.x.LvlJc = nil
	} else {
		n.x.LvlJc = wml.NewCT_Jc()
		n.x.LvlJc.ValAttr = j
	}
}

// RunProperties returns the RunProperties controlling numbering level font, etc.
func (n NumberingLevel) RunProperties() RunProperties {
	if n.x.RPr == nil {
		n.x.RPr = wml.NewCT_RPr()
	}
	return RunProperties{n.x.RPr}
}

// SetStart sets the starting number for this level.
func (n NumberingLevel) SetStart(start int64) {
	if n.x.Start == nil {
		n.x.Start = wml.NewCT_DecimalNumber()
	}
	n.x.Start.ValAttr = start
}

// Start returns the starting number for this level, or 1 if not set.
func (n NumberingLevel) Start() int64 {
	if n.x.Start == nil {
		return 1
	}
	return n.x.Start.ValAttr
}

// SetRestartLevel sets the level at which numbering restarts.
// If restartLevel is -1, numbering is continuous (never restarts).
// Otherwise, numbering restarts when the specified level (0-based) changes.
func (n NumberingLevel) SetRestartLevel(restartLevel int) {
	if restartLevel < 0 {
		// Continuous numbering - remove restart
		n.x.LvlRestart = nil
		return
	}
	if n.x.LvlRestart == nil {
		n.x.LvlRestart = wml.NewCT_DecimalNumber()
	}
	n.x.LvlRestart.ValAttr = int64(restartLevel)
}

// RestartLevel returns the level at which numbering restarts, or -1 for continuous.
func (n NumberingLevel) RestartLevel() int {
	if n.x.LvlRestart == nil {
		return -1
	}
	return int(n.x.LvlRestart.ValAttr)
}
