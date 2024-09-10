package services

import (
	"bytes"
	"encoding/json"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"
	//"github.com/dhf0820/uc_core/service"

	"fmt"
	//fhir "github.com/dhf0820/fhirR4go"
	"log"
	"net/http"
	"net/http/httptest"

	fhir "github.com/dhf0820/fhir4"
	"github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"
	"github.com/dhf0820/uc_core/util"

	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchPatient(t *testing.T) {
	Convey("Subject: SearchForPatient", t, func() {
		godotenv.Load("./.env.cerner_test")
		// os.Setenv("CONFIG_ADDRESS", "http://192.168.1.117:30300/api/rest/v1")
		// //os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")
		// os.Setenv("COMPANY", "test")

		// conf, err := service.InitCore("uc_core", "go_test", "test")
		// log.Printf("\n\n\nInitCore returned\n\n\n")
		// if err != nil {
		// 	log.Printf("Err = %s\n\n", err.Error())
		// }
		// if conf == nil {
		// 	log.Println("Conf is nil")
		// }
		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)

		// _, err := service.Initialize()
		// //, err := service.InitCore("uc_core","test", "test")
		// if err != nil {
		// 	t.Fatalf("InitCore failed: %s", err.Error())
		// }

		///   Handle query for patient, create a cache header info for the queryquerying for results cache each element of the bundle
		//in the back ground check on the size of the results aver couple of seconds.  When the number _count requested, Return them
		// in a new standart cache results

		log.Printf("\n\n\ntesting TestSearchFhirForPatient\n\n\n")
		Convey("Given a valid family/given name", func() {
			log.Printf("\n\nGiven a valid family/given name\n\n\n")
			w := httptest.NewRecorder()
			fmt.Printf("Creating Search Request\n")
			req, _ := http.NewRequest("GET", "/634f0ec03240a53a52a83a9d/Patient?family=smart&given=fred&_count=2", nil)
			//req, _ := http.NewRequest("GET", "/api/rest/v1/Patient?family=smart&given=fred&_count=2", nil)
			godotenv.Load("./.env.core_test")
			err := os.Setenv("ACCESS_SECRET", util.RandomString(32))
			So(err, ShouldBeNil)
			maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
			So(err, ShouldBeNil)
			So(maker, ShouldNotBeNil)
			username := util.RandomOwner()
			duration := time.Minute
			userId := "user123456"
			role := "Provider"
			ip := "192.168.1.1.99"
			fullName := "Debbie Harman MD"
			//issuedAt := time.Now()
			//expiredAt := issuedAt.Add(duration)

			newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
			So(err, ShouldBeNil)
			So(newToken, ShouldNotBeNil)
			So(payload, ShouldNotBeNil)

			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
			//req.Header.Add("tracing-id", "123")
			// vars := map[string]string{
			// 	"doc_id" : "40441",
			// }
			//req = mux.SetURLVars(req, vars)
			req.Header.Set("facility", "demo")
			req.Header.Set("FhirVersion", "r4")
			req.Header.Set("FhirSystemId", "634f0ec03240a53a52a83a9d")
			req.Header.Set("AUTHORIZATION", newToken)
			fmt.Printf("")

			getPatient(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			var bundle []fhir.Bundle
			fmt.Printf("TestSearchPatient:105  --  Resp = %s\n", spew.Sdump(resp))
			err = json.NewDecoder(resp.Body).Decode(&bundle)
			So(err, ShouldBeNil)
			fmt.Printf("bunde: %s\n", spew.Sdump(bundle))

		})
	})
}

func TestPatientGet(t *testing.T) {
	Convey("Subject: GetPatient", t, func() {
		//godotenv.Load("./.env.core_test")
		// os.Setenv("CONFIG_ADDRESS", "http://192.168.1.117:30300/api/rest/v1")
		// //os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")
		// os.Setenv("COMPANY", "test")

		// conf, err := service.InitCore("uc_core", "go_test", "test")
		// log.Printf("\n\n\nInitCore returned\n\n\n")
		// if err != nil {
		// 	log.Printf("Err = %s\n\n", err.Error())
		// }
		// if conf == nil {
		// 	log.Println("Conf is nil")
		// }
		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)

		// _, err := service.Initialize()
		// //, err := service.InitCore("uc_core","test", "test")
		// if err != nil {
		// 	t.Fatalf("InitCore failed: %s", err.Error())
		// }

		///   Handle query for patient, create a cache header info for the queryquerying for results cache each element of the bundle
		//in the back ground check on the size of the results aver couple of seconds.  When the number _count requested, Return them
		// in a new standart cache results

		log.Printf("\n\n\ntesting TestGetPatient\n\n\n")
		Convey("Given a valid patient ID", func() {
			log.Printf("\n\nGiven a valid PatientId\n\n\n")
			w := httptest.NewRecorder()
			fmt.Printf("Creating Get Request\n")
			//req := httptest.NewRequest("GET", "/62f14531ba5395278cd530c4/Patient/12724066", nil)
			//req, _ := http.NewRequest("GET", "/api/rest/v1/Patient?family=smart&given=fred&_count=2", nil)
			godotenv.Load("./.env.core_test")
			err := os.Setenv("ACCESS_SECRET", util.RandomString(32))
			So(err, ShouldBeNil)
			maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
			So(err, ShouldBeNil)
			So(maker, ShouldNotBeNil)
			username := util.RandomOwner()
			duration := time.Minute
			userId := "user123456"
			role := "Provider"
			ip := "192.168.1.1.99"
			fullName := "Debbie Harman MD"
			//issuedAt := time.Now()
			//expiredAt := issuedAt.Add(duration)

			newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
			So(err, ShouldBeNil)
			So(newToken, ShouldNotBeNil)
			So(payload, ShouldNotBeNil)

			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
			//req.Header.Add("tracing-id", "123")
			// vars := map[string]string{
			// 	"doc_id" : "40441",
			// }
			//req = mux.SetURLVars(req, vars)
			//fmt.Printf("testGetPatient:180  --  req: %s\n", spew.Sdump(req))
			connectorPayload := common.ConnectorPayload{}
			//fs, err := GetFhirSystem("62f14531ba5395278cd530c4")
			//So(err, ShouldBeNil)
			//So(fs, ShouldNotBeNil)
			fs := common.FhirSystem{}
			fs.FhirUrl = "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d"
			fs.UcUrl = "http://test.universalcharts.com/634f0ec03240a53a52a83a9d"
			fs.FacilityName = "Mercy Redding"
			fs.DisplayName = "CernerOpen"
			fs.FacilityCode = "demo"
			fs.FhirVersion = "r4"
			fs.Identifiers = []*common.KVData{}
			ident := common.KVData{}
			ident.Name = "mrn"
			ident.Value = "urn:oid:2.16.840.1.113883.6.1000|"
			fs.Identifiers = append(fs.Identifiers, &ident)
			fs.ID, err = primitive.ObjectIDFromHex("634f0ec03240a53a52a83a9d")
			So(err, ShouldBeNil)
			So(fs.ID, ShouldNotBeNil)
			fs.FacilityId, err = primitive.ObjectIDFromHex("634f0ec03240a53a52a83a9d")
			fs.Insert = "false"
			So(err, ShouldBeNil)
			So(fs.ID, ShouldNotBeNil)

			connectorPayload.FhirSystem = &fs
			requestBody, err := json.Marshal(connectorPayload)
			So(err, ShouldBeNil)
			req := httptest.NewRequest("GET", "/634f0ec03240a53a52a83a9d/Patient/12743120", bytes.NewBuffer(requestBody))
			fmt.Printf("Setting Headers\n")
			// req.Header.Set("facility", "demo")
			// req.Header.Set("FhirVersion", "r4")
			// req.Header.Set("FhirSystemId", "62f14531ba5395278cd530c4")
			// req.Header.Set("AUTHORIZATION", newToken)

			h := map[string][]string{
				"facility":      {"demo"},
				"FhirVersion":   {"r4"},
				"FhirSystemId":  {"634f0ec03240a53a52a83a9d"},
				"Authorization": {newToken},
			}
			req.Header = h
			if req == nil {
				fmt.Printf("req is NIL\n")
			}
			if w == nil {
				fmt.Printf("w is nil\n")
			}

			fmt.Printf("testGetPatient:202  --  req: %s\n", spew.Sdump(req))

			getPatient(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			//var bundle []fhir.Bundle
			var patient common.ResourceResponse
			//fmt.Printf("testGetPatient:229  --  Resp = %s\n", spew.Sdump(resp))
			//err = json.NewDecoder(resp.Body).Decode(&bundle)
			err = json.NewDecoder(resp.Body).Decode(&patient)
			So(err, ShouldBeNil)
			fmt.Printf("getPatient:234  --  Patient: %s\n", spew.Sdump(patient))

		})
	})
}
