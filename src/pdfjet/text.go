package pdfjet

/**
 * text.go
 *
Copyright 2020 Innovatics Inc.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice,
  text list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  text list of conditions and the following disclaimer in the documentation
  and / or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

import (
	"single"
	"strings"
)

// Text structure
// Please see Example_45
type Text struct {
	paragraphs                                       []*Paragraph
	font, fallbackFont                               *Font
	x1, y1, xText, yText, width                      float32
	leading, paragraphLeading, spaceBetweenTextLines float32
	beginParagraphPoints                             [][2]float32
	drawBorder                                       bool
}

// NewText is the constructor.
func NewText(paragraphs []*Paragraph) *Text {
	text := new(Text)
	text.paragraphs = paragraphs
	text.font = paragraphs[0].lines[0].GetFont()
	text.fallbackFont = paragraphs[0].lines[0].GetFallbackFont()
	text.leading = text.font.GetBodyHeight()
	text.paragraphLeading = 2 * text.leading
	text.beginParagraphPoints = make([][2]float32, 0)
	text.spaceBetweenTextLines = text.font.StringWidth(text.fallbackFont, single.Space)
    text.drawBorder = true
	return text
}

// SetLocation sets the location of the text.
func (text *Text) SetLocation(x, y float32) *Text {
	text.x1 = x
	text.y1 = y
	return text
}

// SetWidth sets the width of the text component.
func (text *Text) SetWidth(width float32) *Text {
	text.width = width
	return text
}

// SetLeading sets the leading of the text.
func (text *Text) SetLeading(leading float32) *Text {
	text.leading = leading
	return text
}

// SetParagraphLeading sets the paragraph leading.
func (text *Text) SetParagraphLeading(paragraphLeading float32) *Text {
	text.paragraphLeading = paragraphLeading
	return text
}

// GetBeginParagraphPoints returns the begin paragraph points.
func (text *Text) GetBeginParagraphPoints() [][2]float32 {
	return text.beginParagraphPoints
}

// SetSpaceBetweenTextLines sets the space between text lines.
func (text *Text) SetSpaceBetweenTextLines(spaceBetweenTextLines float32) *Text {
	text.spaceBetweenTextLines = spaceBetweenTextLines
	return text
}

// GetSize returns the size of the text block.
func (text *Text) GetSize() [2]float32 {
	return [2]float32{text.width, (text.yText + text.font.descent) - (text.y1 + text.paragraphLeading)}
}

// DrawOn draws the text on the page.
func (text *Text) DrawOn(page *Page) [2]float32 {
	text.xText = text.x1
	text.yText = text.y1 + text.font.ascent
	for _, paragraph := range text.paragraphs {
		var buf strings.Builder
		for _, textLine := range paragraph.lines {
			buf.WriteString(textLine.text)
		}
		for i, textLine := range paragraph.lines {
			if i == 0 {
				text.beginParagraphPoints = append(text.beginParagraphPoints, [2]float32{text.xText, text.yText})
			}
			altDescription := buf.String()
			actualText := buf.String()
			if i == 0 {
				altDescription = single.Space
				actualText = single.Space
			}
			textLine.SetAltDescription(altDescription)
			textLine.SetActualText(actualText)
			xy := text.drawTextLine(page, text.xText, text.yText, textLine)
			text.xText = xy[0]
			if textLine.GetTrailingSpace() {
				text.xText += text.spaceBetweenTextLines
			}
			text.yText = xy[1]
		}
		text.xText = text.x1
		text.yText += text.paragraphLeading
	}

	height := ((text.yText - text.paragraphLeading) - text.y1) + text.font.descent
	if page != nil && text.drawBorder {
		box := NewBox()
		box.SetLocation(text.x1, text.y1)
		box.SetSize(text.width, height)
		box.DrawOn(page)
	}

	return [2]float32{text.x1 + text.width, text.y1 + height}
}

func (text *Text) drawTextLine(page *Page, x, y float32, textLine *TextLine) []float32 {
	text.xText = x
	text.yText = y

	var tokens []string
	if text.stringIsCJK(textLine.text) {
		tokens = text.tokenizeCJK(textLine, text.width)
	} else {
		tokens = strings.Fields(textLine.text)
	}

	firstTextSegment := true
	var buf strings.Builder
	for i, token := range tokens {
		if i > 0 {
			token = single.Space + tokens[i]
		}
		lineWidth := textLine.font.StringWidth(textLine.fallbackFont, buf.String())
		tokenWidth := textLine.font.StringWidth(textLine.fallbackFont, token)
		if (lineWidth + tokenWidth) < (text.x1+text.width)-text.xText {
			buf.WriteString(token)
		} else {
			if page != nil {
				altDescription := single.Space
				actualText := single.Space
				if firstTextSegment {
					altDescription = textLine.GetAltDescription()
					actualText = textLine.GetActualText()
				}
				textLine2 := NewTextLine(textLine.font, buf.String())
				textLine2.SetFallbackFont(textLine.fallbackFont)
				textLine2.SetLocation(text.xText, text.yText+textLine.GetVerticalOffset())
				textLine2.SetColor(textLine.GetColor())
				textLine2.SetUnderline(textLine.GetUnderline())
				textLine2.SetStrikeout(textLine.GetStrikeout())
				textLine2.SetLanguage(textLine.GetLanguage())
				textLine2.SetAltDescription(altDescription)
				textLine2.SetActualText(actualText)
				textLine2.DrawOn(page)
			}
			firstTextSegment = false
			text.xText = text.x1
			text.yText += text.leading
			buf.Reset()
			buf.WriteString(tokens[i])
		}
	}
	if page != nil {
		altDescription := single.Space
		actualText := single.Space
		if firstTextSegment {
			altDescription = textLine.GetAltDescription()
			actualText = textLine.GetActualText()
		}
		textLine2 := NewTextLine(textLine.font, buf.String())
		textLine2.SetFallbackFont(textLine.fallbackFont)
		textLine2.SetLocation(text.xText, text.yText+textLine.GetVerticalOffset())
		textLine2.SetColor(textLine.GetColor())
		textLine2.SetUnderline(textLine.GetUnderline())
		textLine2.SetStrikeout(textLine.GetStrikeout())
		textLine2.SetLanguage(textLine.GetLanguage())
		textLine2.SetAltDescription(altDescription)
		textLine2.SetActualText(actualText)
		textLine2.DrawOn(page)
	}

	return []float32{text.xText + textLine.font.StringWidth(textLine.fallbackFont, buf.String()), text.yText}
}

func (text *Text) stringIsCJK(str string) bool {
	// CJK Unified Ideographs Range: 4E00–9FD5
	// Hiragana Range: 3040–309F
	// Katakana Range: 30A0–30FF
	// Hangul Jamo Range: 1100–11FF
	numOfCJK := 0
	runes := []rune(str)
	for _, ch := range runes {
		if (ch >= 0x4E00 && ch <= 0x9FD5) ||
			(ch >= 0x3040 && ch <= 0x309F) ||
			(ch >= 0x30A0 && ch <= 0x30FF) ||
			(ch >= 0x1100 && ch <= 0x11FF) {
			numOfCJK++
		}
	}
	return numOfCJK > (len(runes) / 2)
}

func (text *Text) tokenizeCJK(textLine *TextLine, textWidth float32) []string {
	tokens := make([]string, 0)
	var sb strings.Builder
	runes := []rune(textLine.text)
	for _, ch := range runes {
		if text.font.StringWidth(text.fallbackFont, sb.String() + string(ch)) < textWidth {
			sb.WriteRune(ch)
		} else {
			tokens = append(tokens, sb.String())
			sb.Reset()
			sb.WriteRune(ch)
		}
	}
	if len(sb.String()) > 0 {
		tokens = append(tokens, sb.String())
	}
	return tokens
}
