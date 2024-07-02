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

			// {Term{"author", String, "", []IndexOption{}}, Displayable, "author_teim", "failed string/searchable"},
			// {Term{"author", String, "", []IndexOption{}}, Searchable, "author_teim", "failed string/searchable"},
			// {Term{"author", String, "", []IndexOption{}}, Displayable, "author_teim", "failed string/searchable"},
		}

		for _, scenario := range scenarios {
			got, _ := GenFieldName(scenario.Term, scenario.IndexOption)
			if got != scenario.Expected {
				t.Errorf("unexpected result: %s: want: '%s', got: '%s'", scenario.Comment, scenario.Expected, got)
			}
		}
	})
}
