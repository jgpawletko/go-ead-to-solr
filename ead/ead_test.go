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
			{Term{"foo", String, "", []IndexOption{}}, Searchable, "foo_teim", "failed string/searchable"},
		}

		for _, scenario := range scenarios {
			got, _ := GenFieldName(scenario.Term, scenario.IndexOption)
			if got != scenario.Expected {
				t.Errorf("unexpected result: %s: want: '%s', got: '%s'", scenario.Comment, scenario.Expected, got)
			}
		}
	})
}
