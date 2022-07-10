package xsd

import (
	"encoding/xml"
)

type Sequence struct {
	XMLName     xml.Name  `xml:"http://www.w3.org/2001/XMLSchema sequence"`
	MinOccurs   string    `xml:"minOccurs,attr"`
	MaxOccurs   string    `xml:"maxOccurs,attr"`
	ElementList []Element `xml:"element"`
	Choices     []Choice  `xml:"choice"`
	allElements []Element `xml:"-"`
}

func (s *Sequence) Elements() []Element {
	return s.allElements
}

func (s *Sequence) compile(sch *Schema, parentElement *Element) {
	for idx := range s.ElementList {
		el := &s.ElementList[idx]
		if el.Ref != "" && len(s.ElementList) == 1 {
			if el.MaxOccurs == "" {
				el.MaxOccurs = s.MaxOccurs
			}
			if el.MinOccurs == "" {
				el.MinOccurs = s.MinOccurs
			}
		}
		el.compile(sch, parentElement)
	}
	s.allElements = s.ElementList

	for idx := range s.Choices {
		c := &s.Choices[idx]
		c.compile(sch, parentElement)

		s.allElements = append(s.allElements, c.Elements()...)
	}
}

type SequenceAll struct {
	XMLName     xml.Name  `xml:"http://www.w3.org/2001/XMLSchema all"`
	ElementList []Element `xml:"element"`
	Choices     []Choice  `xml:"choice"`
	allElements []Element `xml:"-"`
}

func (s *SequenceAll) Elements() []Element {
	return s.allElements
}

func (s *SequenceAll) compile(sch *Schema, parentElement *Element) {
	for idx := range s.ElementList {
		el := &s.ElementList[idx]
		el.compile(sch, parentElement)
	}
	s.allElements = s.ElementList

	for idx := range s.Choices {
		c := &s.Choices[idx]
		c.compile(sch, parentElement)

		s.allElements = append(s.allElements, c.Elements()...)
	}
}
