package ead

var EADTerminology = Terminology{
	Root: "ead",
	Terms: []Term{
		{"author", String, `filedesc/titlestmt/author`, []IndexOption{Searchable, Displayable}},
	},
}
