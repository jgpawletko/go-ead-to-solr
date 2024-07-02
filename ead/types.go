package ead

import "encoding/xml"

type IndexOption int64

const (
	Dateable IndexOption = iota
	Displayable
	Facetable
	Searchable
	Sortable
	StoredSearchable
	StoredSortable
	Symbol
)

type DataType int64

const (
	String DataType = iota
	Text
	Date
	Int
)

type Term struct {
	Name     string
	DataType DataType
	XPath    string
	IndexAs  []IndexOption
}

type Terminology struct {
	Root  string
	Terms []Term
}

type Field struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type SolrDoc struct {
	XMLName xml.Name `xml:"doc"`
	Fields  []Field  `xml:"field"`
}

type SolrAdd struct {
	XMLName xml.Name `xml:"add"`
	SolrDoc *SolrDoc
}
