package services

import (
// "github.com/dhf0820/fhir4"
// //log "github.com/sirupsen/logrus"
// //. "github.com/smartystreets/goconvey/convey"

// "fmt"
// //"os"
// "context"
// "testing"
// //"time"

// //"github.com/davecgh/go-spew/spew"
// //log "github.com/sirupsen/logrus"
// . "github.com/smartystreets/goconvey/convey"
)

// func TestDocumentReferenceCacheBundle(t *testing.T) {
// 	fmt.Printf("\n\n\n\n\nDocumentReferenceTest:19  --  Test Adding a documentReference to BundleCache")
// 	c := New(baseurl)
// 	Convey("RunDocumentReferenceQuery", t, func() {
// 		// bundle, err := c.DocumentReferenceSearch("patient=12724066&_count=5")
// 		// So(err, ShouldBeNil)
// 		// So(bundle, ShouldNotBeNil)
// 		// fmt.Printf("DocumentReferenceSearch returned: %s\n", spew.Sdump(bundle))
// 		c.Query("DocumentReference?Patient=12724066")
// 		cerFhirId := "62f1c5dab3070d0b40e7aac1"
// 		//caFhirId := "62f14531ba5395278cd530c4"
// 		fhirSystem, err := GetFhirSystem(cerFhirId)
// 		userId := "62d0af5dec383ade03a96b7e"
// 		data, err := c.Query("DocumentReference?patient=12724066")
// 		So(err, ShouldBeNil)
// 		So(data, ShouldNotBeNil)
// 		bundle, err := fhir4.UnmarshalBundle(data)
// 		So(err, ShouldBeNil)
// 		So(bundle, ShouldNotBeNil)
// 		doc, err := fhir4.UnmarshalDocumentReference(bundle.Entry[0].Resource)
// 		So(err, ShouldBeNil)
// 		So(doc, ShouldNotBeNil)
// 		// CacheResource(ctx context.Context, queryId, userId,
// 		// 	patientId string, fhirSystem *common.FhirSystem, resource Interface,
// 		// 	resourceType string
// 		docId := doc.Id
// 		err = CacheResource(context.Background(), "queryId", userId, "patientId", fhirSystem, &doc, "DocumentReference", *docId)
// 		So(err, ShouldBeNil)
// 		//fmt.Printf("DocumentReferenceSearch returned: %s\n", spew.Sdump(doc))
// 		fmt.Printf("Patient:= %s\n", *doc.Subject.Reference)
// 	})
// }
