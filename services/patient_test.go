package services

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/fhir4"
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"

	"fmt"
	"os"
	"testing"

	"github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/uc_core/util"
	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func TestPostPatient(t *testing.T) {
// 	fmt.Printf("Test Add a new patient to Server")
// 	c := New(baseurl)
// 	Convey("PostNewPatient", t, func() {
// 		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")
// 		cfg, err := GetServiceConfig("uc_fhir4", "local_test", "test") //GetConfig("delivery", "test")
// 		fmt.Printf("cfg = %s\n", spew.Sdump(cfg))
// 		So(err, ShouldBeNil)
// 		mongodb, err := OpenMongoDB()
// 		So(err, ShouldBeNil)
// 		So(mongodb, ShouldNotBeNil)
// 		//conf, err := Initialize()
// 		So(err, ShouldBeNil)
// 		//caFhirId := "62d0ad3c9d0119afff9978b3"
// 		cerFhirId := "62f1c5dab3070d0b40e7aac1"
// 		err = os.Setenv("ACCESS_SECRET", util.RandomString(32))
// 		So(err, ShouldBeNil)
// 		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
// 		So(err, ShouldBeNil)
// 		So(maker, ShouldNotBeNil)
// 		username := util.RandomOwner()
// 		duration := time.Minute
// 		//userId := "user123456"
// 		userId := "62d0af5dec383ade03a96b7e"
// 		role := "Provider"
// 		ip := "192.168.1.1.99"
// 		fullName := "Debbie Harman MD"
// 		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
// 		So(err, ShouldBeNil)
// 		So(newToken, ShouldNotBeNil)
// 		So(payload, ShouldNotBeNil)
// 		fhirSystem, err := GetFhirSystem(cerFhirId)
// 		So(err, ShouldBeNil)
// 		cc := common.ConnectorConfig{}
// 		cp := common.ConnectorPayload{}
// 		cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
// 		cc.Name = "uc_cerner"
// 		cc.Version = "local_test"
// 		cc.CacheUrl = "http://uc_cache:9200"
// 		// "cacheurl" : "http://uc_cache:9200",
// 		// "cache_url" : "http://uc_cache:9200"
// 		data := []*common.KVData{}
// 		cacheServer := common.KVData{}
// 		cacheServer.Name = "cacheServer"
// 		cacheServer.Value = "http://192.168.1.117:30201"
// 		data = append(data, &cacheServer)
// 		hostServer := common.KVData{}
// 		hostServer.Name = "cacheHost"
// 		hostServer.Value = "http://ucCache:9200"
// 		data = append(data, &hostServer)
// 		cc.Data = data
// 		cp.FhirSystem = fhirSystem
// 		cp.ConnectorConfig = &cc
// 		dlhFhirId := "6329112852f3616990e2f763"
// 		dlhFhirSystem, err := GetFhirSystem(dlhFhirId)
// 		So(err, ShouldBeNil)
// 		So(dlhFhirSystem, ShouldNotBeNil)
// 		fmt.Printf("dlhFhirSystemURL %s\n", dlhFhirSystem.FhirUrl)
// 		//Get a Cerner Patient to save
// 		cnt, bundle, header, err := FindResource(&cp, "Patient", userId, "family=smart&_count=12", newToken)
// 		//bundle, err := c.PatientSearch(fhirSystem, "family=smart&given=sally&_count=12", "Patient", newToken)
// 		So(err, ShouldBeNil)
// 		So(bundle, ShouldNotBeNil)
// 		So(header, ShouldNotBeNil)
// 		So(cnt, ShouldNotEqual, 0)
// 		fmt.Printf("TestPostPatient:67 returned %d resources\n", cnt)
// 		//fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
// 		// data, err := c.Query("Patient/12724066")
// 		// So(err, ShouldBeNil)
// 		// So(data, ShouldNotBeNil)

// 		pat, err := fhir4.UnmarshalPatient(bundle.Entry[1].Resource)
// 		//fmt.Printf("PATIENT: %s\n", spew.Sdump(pat))
// 		So(err, ShouldBeNil)
// 		So(pat, ShouldNotBeNil)
// 		fmt.Printf("Number of entries: %d\n", len(bundle.Entry))

// 		//fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
// 		fmt.Printf("Patient.ID := %s\n", *pat.Id)
// 		pat.Id = nil
// 		pat.Meta = nil
// 		newPat, err := c.PostPatient(dlhFhirSystem.FhirUrl, "210205", &pat)
// 		So(err, ShouldBeNil)
// 		So(newPat, ShouldNotBeNil)
// 		fmt.Printf("NewPatient: %s\n", spew.Sdump(newPat))
// 		//time.Sleep(15 * time.Second)
// 	})
// }

func TestPatientSearch(t *testing.T) {
	fmt.Printf("Test run a FHIR query")
	//c := New(baseurl)
	Convey("RunPatientQuery", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
		_, err := GetServiceConfig("uc_fhir4", "local_test", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		mongodb, err := OpenMongoDB()
		So(err, ShouldBeNil)
		So(mongodb, ShouldNotBeNil)
		//conf, err := Initialize()
		So(err, ShouldBeNil)
		//caFhirId := "62d0ad3c9d0119afff9978b3"
		cerFhirId := "62f1c5dab3070d0b40e7aac1"
		fhirSystem, err := GetFhirSystem(cerFhirId)
		So(err, ShouldBeNil)
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
		//err = os.Setenv("ACCESS_SECRET", util.RandomString(32))
		err = os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
		So(err, ShouldBeNil)
		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
		So(err, ShouldBeNil)
		So(maker, ShouldNotBeNil)
		username := util.RandomOwner()
		duration := 10 * time.Minute
		//userId := "user123456"
		userId := "62d0af5dec383ade03a96b7e"
		role := "Provider"
		ip := "192.168.1.1.99"
		fullName := "Debbie Harman MD"
		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		cnt, bundle, header, err := FindResource(&cp, "Patient", userId, "family=smart&_count=12", newToken)
		//bundle, err := c.PatientSearch(fhirSystem, "family=smart&_count=12", "patient", newToken)
		So(header, ShouldNotBeNil)
		So(cnt, ShouldNotEqual, 0)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		fmt.Printf("TestPatientSearch:130 returned %d resources\n", cnt)
		//fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
		// data, err := c.Query("Patient/12724066")
		// So(err, ShouldBeNil)
		// So(data, ShouldNotBeNil)
		pat, err := fhir4.UnmarshalPatient(bundle.Entry[0].Resource)
		So(err, ShouldBeNil)
		So(pat, ShouldNotBeNil)

		//fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
		fmt.Printf("Patient.ID := %s\n", *pat.Id)
		time.Sleep(15 * time.Second)
	})
}

func TestCaPatientSearch(t *testing.T) {
	fmt.Printf("Test query through to CA for patient")
	//c := New(baseurl)
	Convey("RunCaPatientQuery", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
		_, err := GetServiceConfig("uc_fhir4", "local_test", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		mongodb, err := OpenMongoDB()
		So(err, ShouldBeNil)
		So(mongodb, ShouldNotBeNil)
		//conf, err := Initialize()
		So(err, ShouldBeNil)
		caFhirId := "62f14531ba5395278cd530c4"
		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
		fhirSystem, err := GetFhirSystem(caFhirId)
		So(err, ShouldBeNil)
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
		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
		So(err, ShouldBeNil)
		So(maker, ShouldNotBeNil)
		username := util.RandomOwner()
		duration := time.Minute
		//userId := "user123456"
		userId := "62d0af5dec383ade03a96b7e"
		role := "Provider"
		ip := "192.168.1.1.99"
		fullName := "Debbie Harman MD"
		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		// bundle, err := c.PatientSearch(fhirSystem, "family=HARMAN", "Patient", newToken)
		// So(err, ShouldBeNil)
		// So(bundle, ShouldNotBeNil)
		cnt, bundle, header, err := FindResource(&cp, "Patient", userId, "family=smart&_count=12", newToken)
		//bundle, err := c.PatientSearch(fhirSystem, "family=smart&_count=12", "patient", newToken)
		So(header, ShouldNotBeNil)
		So(cnt, ShouldNotEqual, 0)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		fmt.Printf("TestCaPatientSearch:184 returned %d resources\n", cnt)
		//fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
		// data, err := c.Query("Patient/12724066")
		// So(err, ShouldBeNil)
		// So(data, ShouldNotBeNil)
		pat, err := fhir4.UnmarshalPatient(bundle.Entry[0].Resource)
		So(err, ShouldBeNil)
		So(pat, ShouldNotBeNil)

		//fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
		fmt.Printf("Patient.ID = %s Name = %s\n", *pat.Id, *pat.Name[0].Family)
		//ime.Sleep(15 * time.Second)
	})
}

func TestGetPatient(t *testing.T) {
	fmt.Printf("Test run a FHIR query")
	//c := New(baseurl)
	Convey("RunPatientGet", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30300/api/rest/v1")
		_, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		mongodb, err := OpenMongoDB()
		So(err, ShouldBeNil)
		So(mongodb, ShouldNotBeNil)
		//conf, err := Initialize()
		So(err, ShouldBeNil)
		pat, err := GetPatient("12743944")
		So(err, ShouldBeNil)
		So(pat, ShouldNotBeNil)
		//fmt.Printf("Found Patient: %s\n", spew.Sdump(pat))
	})
}

func TestCreateIdentifier(t *testing.T) {
	fmt.Printf("Test run a FHIR query")
	//c := New(baseurl)
	Convey("CreateIdentifier", t, func() {
		// os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30300/api/rest/v1")
		// _, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
		// So(err, ShouldBeNil)
		// mongodb, err := OpenMongoDB()
		// So(err, ShouldBeNil)
		// So(mongodb, ShouldNotBeNil)
		// //conf, err := Initialize()
		// So(err, ShouldBeNil)
		ident := CreateIdentifier("63c703eac5cc538807e9b775")
		So(ident, ShouldNotBeNil)
		fmt.Printf("Ident: %s\n", spew.Sdump(ident))
	})
}
