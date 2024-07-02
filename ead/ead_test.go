package ead

import (
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
			{Term{"foo", String, "", []IndexOption{}}, StoredSearchable, "foo_tesim", "failed String/StoredSearchable"},
			{Term{"bar", Text, "", []IndexOption{}}, StoredSearchable, "bar_tesim", "failed Text/StoredSearchable"},
			{Term{"baz", Date, "", []IndexOption{}}, StoredSearchable, "baz_dtsim", "failed Date/StoredSearchable"},
			{Term{"quux", Int, "", []IndexOption{}}, StoredSearchable, "quux_isim", "failed Int/StoredSearchable"},

			{Term{"foo", String, "", []IndexOption{}}, Searchable, "foo_teim", "failed String/Searchable"},
			{Term{"bar", Text, "", []IndexOption{}}, Searchable, "bar_teim", "failed Text/Searchable"},
			{Term{"baz", Date, "", []IndexOption{}}, Searchable, "baz_dtim", "failed Date/Searchable"},
			{Term{"quux", Int, "", []IndexOption{}}, Searchable, "quux_iim", "failed Int/Searchable"},

			{Term{"foo", String, "", []IndexOption{}}, Dateable, "foo_dtsim", "failed String/Dateable"},
			{Term{"bar", Text, "", []IndexOption{}}, Dateable, "bar_dtsim", "failed Text/Dateable"},
			{Term{"baz", Date, "", []IndexOption{}}, Dateable, "baz_dtsim", "failed Date/Dateable"},
			{Term{"quux", Int, "", []IndexOption{}}, Dateable, "quux_dtsim", "failed Int/Dateable"},

			{Term{"foo", String, "", []IndexOption{}}, Facetable, "foo_sim", "failed String/Facetable"},
			{Term{"bar", Text, "", []IndexOption{}}, Facetable, "bar_sim", "failed Text/Facetable"},
			{Term{"baz", Date, "", []IndexOption{}}, Facetable, "baz_sim", "failed Date/Facetable"},
			{Term{"quux", Int, "", []IndexOption{}}, Facetable, "quux_sim", "failed Int/Facetable"},

			{Term{"foo", String, "", []IndexOption{}}, Sortable, "foo_ssi", "failed String/Sortable"},
			{Term{"bar", Text, "", []IndexOption{}}, Sortable, "bar_tesi", "failed Text/Sortable"},
			{Term{"baz", Date, "", []IndexOption{}}, Sortable, "baz_dtsi", "failed Date/Sortable"},
			{Term{"quux", Int, "", []IndexOption{}}, Sortable, "quux_isi", "failed Int/Sortable"},
		}

		for _, scenario := range scenarios {
			got, _ := GenFieldName(scenario.Term, scenario.IndexOption)
			if got != scenario.Expected {
				t.Errorf("unexpected result: %s: want: '%s', got: '%s'", scenario.Comment, scenario.Expected, got)
			}
		}
	})
}
