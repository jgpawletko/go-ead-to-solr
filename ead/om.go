package ead

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lestrrat-go/libxml2/parser"
	"github.com/lestrrat-go/libxml2/xpath"
)

var xpathFunctionMatcher = regexp.MustCompile(`\/?\w+\(\)`) // matches functions in xpath
var xpathWordMatcher = regexp.MustCompile(`^(\w+).*\/?`)    // matches words in xpath

// xpathToExpression is a helper function to add the default namespace to an xpath so that it can be used in a XPath query
func XpathToExpression(xp string) string {
	// https://go.dev/play/p/-iILPI4g6q6
	expr := ""
	substrings := strings.SplitAfter(xp, "/")
	for _, ss := range substrings {

		// if it's a slash, echo it
		if ss == "/" {
			expr += "/"
			continue
		}

		// if it's a function, echo it
		b := xpathFunctionMatcher.Match([]byte(ss))
		if b {
			expr += ss
			continue
		}

		// if it's a word, convert it
		if xpathWordMatcher.Match([]byte(ss)) {
			expr = expr + "_:" + ss
			continue
		}

		// if it's none of the above, echo it
		expr += ss
	}
	return expr
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
		exprString := XpathToExpression(`//` + term.XPath)
		nodes := xpath.NodeList(ctx.Find(exprString))

		for _, indexOption := range term.IndexAs {
			for _, n := range nodes {
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

func (sd *SolrDoc) AddField(name string, dt DataType, value string, indexOptions []IndexOption) error {
	for _, indexOption := range indexOptions {
		solrFieldName, err := GenFieldName(Term{Name: name, DataType: dt}, indexOption)
		if err != nil {
			return err
		}
		sd.Fields = append(sd.Fields, Field{Name: solrFieldName, Value: value})
	}
	return nil
}
