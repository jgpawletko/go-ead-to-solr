## Tech Notes

**Strategy:** `code spike`

**Thinking:** `first make it run, then make it fast`

**Discussion:**  
Given that we already have the files containing the XPath expressions, e.g., `document.rb`, then we could just implement the existing functionality to Go and see if the resulting code was suffciently performant.


** code flow:** 
```
input  --> path to EAD file
output --> Solr XML files

parameters:
-p <path to EAD file>
-o <path to output directory> [DEFAULT .]

code execution:
  run standard file-existence and readability tests
  open EAD file
  for each Term in the Terminology
    for each element returned by the XPath query
      for each IndexOption
        load the field into SolrDoc
      end
    end
  end 
end

write SolrDocAdd XML document
```

----
**References:**  

* Adding XML Header to output:  `<?xml version="1.0" encoding="UTF-8"?>`
  https://stackoverflow.com/a/64250018

* What do the Solr `<field @name>` suffixes mean?    
  https://github.com/samvera/active_fedora/blob/12.0-stable/lib/active_fedora/indexing/default_descriptors.rb  

* Where do the `XPath`s and stuff come from?  
  https://github.com/NYULibraries/ead_indexer/blob/master/lib/ead_indexer/document.rb

* What suffixes are used in the Solr schema?  
  https://github.com/NYULibraries/specialcollections/blob/master/solr/conf/schema.xml#L593-L748

* What `IndexOptions` are actually used by the `ead_indexer` Gem?
  ```
  # path to root of ead_indexer Gem
  # e.g., 
  [~/dev/.../ead_indexer](master)$ pwd
  /Users/$USER/dev/nyulibraries/ead_indexer

  $ egrep -lr 'index_as|insert_field' .
  ./lib/ead_indexer/component.rb
  ./lib/ead_indexer/document.rb
  ./lib/ead_indexer/behaviors/dates.rb
 
  $ grep -r index_as . | rev | cut -d'[' -f1|rev| cut -d] -f1 | tr , $'\n' |sort | uniq
  :displayable
  :facetable
  :searchable

  $ grep -r insert_field . | cut -d, -f4-| tr -d '[:blank:]'| cut -d\) -f1 | tr , $'\n' |sort | uniq
  :displayable
  :facetable
  :searchable
  :sortable
  :stored_sortable
  ```

  **Bottom line:**  
  At a minimum, we must support the following `IndexOptions`:
  ```
  :displayable
  :facetable
  :searchable
  :sortable
  :stored_sortable
  ```
