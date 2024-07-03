package ead

var EADTerminology = Terminology{
	Root: "ead",
	Terms: []Term{
		{"author", String, `filedesc/titlestmt/author`, []IndexOption{Searchable, Displayable}},

		// Descriptive information in <did>
		{"unittitle", String, `archdesc[@level='collection']/did/unittitle`, []IndexOption{Searchable, Displayable}},
		{"unitid", String, `archdesc[@level='collection']/did/unitid`, []IndexOption{Searchable, Displayable}},
		{"abstract", String, `archdesc[@level='collection']/did/abstract`, []IndexOption{Searchable, Displayable}},
		{"creator", String, `archdesc[@level='collection']/did/origination[@label='creator']/*[name() = 'corpname' or name() = 'famname' or name() = 'persname']`, []IndexOption{Searchable, Displayable}},

		// Dates
		{"unitdate_normal", String, `archdesc[@level='collection']/did/unitdate/@normal`, []IndexOption{Displayable, Searchable, Facetable}},
		{"unitdate", String, `archdesc[@level='collection']/did/unitdate[not(@type)]`, []IndexOption{Searchable}},
		{"unitdate_bulk", String, `archdesc[@level='collection']/did/unitdate[@type='bulk']`, []IndexOption{Searchable}},
		{"unitdate_inclusive", String, `archdesc[@level='collection']/did/unitdate[@type='inclusive']`, []IndexOption{Searchable}},

		// Fulltext in <p> under the following descriptive fields
		{"scopecontent", String, `archdesc[@level='collection']/scopecontent/p`, []IndexOption{Searchable}},
		{"bioghist", String, `archdesc[@level='collection']/bioghist/p`, []IndexOption{Searchable}},
		{"acqinfo", String, `archdesc[@level='collection']/acqinfo/p`, []IndexOption{Searchable}},
		{"custodhist", String, `archdesc[@level='collection']/custodhist/p`, []IndexOption{Searchable}},
		{"appraisal", String, `archdesc[@level='collection']/appraisal/p`, []IndexOption{Searchable}},
		{"phystech", String, `archdesc[@level='collection']/phystech/p`, []IndexOption{Searchable}},

		// Find the following wherever they exist in the tree structure under <archdesc level="collection">
		// except under the inventory which starts at <dsc>
		{"chronlist", String, `archdesc[@level='collection']/*[name() != 'dsc']//chronlist/chronitem//text()`, []IndexOption{Searchable}},
		{"corpname", String, `archdesc[@level='collection']/*[name() != 'dsc']//corpname`, []IndexOption{Searchable, Displayable}},
		{"famname", String, `archdesc[@level='collection']/*[name() != 'dsc']//famname`, []IndexOption{Searchable, Displayable}},
		{"function", String, `archdesc[@level='collection']/*[name() != 'dsc']//function`, []IndexOption{Searchable, Displayable}},
		{"genreform", String, `archdesc[@level='collection']/*[name() != 'dsc']//genreform`, []IndexOption{Searchable, Displayable}},
		{"geogname", String, `archdesc[@level='collection']/*[name() != 'dsc']//geogname`, []IndexOption{Searchable, Displayable}},
		{"name", String, `archdesc[@level='collection']/*[name() != 'dsc']//name`, []IndexOption{Searchable, Displayable}},
		{"occupation", String, `archdesc[@level='collection']/*[name() != 'dsc']//occupation`, []IndexOption{Searchable, Displayable}},
		{"persname", String, `archdesc[@level='collection']/*[name() != 'dsc']//persname`, []IndexOption{Searchable, Displayable}},
		{"subject", String, `archdesc[@level='collection']/*[name() != 'dsc']//subject`, []IndexOption{Searchable, Displayable}},
		{"title", String, `archdesc[@level='collection']/*[name() != 'dsc']//title`, []IndexOption{Searchable, Displayable}},
		{"note", String, `archdesc[@level='collection']/*[name() != 'dsc']//note`, []IndexOption{Searchable, Displayable}},

		// Copy fields
		// collection is an alias for unittitle
		{"collection", String, `archdesc[@level='collection']/did/unittitle`, []IndexOption{Facetable, Displayable, Searchable}},
	},
}

// func creatorFieldsToXPath() string {
// 	str := ""
// 	for _, field := range []string{"corpname", "famname", "persname"} {
// 		str += "name() = '" + field + "' or "
// 	}
// 	return str
// }

// # Places to look for names
// NAME_FIELDS = [:corpname, :famname, :persname]

// module ClassMethods
//   def creator_fields_to_xpath
// 	@creator_fields_to_xpath ||= NAME_FIELDS.map {|field| "name() = '#{field}'"}.join(" or ")
//   end
// end
