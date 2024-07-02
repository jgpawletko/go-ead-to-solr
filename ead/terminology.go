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

		// # Find the following wherever they exist in the tree structure under <archdesc level="collection">
		// # except under the inventory which starts at <dsc>
		// t.chronlist(path:"archdesc[@level='collection']/*[name() != 'dsc']//chronlist/chronitem//text()",index_as:[:searchable])
		// t.corpname(path:"archdesc[@level='collection']/*[name() != 'dsc']//corpname",index_as:[:searchable,:displayable])
		// t.famname(path:"archdesc[@level='collection']/*[name() != 'dsc']//famname",index_as:[:searchable,:displayable])
		// t.function(path:"archdesc[@level='collection']/*[name() != 'dsc']//function",index_as:[:searchable,:displayable])
		// t.genreform(path:"archdesc[@level='collection']/*[name() != 'dsc']//genreform",index_as:[:searchable,:displayable])
		// t.geogname(path:"archdesc[@level='collection']/*[name() != 'dsc']//geogname",index_as:[:searchable,:displayable])
		// t.name(path:"archdesc[@level='collection']/*[name() != 'dsc']//name",index_as:[:searchable,:displayable])
		// t.occupation(path:"archdesc[@level='collection']/*[name() != 'dsc']//occupation",index_as:[:searchable,:displayable])
		// t.persname(path:"archdesc[@level='collection']/*[name() != 'dsc']//persname",index_as:[:searchable,:displayable])
		// t.subject(path:"archdesc[@level='collection']/*[name() != 'dsc']//subject",index_as:[:searchable,:displayable])
		// t.title(path:"archdesc[@level='collection']/*[name() != 'dsc']//title",index_as:[:searchable,:displayable])
		// t.note(path:"archdesc[@level='collection']/*[name() != 'dsc']//note",index_as:[:searchable,:displayable])

	},
}

func creatorFieldsToXPath() string {
	str := ""
	for _, field := range []string{"corpname", "famname", "persname"} {
		str += "name() = '" + field + "' or "
	}
	return str
}

// # Places to look for names
// NAME_FIELDS = [:corpname, :famname, :persname]

// module ClassMethods
//   def creator_fields_to_xpath
// 	@creator_fields_to_xpath ||= NAME_FIELDS.map {|field| "name() = '#{field}'"}.join(" or ")
//   end
// end
