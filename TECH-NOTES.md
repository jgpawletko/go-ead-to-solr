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

* What do the Solr `<field @name>` suffixes mean?    
  https://github.com/samvera/active_fedora/blob/12.0-stable/lib/active_fedora/indexing/default_descriptors.rb  

* Where do the `XPath`s and stuff come from?  
  https://github.com/NYULibraries/ead_indexer/blob/master/lib/ead_indexer/document.rb

