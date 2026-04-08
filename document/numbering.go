// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"fmt"
	"strings"

	"gooxml"
	"gooxml/schema/soo/ofc/sharedTypes"
	"gooxml/schema/soo/wml"
)

// Numbering is the document wide numbering styles contained in numbering.xml.
type Numbering struct {
	x *wml.Numbering
}

// NewNumbering constructs a new numbering.
func NewNumbering() Numbering {
	n := wml.NewNumbering()
	return Numbering{n}
}

// X returns the inner wrapped XML type.
func (n Numbering) X() *wml.Numbering {
	return n.x
}

// Clear resets the numbering.
func (n Numbering) Clear() {
	n.x.AbstractNum = nil
	n.x.Num = nil
	n.x.NumIdMacAtCleanup = nil
	n.x.NumPicBullet = nil
}

// InitializeDefault constructs a default numbering with bullet style.
func (n Numbering) InitializeDefault() {
	n.InitializeBullet()
}

// InitializeBullet constructs a default numbering with bullet style.
func (n Numbering) InitializeBullet() {
	abs := wml.NewCT_AbstractNum()
	abs.MultiLevelType = wml.NewCT_MultiLevelType()
	abs.MultiLevelType.ValAttr = wml.ST_MultiLevelTypeHybridMultilevel

	n.x.AbstractNum = append(n.x.AbstractNum, abs)
	abs.AbstractNumIdAttr = 1
	const indentStart = 720
	const indentDelta = 720
	const hangingIndent = 360
	for i := 0; i < 9; i++ {
		lvl := wml.NewCT_Lvl()
		lvl.IlvlAttr = int64(i)
		lvl.Start = wml.NewCT_DecimalNumber()
		lvl.Start.ValAttr = 1

		lvl.NumFmt = wml.NewCT_NumFmt()
		lvl.NumFmt.ValAttr = wml.ST_NumberFormatBullet

		lvl.Suff = wml.NewCT_LevelSuffix()
		lvl.Suff.ValAttr = wml.ST_LevelSuffixNothing

		lvl.LvlText = wml.NewCT_LevelText()
		lvl.LvlText.ValAttr = gooxml.String("")

		lvl.LvlJc = wml.NewCT_Jc()
		lvl.LvlJc.ValAttr = wml.ST_JcLeft

		lvl.RPr = wml.NewCT_RPr()
		lvl.RPr.RFonts = wml.NewCT_Fonts()
		lvl.RPr.RFonts.AsciiAttr = gooxml.String("Symbol")
		lvl.RPr.RFonts.HAnsiAttr = gooxml.String("Symbol")
		lvl.RPr.RFonts.HintAttr = wml.ST_HintDefault

		lvl.PPr = wml.NewCT_PPrGeneral()

		indent := int64(i*indentDelta + indentStart)
		lvl.PPr.Ind = wml.NewCT_Ind()
		lvl.PPr.Ind.LeftAttr = &wml.ST_SignedTwipsMeasure{}
		lvl.PPr.Ind.LeftAttr.Int64 = gooxml.Int64(indent)
		lvl.PPr.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		lvl.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(hangingIndent))

		abs.Lvl = append(abs.Lvl, lvl)
	}
	num := wml.NewCT_Num()
	num.NumIdAttr = 1
	num.AbstractNumId = wml.NewCT_DecimalNumber()
	num.AbstractNumId.ValAttr = 1
	n.x.Num = append(n.x.Num, num)
}

// InitializeDecimal constructs a numbering with decimal style (1, 2, 3, ...).
func (n Numbering) InitializeDecimal() {
	n.initializeNumbered(wml.ST_NumberFormatDecimal, "%1.")
}

// InitializeRoman constructs a numbering with upper roman style (I, II, III, ...).
func (n Numbering) InitializeRoman() {
	n.initializeNumbered(wml.ST_NumberFormatUpperRoman, "%1.")
}

// InitializeRomanLower constructs a numbering with lower roman style (i, ii, iii, ...).
func (n Numbering) InitializeRomanLower() {
	n.initializeNumbered(wml.ST_NumberFormatLowerRoman, "%1.")
}

// InitializeLetter constructs a numbering with upper letter style (A, B, C, ...).
func (n Numbering) InitializeLetter() {
	n.initializeNumbered(wml.ST_NumberFormatUpperLetter, "%1.")
}

// InitializeLetterLower constructs a numbering with lower letter style (a, b, c, ...).
func (n Numbering) InitializeLetterLower() {
	n.initializeNumbered(wml.ST_NumberFormatLowerLetter, "%1.")
}

// InitializeChinese constructs a numbering with chinese counting style (一, 二, 三, ...).
func (n Numbering) InitializeChinese() {
	n.initializeNumbered(wml.ST_NumberFormatChineseCounting, "%1、")
}

// InitializeDecimalParenthesis constructs a numbering with decimal enclosed parenthesis style ((1), (2), (3), ...).
func (n Numbering) InitializeDecimalParenthesis() {
	n.initializeNumbered(wml.ST_NumberFormatDecimalEnclosedParen, "%1)")
}

// common helper to create numbered list styles
func (n Numbering) initializeNumbered(format wml.ST_NumberFormat, textPattern string) {
	abs := wml.NewCT_AbstractNum()
	abs.MultiLevelType = wml.NewCT_MultiLevelType()
	abs.MultiLevelType.ValAttr = wml.ST_MultiLevelTypeHybridMultilevel

	n.x.AbstractNum = append(n.x.AbstractNum, abs)
	abs.AbstractNumIdAttr = 1
	const indentStart = 720
	const indentDelta = 720
	const hangingIndent = 360
	for i := 0; i < 9; i++ {
		lvl := wml.NewCT_Lvl()
		lvl.IlvlAttr = int64(i)
		lvl.Start = wml.NewCT_DecimalNumber()
		lvl.Start.ValAttr = 1

		lvl.NumFmt = wml.NewCT_NumFmt()
		lvl.NumFmt.ValAttr = format

		lvl.Suff = wml.NewCT_LevelSuffix()
		lvl.Suff.ValAttr = wml.ST_LevelSuffixNothing

		lvl.LvlText = wml.NewCT_LevelText()
		// Replace %1 with the level number placeholder
		placeholder := fmt.Sprintf("%%%d", i+1)
		text := strings.ReplaceAll(textPattern, "%1", placeholder)
		lvl.LvlText.ValAttr = gooxml.String(text)

		lvl.LvlJc = wml.NewCT_Jc()
		lvl.LvlJc.ValAttr = wml.ST_JcLeft

		lvl.PPr = wml.NewCT_PPrGeneral()

		indent := int64(i*indentDelta + indentStart)
		lvl.PPr.Ind = wml.NewCT_Ind()
		lvl.PPr.Ind.LeftAttr = &wml.ST_SignedTwipsMeasure{}
		lvl.PPr.Ind.LeftAttr.Int64 = gooxml.Int64(indent)
		lvl.PPr.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		lvl.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(hangingIndent))

		abs.Lvl = append(abs.Lvl, lvl)
	}
	num := wml.NewCT_Num()
	num.NumIdAttr = 1
	num.AbstractNumId = wml.NewCT_DecimalNumber()
	num.AbstractNumId.ValAttr = 1
	n.x.Num = append(n.x.Num, num)
}

// Definitions returns the defined numbering definitions.
func (n Numbering) Definitions() []NumberingDefinition {
	ret := []NumberingDefinition{}
	for _, n := range n.x.AbstractNum {
		ret = append(ret, NumberingDefinition{n})
	}
	return ret
}

// AddDefinition adds a new numbering definition.
func (n Numbering) AddDefinition() NumberingDefinition {
	nx := wml.NewCT_Num()

	nextID := int64(1)
	for _, nd := range n.Definitions() {
		if nd.AbstractNumberID() >= nextID {
			nextID = nd.AbstractNumberID() + 1
		}
	}
	nx.NumIdAttr = nextID
	nx.AbstractNumId = wml.NewCT_DecimalNumber()
	nx.AbstractNumId.ValAttr = nextID

	an := wml.NewCT_AbstractNum()
	an.AbstractNumIdAttr = nextID

	n.x.AbstractNum = append(n.x.AbstractNum, an)
	n.x.Num = append(n.x.Num, nx)
	return NumberingDefinition{an}
}
