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

// Footnote represents a footnote in the document.
type Footnote struct {
	d *Document
	x *wml.CT_FtnEdn
}

// X returns the inner wrapped XML type.
func (f Footnote) X() *wml.CT_FtnEdn {
	return f.x
}

// ID returns the footnote ID.
func (f Footnote) ID() int64 {
	return f.x.IdAttr
}

// Paragraphs returns the paragraphs in the footnote.
func (f Footnote) Paragraphs() []Paragraph {
	var paragraphs []Paragraph
	for _, block := range f.x.EG_BlockLevelElts {
		if block.EG_ContentBlockContent != nil {
			for _, content := range block.EG_ContentBlockContent {
				if content.P != nil {
					for _, p := range content.P {
						paragraphs = append(paragraphs, Paragraph{f.d, p})
					}
				}
			}
		}
	}
	return paragraphs
}

// AddParagraph adds a new paragraph to the footnote.
func (f Footnote) AddParagraph() Paragraph {
	p := wml.NewCT_P()

	runContent := wml.NewEG_PContent()
	contentRun := wml.NewEG_ContentRunContent()
	runContent.EG_ContentRunContent = append(runContent.EG_ContentRunContent, contentRun)
	p.EG_PContent = append(p.EG_PContent, runContent)

	blockElt := wml.NewEG_BlockLevelElts()
	contentBlock := wml.NewEG_ContentBlockContent()
	contentBlock.P = append(contentBlock.P, p)
	blockElt.EG_ContentBlockContent = append(blockElt.EG_ContentBlockContent, contentBlock)

	f.x.EG_BlockLevelElts = append(f.x.EG_BlockLevelElts, blockElt)

	return Paragraph{f.d, p}
}

// AddFootnoteRef adds a footnote reference mark (the number) to the footnote.
// This should be called at the beginning of the first paragraph to show the footnote number.
// Returns the Run so you can set the style (font, size, etc.) of the number.
func (f Footnote) AddFootnoteRef(p Paragraph) Run {
	xP := p.X()
	if xP == nil {
		return Run{}
	}

	run := wml.NewCT_R()
	ic := wml.NewEG_RunInnerContent()
	ic.FootnoteRef = wml.NewCT_Empty()
	run.EG_RunInnerContent = append(run.EG_RunInnerContent, ic)

	runContent := wml.NewEG_PContent()
	contentRun := wml.NewEG_ContentRunContent()
	contentRun.R = run
	runContent.EG_ContentRunContent = append(runContent.EG_ContentRunContent, contentRun)
	xP.EG_PContent = append(xP.EG_PContent, runContent)

	return Run{d: f.d, x: run}
}
