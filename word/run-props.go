package word

import (
	"encoding/xml"
)

// Run props
type RunProperties struct {
	XMLName xml.Name `xml:"w:rPr"`
	Content []interface{}
	B       *Bold
	I       *Italics
	Sz      *Fontsize
	SzCs    *ComplexScriptFontsize
}

func NewRunProperties() *RunProperties {
	return &RunProperties{Content: make([]interface{}, 0)}
}

func (r *RunProperties) AddFont() *RunFonts {
	obj := NewRunFonts()
	r.Content = append(r.Content, obj)
	return obj
}

func (r *RunProperties) Bold(bold bool) {
	if bold {
		if r.B == nil {
			r.B = &Bold{Val: "true"}
		} else {
			r.B.Val = "true"
		}
	}
}

func (r *RunProperties) Fontsize(size string) {
	if r.Sz == nil {
		r.Sz = &Fontsize{Val: size}
	} else {
		r.Sz.Val = size
	}
}

func (r *RunProperties) ComplexScriptFontsize(size string) {
	if r.SzCs == nil {
		r.SzCs = &ComplexScriptFontsize{Val: size}
	} else {
		r.SzCs.Val = size
	}
}