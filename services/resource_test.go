package services

import (
	//"encoding/json"

	//"context"

	//fhir "github.com/dhf0820/fhir4"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//fhir "github.com/dshills/gofhir"

	//log "github.com/sirupsen/logrus"
	"fmt"
	"os"
	"testing"

	"github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"
	"github.com/dhf0820/uc_core/util"

	"time"

	//"github.com/davecgh/go-spew/spew"
	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestResourceSearch(t *testing.T) {
	fmt.Printf("Test run a FHIR query")
	//c := New(baseurl)
	Convey("RunDocumentResourceQuery", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://localhost:30300/api/rest/v1")
		// This should not be used. It is now the ConnectorPayload
		// _, err := GetServiceConfig("uc_cerner", "local_test", "test") //GetConfig("delivery", "test")
		// So(err, ShouldBeNil)
		mongodb, err := OpenMongoDB()
		So(err, ShouldBeNil)
		So(mongodb, ShouldNotBeNil)
		userId := "62d0af5dec383ade03a96b7e"
		//userID, err := primitive.ObjectIDFromHex("62d0af5dec383ade03a96b7e")
		cerFhirId := "62f1c5dab3070d0b40e7aac1"
		//caFhirId := "62f14531ba5395278cd530c4"
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
		// Create limited sample ConnectorConfig should have just what the connector requires to retrieve and cache the Ressource

		err = os.Setenv("ACCESS_SECRET", util.RandomString(32))
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
		//issuedAt := time.Now()
		//expiredAt := issuedAt.Add(duration)

		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		//TODO:  Add user Id toAll handlers which will get it fom the header

		// Use id of known patient
		// _, bundle, _, err := SearchPatient(fhirSystem, "Patient", userId, "Patient?family=dawg&given=joel")
		// //_, bundle, _, err := SearchPatient(fhirSystem, "Patient", userId, "Patient?family=smart&given=na")
		// //bundle, err := c.PatientSearch(caFhirId, "family=smart&given=na&_count=12")
		// So(err, ShouldBeNil)
		// So(bundle, ShouldNotBeNil)
		// //fmt.Printf("bundle: %s\n", spew.Sdump(bundle.Entry[0]))
		// // pat := fhir.Patient{}
		// // json.Unmarshal(bundle.Entry[0], &pat)
		// // fmt.Printf("patient[0] = %s\n", spew.Sdump(pat))

		// //fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
		// // data, err := c.Query("Patient/12724066")
		// // So(err, ShouldBeNil)
		// // So(data, ShouldNotBeNil)

		// pat, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
		// So(err, ShouldBeNil)
		// So(pat, ShouldNotBeNil)

		//fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
		patientId := "12724066"
		fmt.Printf("\n\n\nTestResourceSearch:62 -- Patient.ID = %s\n\n\n\n", patientId)
		//fmt.Printf("Patient: - %s\n", spew.Sdump(pat))
		//time.Sleep(10 * time.Second)
		// //Get Documents for this patient

		qry := "/DocumentReference?patient=" + patientId
		fmt.Printf("\n\nTestResourceSearch:68  --  Process DocumentReference")
		cnt, bundle, hdr, err := FindResource(&cp, "DocumentReference", userId, qry, newToken)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		So(cnt, ShouldNotEqual, 0)
		So(hdr, ShouldNotBeNil)
		//fmt.Printf("TestResourceSearch:74  --  Number of entries = %d\n", len(bundle.Entry))
		//fmt.Printf("TestResourceSearch:75  --  Entry = %s\n", spew.Sdump(bundle.Entry))
		//rawDoc := bundle.Entry[0].Resource
		//doc, _ := fhir.UnmarshalDocumentReference(rawDoc)
		//fmt.Printf("TestResourceSearch:78  --  Doc: %s\n", spew.Sdump(doc))
		//fmt.Printf("Doc[0] = %s\n", spew.Sdump(doc))

		// func CachePatient(ctx context.Context, queryId string, userID primitive.ObjectID,
		// 	patientId string, fhirSystem *common.FhirSystem, resource *Interface,
		// 	resourceType string) error {

		//CacheResource(context.Background(), hdr.QueryId, userId, patientId, fhirSystem, &doc, "Documentreference", doc)

		qry = "/Observation?patient=" + patientId
		cnt, bundle, _, err = FindResource(&cp, "Observation", userId, qry, newToken)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		So(cnt, ShouldNotEqual, 0)
		// rawObs := bundle.Entry[0].Resource
		// obs, _ := fhir.UnmarshalObservation(rawObs)
		//fmt.Printf("Observation[0] = %s\n", spew.Sdump(obs))
		qry = fmt.Sprintf("/Condition?patient=%s", "12743119")
		//qry = fmt.Sprintf("/Condition?patient=%s", patientId)
		cnt, bundle, _, err = FindResource(&cp, "Condition", userId, qry, newToken)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		So(cnt, ShouldNotEqual, 0)

	})
}
