#### TODO
```
[] use pointers to pass around large data structures  
[] add SolrDocAdd type  
[] add SolrDocDelete type  
[] add XPath to to Expression that adds default namespace to XPath string  
[] add error test cases for Gen
[] add func (sd *SolrDoc) AddField(v Any, t Term?)
[] break out types to separate file
[] add call to modify.Fabify() before generating XML files
[] figure out how XPaths are handled in ead_indexer gem: is everything converted to `//<xpath>` ?
[] what is up with `langcode`? There are no specified index options:     `t.langcode(path:"archdesc[@level='collection']/did/langmaterial/language/@langcode")`
[] need to handle HTML entity encoding, e.g., in `bioghist_teim`,
   expect:
    <field name="bioghist_teim">Nancy Wallace Henderson is an American playwright and author. Originally from North Carolina, Henderson majored in Dramatic Art at the University of North Carolina at Chapel Hill. There she studied playwriting under Walter Prichard Eaton and continued her studies in New York with Benjamin Bernard Zavin. Henderson won awards from the Dramatics Alliance in California and the University of Wisconsin, and was a member of the Author's Guild, the Dramatics Guild, and the International Women's Writing Guild.</field>

   got:
     <field name="bioghist_teim">Nancy Wallace Henderson is an American playwright and author. Originally from North Carolina, Henderson majored in Dramatic Art at the University of North Carolina at Chapel Hill. There she studied playwriting under Walter Prichard Eaton and continued her studies in New York with Benjamin Bernard Zavin. Henderson won awards from the Dramatics Alliance in California and the University of Wisconsin, and was a member of the Author&#39;s Guild, the Dramatics Guild, and the International Women&#39;s Writing Guild.</field>

[x] need to understand what this proxy does... just grabs the unittitle as collection?
    t.collection(proxy:[:unittitle],index_as:[:facetable,:displayable,:searchable])
    # it aliases "unittitle" as "collection" and indexes "collection" independently from "unittitle"

```