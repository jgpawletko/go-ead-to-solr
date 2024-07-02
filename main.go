package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	ead "github.com/jgpawletko/go-ead-to-solr/ead"
	"github.com/nyulibraries/dlts-finding-aids-ead-go-packages/ead/modify"
)

func assertFile(filePath string) error {

	filePathInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if os.FileInfo.IsDir(filePathInfo) {
		return fmt.Errorf("%s is not a file", filePath)
	}

	return nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("error: incorrect number of EAD file arguments")
		os.Exit(1)
	}

	eadFile := os.Args[1]
	err := assertFile(eadFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	eadFile = strings.TrimSpace(eadFile)
	eadFilePath, _ := filepath.Abs(eadFile)

	EADXML, err := os.ReadFile(eadFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fabifiedEAD, errors := modify.FABifyEAD(EADXML)
	if len(errors) != 0 {
		for _, eMsg := range errors {
			fmt.Printf("%s\n", eMsg)
		}
		os.Exit(1)
	}

	solrDoc, errors := ead.GenSolrDoc([]byte(fabifiedEAD), ead.EADTerminology)
	if len(errors) != 0 {
		for _, eMsg := range errors {
			fmt.Printf("%s\n", eMsg)
		}
		os.Exit(1)
	}

	solrAddDoc := ead.SolrAdd{SolrDoc: solrDoc}

	output, err := xml.MarshalIndent(solrAddDoc, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
	os.Exit(0)
}
