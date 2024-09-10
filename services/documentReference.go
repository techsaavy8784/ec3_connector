package services

import (
	"github.com/dhf0820/fhir4"
	//"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"time"
)

func (c *Connection) DocumentReferenceSearch(query string) (*fhir4.Bundle, error) {
	fmt.Printf("\n\n\nDocumentReferenceSearch:13  --  started\n\n")
	log.Infof("queryString: %s\n", query)
	qry := fmt.Sprintf("DocumentReference?%s", query)
	log.Infof("Final url for DocumentReference query: %s\n", qry)
	startTime := time.Now()
	bytes, err := c.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("Query %s failed: %s", query, err.Error())
	}
	bundle, err := fhir4.UnmarshalBundle(bytes)
	log.Infof("Search DocumentReference returned %d documents in: %s", len(bundle.Entry), time.Since(startTime))
	if err != nil {
		return nil, err
	}
	///entry, err := fhir4.UnmarshalDocumentReference(e1.Resource)
	e1 := bundle.Entry[0]

	entry, err := fhir4.UnmarshalDocumentReference(e1.Resource)

	fmt.Printf("DocumentReferenceSearch:25 - Entry[0] = %s  ID = %s\n", spew.Sdump(entry), *entry.Id)
	return &bundle, nil
}
