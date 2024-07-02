package ead

import (
	"encoding/xml"
	"fmt"
)

type IndexOption int64

const (
	Dateable IndexOption = iota
	Displayable
	Facetable
	Searchable
	Sortable
	StoredSearchable
	StoredSortable
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
		{"author", String, `filedesc/titlestmt/author`, []IndexOption{Searchable, Displayable}},
	},
}

/* func getXMLDoc(EADXML []byte) (*xml.Document, error) {
	p := parser.New()
	doc, err := p.Parse(EADXML)
	defer doc.Free()
	if err != nil {
		return nil, err
	}
	return doc, nil
}
*/
// reference: https://github.com/samvera/active_fedora/blob/12.0-stable/lib/active_fedora/indexing/default_descriptors.rb
func GenFieldName(t Term, indexOption IndexOption) (string, error) {

	suffix := ""

	name := t.Name
	dataType := t.DataType

	switch dataType {
	case String:
		suffix = "s"
	case Text:
		suffix = "te"
	case Date:
		suffix = "dt"
	case Int:
		suffix = "i"
	default:
		return "", fmt.Errorf("invalid data type")
	}

	switch indexOption {
	case StoredSearchable:
		// backward compatibility weirdness here...
		// Searchable strings are "teim"
		if dataType == String {
			suffix = "te"
		}

		suffix += "sim"

	case Searchable:
		// backward compatibility weirdness here...
		// Searchable strings are "teim"
		if dataType == String {
			suffix = "te"
		}

		suffix += "im"

	case StoredSortable:
		// backward compatibility weirdness here...
		// Searchable strings are "teim"
		if dataType == String {
			suffix = "te"
		}

		suffix += "sim"

	case Displayable:
		// NOTE: this is universal and not based on the DataType
		suffix = "ssm"
	case Dateable:
		// NOTE: this is universal and not based on the DataType
		suffix = "dtsim"
	case Facetable:
		// NOTE: this is universal and not based on the DataType
		suffix = "sim"
	case Sortable:

	default:
		return "", fmt.Errorf("invalid index option")
	}

	return name + "_" + suffix, nil
}

// // GenSolrDoc generates a Solr document from an EAD XML document.
// func GenSolrDoc(EADXML []byte, t Terminology) (SolrDoc, []string) {

// 	solrDoc := SolrDoc{}

// 	var errors = []string{}

// 	xmlDoc, err := getXMLDoc(EADXML)
// 	if err != nil {
// 		errors = append(errors, err.Error())
// 		return solrDoc, errors
// 	}

// 	// TODO: use the terminology root value to set the root node?
// 	root, err := xmlDoc.DocumentElement()
// 	if err != nil {
// 		errors = append(errors, "Unable to extract root node")
// 		return solrDoc, append(errors, err.Error())
// 	}

// 	ctx, err := xpath.NewContext(root)
// 	if err != nil {
// 		errors = append(errors, "Unable to initialize XPathContext")
// 		return solrDoc, append(errors, err.Error())
// 	}
// 	defer ctx.Free()

// 	// register the default namespace
// 	prefix := `_`
// 	nsuri := `urn:isbn:1-931666-22-9`
// 	if err := ctx.RegisterNS(prefix, nsuri); err != nil {
// 		errors = append(errors, "Failed to register namespace")
// 		return solrDoc, append(errors, err.Error())
// 	}

// 	for _, term := range t.Terms {
// 		exprString := `/` + t.Root + `/` + term.XPath
// 		nodes := xpath.NodeList(ctx.Find(exprString))

// 		for _, n := range nodes {
// 			for _, indexOption := range term.IndexAs {
// 				solrFieldName, err := GenFieldName(term, indexOption)
// 				if err != nil {
// 					errors = append(errors, err.Error())
// 					return solrDoc, errors
// 				}
// 				field := Field{Name: solrFieldName, Value: n.String()}
// 				solrDoc.Fields = append(solrDoc.Fields, field)
// 			}
// 		}
// 	}

// 	return solrDoc, errors
// }
