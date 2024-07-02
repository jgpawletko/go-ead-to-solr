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

```