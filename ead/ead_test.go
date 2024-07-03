package ead

import (
	"fmt"
	"os"
	"testing"
)

func TestGenFieldName(t *testing.T) {

	type TermTest struct {
		Term        Term
		IndexOption IndexOption
		Expected    string
		Comment     string
	}
	t.Run("Test GenFieldName()", func(t *testing.T) {

		scenarios := []TermTest{
			{Term{"foo", String, "", []IndexOption{}}, Displayable, "foo_ssm", "failed String/Displayable"},
			{Term{"bar", Text, "", []IndexOption{}}, Displayable, "bar_ssm", "failed Text/Displayable"},
			{Term{"baz", Date, "", []IndexOption{}}, Displayable, "baz_ssm", "failed Date/Displayable"},
			{Term{"quux", Int, "", []IndexOption{}}, Displayable, "quux_ssm", "failed Int/Displayable"},

			{Term{"foo", String, "", []IndexOption{}}, Facetable, "foo_sim", "failed String/Facetable"},
			{Term{"bar", Text, "", []IndexOption{}}, Facetable, "bar_sim", "failed Text/Facetable"},
			{Term{"baz", Date, "", []IndexOption{}}, Facetable, "baz_sim", "failed Date/Facetable"},
			{Term{"quux", Int, "", []IndexOption{}}, Facetable, "quux_sim", "failed Int/Facetable"},

			{Term{"foo", String, "", []IndexOption{}}, Searchable, "foo_teim", "failed String/Searchable"},
			{Term{"bar", Text, "", []IndexOption{}}, Searchable, "bar_teim", "failed Text/Searchable"},
			{Term{"baz", Date, "", []IndexOption{}}, Searchable, "baz_dtim", "failed Date/Searchable"},
			{Term{"quux", Int, "", []IndexOption{}}, Searchable, "quux_iim", "failed Int/Searchable"},

			{Term{"foo", String, "", []IndexOption{}}, Sortable, "foo_si", "failed String/Sortable"},
			{Term{"bar", Text, "", []IndexOption{}}, Sortable, "bar_tei", "failed Text/Sortable"},
			{Term{"baz", Date, "", []IndexOption{}}, Sortable, "baz_dti", "failed Date/Sortable"},
			{Term{"quux", Int, "", []IndexOption{}}, Sortable, "quux_ii", "failed Int/Sortable"},

			{Term{"foo", String, "", []IndexOption{}}, StoredSortable, "foo_ssi", "failed String/StoredSortable"},
			{Term{"bar", Text, "", []IndexOption{}}, StoredSortable, "bar_tesi", "failed Text/StoredSortable"},
			{Term{"baz", Date, "", []IndexOption{}}, StoredSortable, "baz_dtsi", "failed Date/StoredSortable"},
			{Term{"quux", Int, "", []IndexOption{}}, StoredSortable, "quux_isi", "failed Int/StoredSortable"},

			{Term{"foo", String, "", []IndexOption{}}, StoredSearchable, "foo_tesim", "failed String/StoredSearchable"},
			{Term{"bar", Text, "", []IndexOption{}}, StoredSearchable, "bar_tesim", "failed Text/StoredSearchable"},
			{Term{"baz", Date, "", []IndexOption{}}, StoredSearchable, "baz_dtsim", "failed Date/StoredSearchable"},
			{Term{"quux", Int, "", []IndexOption{}}, StoredSearchable, "quux_isim", "failed Int/StoredSearchable"},

			{Term{"foo", String, "", []IndexOption{}}, Dateable, "foo_dtsim", "failed String/Dateable"},
			{Term{"bar", Text, "", []IndexOption{}}, Dateable, "bar_dtsim", "failed Text/Dateable"},
			{Term{"baz", Date, "", []IndexOption{}}, Dateable, "baz_dtsim", "failed Date/Dateable"},
			{Term{"quux", Int, "", []IndexOption{}}, Dateable, "quux_dtsim", "failed Int/Dateable"},
		}

		for _, scenario := range scenarios {
			got, _ := GenFieldName(scenario.Term, scenario.IndexOption)
			if got != scenario.Expected {
				t.Errorf("unexpected result: %s: want: '%s', got: '%s'", scenario.Comment, scenario.Expected, got)
			}
		}
	})
}

func TestGenSolrDoc(t *testing.T) {

	t.Run("Test GenSolrDoc()", func(t *testing.T) {

		EADXML, err := os.ReadFile("./testdata/input/fales/mss_104.xml")
		if err != nil {
			t.Errorf(err.Error())
		}

		data, errors := GenSolrDoc(EADXML, EADTerminology)
		if len(errors) != 0 {
			for _, eMsg := range errors {
				fmt.Printf("%s\n", eMsg)
			}
			t.Errorf("failed to generate Solr document")
		}
		fmt.Printf("%v\n", data)
	})
}

func TestXpathToExpression(t *testing.T) {

	t.Run("Test XpathToExpression()", func(t *testing.T) {

		// input, expected, comment
		scenarios := [][]string{
			{"//filedesc/titlestmt/author", "//_:filedesc/_:titlestmt/_:author", "//filedesc/titlestmt/author"},
			{"//archdesc[@level='collection']/*[name() != 'dsc']//chronlist/chronitem//text()",
				"//_:archdesc[@level='collection']/*[name() != 'dsc']//_:chronlist/_:chronitem//text()",
				"//archdesc[@level='collection']/*[name() != 'dsc']//chronlist/chronitem//text()"},
			{"//archdesc[@level='collection']/did/origination[@label='creator']/*[name() = 'corpname' or name() = 'famname' or name() = 'persname']",
				"//_:archdesc[@level='collection']/_:did/_:origination[@label='creator']/*[name() = 'corpname' or name() = 'famname' or name() = 'persname']",
				"//archdesc[@level='collection']/did/origination[@label='creator']/*[name() = 'corpname' or name() = 'famname' or name() = 'persname']"},
			{"//archdesc[@level='collection']/did/unitdate[not(@type)]",
				"//_:archdesc[@level='collection']/_:did/_:unitdate[not(@type)]",
				"//archdesc[@level='collection']/did/unitdate[not(@type)]"},
			{"archdesc[@level='collection']/did/unitdate/@normal",
				"_:archdesc[@level='collection']/_:did/_:unitdate/@normal",
				"archdesc[@level='collection']/did/unitdate/@normal"},
		}

		for _, scenario := range scenarios {
			got := XpathToExpression(scenario[0])
			if got != scenario[1] {
				t.Errorf("unexpected result: %s\nwant: '%s'\n got: '%s'", scenario[2], scenario[1], got)
			}
		}
	})
}
