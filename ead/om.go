package ead

import "encoding/xml"

type IndexOption int64

const (
	Displayable IndexOption = iota
	Facetable
	Searchable
	Sortable
	StoredSortable
)

type Term struct {
	Name    string
	XPath   string
	IndexAs []IndexOption
}

type Terminology struct {
	Root  string
	Terms []Term
}

type Field struct {
	XMLName xml.Name `xml:"field"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

type SolrDoc struct {
	Fields []Field `xml:"doc"`
}

var EADTerminology = Terminology{
	Root: "ead",
	Terms: []Term{
		{"author", `filedesc/titlestmt/author`, []IndexOption{Searchable, Displayable}},
	},
}
