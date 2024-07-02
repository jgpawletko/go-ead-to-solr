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
	},
}

// # Descriptive information in <did>
// t.unitid(path:"archdesc[@level='collection']/did/unitid",index_as:[:searchable,:displayable])
// t.langcode(path:"archdesc[@level='collection']/did/langmaterial/language/@langcode")
// t.abstract(path:"archdesc[@level='collection']/did/abstract",index_as:[:searchable,:displayable])
// t.creator(path:"archdesc[@level='collection']/did/origination[@label='creator']/*[#{creator_fields_to_xpath}]",index_as:[:searchable,:displayable])

// # Dates
// t.unitdate_normal(path:"archdesc[@level='collection']/did/unitdate/@normal",index_as:[:displayable,:searchable,:facetable])
// t.unitdate(path:"archdesc[@level='collection']/did/unitdate[not(@type)]",index_as:[:searchable])
// t.unitdate_bulk(path:"archdesc[@level='collection']/did/unitdate[@type='bulk']",index_as:[:searchable])
// t.unitdate_inclusive(path:"archdesc[@level='collection']/did/unitdate[@type='inclusive']",index_as:[:searchable])

// # Fulltext in <p> under the following descriptive fields
// t.scopecontent(path:"archdesc[@level='collection']/scopecontent/p",index_as:[:searchable])
// t.bioghist(path:"archdesc[@level='collection']/bioghist/p",index_as:[:searchable])
// t.acqinfo(path:"archdesc[@level='collection']/acqinfo/p",index_as:[:searchable])
// t.custodhist(path:"archdesc[@level='collection']/custodhist/p",index_as:[:searchable])
// t.appraisal(path:"archdesc[@level='collection']/appraisal/p",index_as:[:searchable])
// t.phystech(path:"archdesc[@level='collection']/phystech/p",index_as:[:searchable])

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
