// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"gooxml/schema/soo/wml"
)

// NumberingDefinition defines a numbering definition for a list of pragraphs.
type NumberingDefinition struct {
	x *wml.CT_AbstractNum
}

// X returns the inner wrapped XML type.
func (n NumberingDefinition) X() *wml.CT_AbstractNum {
	return n.x
}

// AbstractNumberID returns the ID that is unique within all numbering
// definitions that is used to assign the definition to a paragraph.
func (n NumberingDefinition) AbstractNumberID() int64 {
	return n.x.AbstractNumIdAttr
}

// Levels returns all of the numbering levels defined in the definition.
func (n NumberingDefinition) Levels() []NumberingLevel {
	ret := []NumberingLevel{}
	for _, nl := range n.x.Lvl {
		ret = append(ret, NumberingLevel{nl})
	}
	return ret
}

// AddLevel adds a new numbering level to a NumberingDefinition.
func (n NumberingDefinition) AddLevel() NumberingLevel {
	nl := wml.NewCT_Lvl()
	nl.Start = &wml.CT_DecimalNumber{ValAttr: 1}
	nl.IlvlAttr = int64(len(n.x.Lvl))
	n.x.Lvl = append(n.x.Lvl, nl)
	return NumberingLevel{nl}
}

// MultiLevelType returns the multilevel type, or ST_MultiLevelTypeUnset if not set.
func (n NumberingDefinition) MultiLevelType() wml.ST_MultiLevelType {
	if n.x.MultiLevelType != nil {
		return n.x.MultiLevelType.ValAttr
	} else {
		return wml.ST_MultiLevelTypeUnset
	}
}

// SetMultiLevelType sets the multilevel type.
func (n NumberingDefinition) SetMultiLevelType(t wml.ST_MultiLevelType) {
	if t == wml.ST_MultiLevelTypeUnset {
		n.x.MultiLevelType = nil
	} else {
		n.x.MultiLevelType = wml.NewCT_MultiLevelType()
		n.x.MultiLevelType.ValAttr = t
	}
}

// SetName sets the name of the numbering definition.
func (n NumberingDefinition) SetName(name string) {
	if name == "" {
		n.x.Name = nil
	} else {
		n.x.Name = wml.NewCT_String()
		n.x.Name.ValAttr = name
	}
}

// Name returns the name of the numbering definition, or empty string if not set.
func (n NumberingDefinition) Name() string {
	if n.x.Name == nil {
		return ""
	}
	return n.x.Name.ValAttr
}

// Level returns the numbering level at the specified index, or creates it if it doesn't exist.
func (n NumberingDefinition) Level(index int) NumberingLevel {
	if index < 0 {
		index = 0
	}
	// Ensure the level exists
	for len(n.x.Lvl) <= index {
		n.AddLevel()
	}
	return NumberingLevel{n.x.Lvl[index]}
}

// SetLevelStart sets the starting number for a specific level.
func (n NumberingDefinition) SetLevelStart(level int, start int64) {
	lvl := n.Level(level)
	lvl.SetStart(start)
}

// SetLevelFormat sets the numbering format for a specific level.
func (n NumberingDefinition) SetLevelFormat(level int, format wml.ST_NumberFormat) {
	lvl := n.Level(level)
	lvl.SetFormat(format)
}

// SetLevelText sets the numbering text pattern for a specific level.
// Use %1, %2, etc. to represent the number at each level.
func (n NumberingDefinition) SetLevelText(level int, text string) {
	lvl := n.Level(level)
	lvl.SetText(text)
}
