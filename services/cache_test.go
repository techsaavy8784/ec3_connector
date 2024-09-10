package services

import (
	//"context"

	fhir "github.com/dhf0820/fhir4"
	//common "github.com/dhf0820/uc_common"
	//"github.com/samply/golang-fhir-models/fhir-models/fhir"
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"

	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/token"
	"github.com/dhf0820/uc_core/util"

	common "github.com/dhf0820/uc_common"
	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//const pid = "Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB"

//const ordercode = "8310-5"
//const baseurl = "https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/"

//const pid = "4342009"
//const baseurl = "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/"

// func TestPatientCache(t *testing.T) {
// 	fmt.Printf("Test run a FHIR query")
// 	c := New(baseurl)
// 	Convey("Run a query", t, func() {

// 		caFhirId := "62f14531ba5395278cd530c4"
// 		patient, err := c.GetPatient("375", caFhirId)
// 		//bundle, err := c.PatientSearch(caFhirId, "family=smart")
// 		So(err, ShouldBeNil)
// 		So(patient, ShouldNotBeNil)
// 		// data, err := c.Query("Patient/12724066")
// 		// So(err, ShouldBeNil)
// 		// So(data, ShouldNotBeNil)
// 		pat, err := fhir4.UnmarshalPatient(bundle.Entry[0].Resource)
// 		So(err, ShouldBeNil)
// 		So(pat, ShouldNotBeNil)
// 		fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(pat))
// 	})
// }
func TestDocumentReferenceCacheBundle(t *testing.T) {
	fmt.Printf("Test run a FHIR query\n")
	c := New(baseurl)
	Convey("Run a query", t, func() {
		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
		//caFhirId := "62f14531ba5395278cd530c4"
		bundle, err := c.DocumentReferenceSearch("patient=12724066") //family=smart")
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		// data, err := c.Query("Patient/12724066")
		// So(err, ShouldBeNil)
		// So(data, ShouldNotBeNil)
		pat, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
		So(err, ShouldBeNil)
		So(pat, ShouldNotBeNil)
		//fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(pat))
	})
}

func TestDocumentReferenceCache(t *testing.T) {
	fmt.Printf("\n\n\n\ncacheTest:68  --  Test FHIR adding DocumentReference to query\n")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://localhost:30300/api/rest/v1")
		conf, err := GetServiceConfig("uc_fhir4", "local_test", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		//c := New(baseurl)
		err = os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!") //util.RandomString(32))
		So(err, ShouldBeNil)
		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
		So(err, ShouldBeNil)
		So(maker, ShouldNotBeNil)
		username := util.RandomOwner()
		duration := time.Minute
		//userId := "user123456"
		role := "Provider"
		ip := "192.168.1.1.99"
		fullName := "Debbie Harman MD"
		userId := "62d0af5dec383ade03a96b7e"
		//issuedAt := time.Now()
		//expiredAt := issuedAt.Add(duration)

		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		newToken = "Bearer " + newToken
		patientId := "12724066"
		//caFhirId := "62f14531ba5395278cd530c4"
		cerFhirId := "62f1c5dab3070d0b40e7aac1"
		fhirSystem, err := GetFhirSystem(cerFhirId)
		So(err, ShouldBeNil)
		So(fhirSystem, ShouldNotBeNil)
		cc := common.ConnectorConfig{}
		cp := common.ConnectorPayload{}
		cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
		cc.Name = "uc_cerner"
		cc.Version = "local_test"
		cc.CacheUrl = "http://uc_cache:9200"
		// "cacheurl" : "http://uc_cache:9200",
		// "cache_url" : "http://uc_cache:9200"
		data := []*common.KVData{}
		cacheServer := common.KVData{}
		cacheServer.Name = "cacheServer"
		cacheServer.Value = "http://192.168.1.117:30201"
		data = append(data, &cacheServer)
		hostServer := common.KVData{}
		hostServer.Name = "cacheHost"
		hostServer.Value = "http://ucCache:9200"
		data = append(data, &hostServer)
		cc.Data = data
		cp.FhirSystem = fhirSystem
		cp.ConnectorConfig = &cc
		//FindResource(fhirSystem *common.FhirSystem, url, resource, userId, query)
		cnt, bundle, hdr, err := FindResource(&cp, "DocumentReference", userId, "DocumentReference?patient="+patientId, newToken)

		// //data, err := c.Query("DocumentReference?patient=12724066")
		So(err, ShouldBeNil)
		So(cnt, ShouldNotEqual, 0)
		So(hdr, ShouldNotBeNil)
		So(bundle, ShouldNotBeNil)
		So(len(bundle.Entry), ShouldEqual, cnt)
		fmt.Printf("Count = %d\n", cnt)
		fmt.Printf("Header : %s\n", spew.Sdump(hdr))
		// // data, err = c.Query("DocumentReference?patient=12724066")
		// // So(err, ShouldBeNil)
		// // So(data, ShouldNotBeNil)
		// drBundle, err := fhir.UnmarshalBundle(data)
		// So(err, ShouldBeNil)
		// So(bundle, ShouldNotBeNil)
		// doc, err := fhir.UnmarshalDocumentReference(drBundle.Entry[0].Resource)
		// So(err, ShouldBeNil)
		// So(doc, ShouldNotBeNil)
		// docId := *doc.Id
		// fmt.Printf("Number of DocRefs = %d\n", len(bundle.Entry))
		// fmt.Printf("Id of first Document: %s\n", docId)
		// err = CacheResource(context.Background(), "queryId", userId, patientId, fhirSystem, &doc, "DocumentReference", docId)
		// So(err, ShouldBeNil)
	})
}

func TestDocumentReferenceBundleCache(t *testing.T) {
	fmt.Printf("\n\n\n\ncacheTest:68  --  Test FHIR ading DocumentReference to query\n")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
		conf, err := GetServiceConfig("uc_fhir4", "linode", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		c := New(baseurl)
		//caFhirId := "62f14531ba5395278cd530c4"
		cerFhirId := "62f1c5dab3070d0b40e7aac1"
		fhirSystem, err := GetFhirSystem(cerFhirId)
		So(err, ShouldBeNil)
		So(fhirSystem, ShouldNotBeNil)
		data, err := c.Query("DocumentReference?patient=12724066")
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		bundle, err := fhir.UnmarshalBundle(data)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)

		// doc, err := fhir4.UnmarshalDocumentReference(bundle.Entry[0].Resource)
		// So(err, ShouldBeNil)
		// So(doc, ShouldNotBeNil)
		//docId := *doc.Id
		// patientId := "12724066"
		// userId := "62d0af5dec383ade03a96b7e"
		// resourceType := "DocumentReference"
		fmt.Printf("TestDocumentReferenceBundleCache:137  --  Number of DocRefs = %d\n", len(bundle.Entry))
		//fmt.Printf("Id of first Document: %s\n", docId)
		//fmt.Printf("Document[0] returned: %s\n", spew.Sdump(doc))
		// CacheResourceBundleElements(context.Background(), userId,
		// 	patientId, fhirSystem, &bundle, resourceType)
	})
}
func TestGetCache(t *testing.T) {
	fmt.Printf("Test Get Cache for page")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
		conf, err := GetServiceConfig("uc_fhir4", "linode", "test") //GetConfig("delivery", "test")
		//conf, err := Initialize()
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		Convey("Run a query", func() {
			// pageId := 5
			// //queryId := "62ddb9f691f15a1e2d5206f7"
			// queryId := "62ddcd3261024840e7244591"
			// startTime := time.Now()
			// total, bundle, header, err := GetCache(queryId, pageId)
			// fmt.Printf("GetCache elapsed Time: %s\n", time.Since(startTime))
			// //
			// //fmt.Printf("Error: %s\n", err.Error())
			// So(err, ShouldBeNil)
			// So(header, ShouldNotBeNil)
			// So(bundle, ShouldNotBeNil)
			// So(total, ShouldNotEqual, 0)
			// fmt.Printf("GetCache Returned header returned: %s\n", spew.Sdump(header))
		})
	})
}

func TestGetDocumentReferenceCachePage(t *testing.T) {
	fmt.Printf("Test Get Cache for page")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
		conf, err := GetServiceConfig("uc_fhir4", "linode", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		Convey("Query for page 1  of num per page of 3", func() {
			pageId := int64(1)
			perPage := int64(1)                  //count
			userId := "62d0af5dec383ade03a96b7e" //e"
			docs, err := GetDocumentReferenceCachePage(userId, perPage, pageId)
			So(err, ShouldBeNil)
			So(docs, ShouldNotBeNil)
			fmt.Printf("GetCache Returned resource: %s\n", spew.Sdump(docs))
		})
	})
}

func TestGetObservationCachePage(t *testing.T) {
	fmt.Printf("Test Get Observation Cache for page")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
		conf, err := GetServiceConfig("uc_fhir4", "linode", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		Convey("Query for page 1  of num per page of 2", func() {
			pageId := int64(1)
			perPage := int64(2) //count
			userId := "62d0af5dec383ade03a96b7e"
			obs, err := GetObservationCachePage(userId, perPage, pageId)
			So(err, ShouldBeNil)
			So(obs, ShouldNotBeNil)
			fmt.Printf("GetObservationCache Returned resource: %s\n", spew.Sdump(obs))
		})
	})
}
