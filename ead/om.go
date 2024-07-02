package ead

import (
	"encoding/xml"
	"fmt"
	"regexp"

	"github.com/lestrrat-go/libxml2/parser"
	"github.com/lestrrat-go/libxml2/xpath"
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

// xpathToExpression is a helper function to add the default namespace to an xpath so that it can be used in a XPath query
func XpathToExpression(s string) string {
	defaultNameSpaceMatcher := regexp.MustCompile(`/(\w+)`)
	s = defaultNameSpaceMatcher.ReplaceAllString(s, `/_:$1`)
	return s
}

// this is a bit ugly, but unfortunately there is coupling between the data type and the index option...
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
		return "", fmt.Errorf("unsupported data type")
	}

	switch indexOption {
	// the following cases ARE data type dependent,
	// so the suffix appended to
	// additionally, for backward compatibility, the
	// Strings are treated like Text
	case StoredSearchable:
		if dataType == String {
			suffix = "te"
		}
		suffix += "sim"

	case Searchable:
		if dataType == String {
			suffix = "te"
		}
		suffix += "im"

	// the following cases ARE data type dependent,
	// so the IndexOption string is APPENDED to the suffix
	case Sortable:
		suffix += "i"

	case StoredSortable:
		suffix += "si"

	// the following cases ARE NOT data type dependent,
	// so the IndexOption string is ASSIGNED to the suffix
	case Displayable:
		suffix = "ssm"
	case Dateable:
		suffix = "dtsim"
	case Facetable:
		suffix = "sim"
	case Symbol:
		suffix = "ssim"
	default:
		return "", fmt.Errorf("unsupported index option")
	}

	return name + "_" + suffix, nil
}

// GenSolrDoc generates a Solr document from an EAD XML document.
func GenSolrDoc(EADXML []byte, t Terminology) (*SolrDoc, []string) {

	solrDoc := SolrDoc{}

	var errors = []string{}

	p := parser.New()
	xmlDoc, err := p.Parse(EADXML)
	defer xmlDoc.Free()
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}

	// TODO: use the terminology root value to set the root node?
	root, err := xmlDoc.DocumentElement()
	if err != nil {
		errors = append(errors, fmt.Sprintf("problem extracting root node: %s", err.Error()))
		return nil, append(errors, err.Error())
	}

	ctx, err := xpath.NewContext(root)
	if err != nil {
		errors = append(errors, "unable to initialize XPathContext")
		return nil, append(errors, err.Error())
	}
	defer ctx.Free()

	// register the default namespace
	prefix := `_`
	nsuri := `urn:isbn:1-931666-22-9`
	if err := ctx.RegisterNS(prefix, nsuri); err != nil {
		errors = append(errors, "failed to register namespace")
		return nil, append(errors, err.Error())
	}

	for _, term := range t.Terms {
		//exprString := `/` + t.Root + `/` + term.XPath
		//exprString := `//` + term.XPath
		// exprString := `/_:ead/_:eadheader[1]/_:filedesc[1]/_:titlestmt[1]/_:author[1]`
		// exprString := `/_:ead/_:eadheader/_:filedesc/_:titlestmt/_:author`
		// exprString := `/_:ead/_:filedesc/_:titlestmt/_:author`
		exprString := XpathToExpression(`//` + term.XPath)
		nodes := xpath.NodeList(ctx.Find(exprString))

		for _, n := range nodes {
			for _, indexOption := range term.IndexAs {
				solrFieldName, err := GenFieldName(term, indexOption)
				if err != nil {
					return nil, append(errors, err.Error())
				}
				field := Field{Name: solrFieldName, Value: n.NodeValue()}
				solrDoc.Fields = append(solrDoc.Fields, field)
			}
		}
	}

	return &solrDoc, errors
}
