package services

import (
	//"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	//"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"
	"github.com/dhf0820/uc_core/service"

	"github.com/davecgh/go-spew/spew"
	fhir4 "github.com/dhf0820/fhir4"
	//fhirR4go "github.com/dhf0820/fhirR4go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleFindResource(t *testing.T) {
	Convey("TestSimpleFindResource", t, func() {
		req, err := http.NewRequest("GET", "/api/rest/v1/FindPatient?family=smart", nil)
		//req, err := http.NewRequest("GET", "/api/rest/v1/GetPatient/12345", nil)
		So(err, ShouldBeNil)
		w := httptest.NewRecorder()
		os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
		dur := time.Duration(300) * time.Second
		jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
		So(err, ShouldBeNil)
		So(jwt, ShouldNotBeNil)
		req.Header.Set("Authorization", jwt)
		fmt.Printf("\nCalling Router\n")
		NewRouter().ServeHTTP(w, req)
	})
}

func TestFindResource(t *testing.T) {
	Convey("TestFindResource", t, func() {
		//req, err := http.NewRequest("GET", "/api/rest/v1/healthcheck", nil) //"api/Patient?family=smart", nil)
		// req, err := http.NewRequest("GET", "/api/rest/v1/Patient?family=smart", nil)
		// So(err, ShouldBeNil)
		//req, err := http.NewRequest("GET", "/api/rest/v1/Patient/12345678", nil)
		// vars := map[string]string{
		// 	"resource": "Patient",
		// }
		// req = mux.SetURLVars(req, vars)
		fmt.Printf("\n\nSetting FhirSystem body\n")
		fs := `{
				"id": "62f1c5dab3070d0b40e7aac1",
				"facilityId": "62e89b57e2da183de83c27a2",
				"facilityName": "Mercy Redding",
				"displayName": "Cerner Open",
				"description": "Medical Records from 2013-Current",
				"fhirVersion": "r4",
				"authUrl": "",
				"ucUrl": "http://192.168.1.117:30300/62f1c5dab3070d0b40e7aac1",
				"fhirUrl": "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d",
				"insert": "false",
				"identifiers": [
					{
						"Name": "Mrn",
						"Value": "urn:oid:2.16.840.1.113883.6.1000|"
					},
					{
						"Name": "Community",
						"Value": "urn:oid:2.16.840.1.113883.3.787.0.0|"
					},
					{
						"Name": "Military",
						"Value": "urn:oid:2.16.840.1.113883.3.42.10001.100001.12|"
					},
					{
						"Name": "Millennium",
						"Value": "https://fhir.cerner.com/ec2458f2-1e24-41c8-b71b-0e701af7583d/codeSet/4|"
					}
				],
				"facilityCode": "",
				"serviceName" : "uc_cerner",
				"returnBundle" : "true",
				"connector" : "uc_cerner:local_test"
			}`
		//}`

		fsByte := []byte(fs)
		fhirSystem := &common.FhirSystem{}
		err := json.Unmarshal(fsByte, fhirSystem)
		So(err, ShouldBeNil)
		fmt.Printf("\n\nTest:102  --  fhirSystem = %s\n", spew.Sdump(fhirSystem))
		// fsBody := json.RawMessage(fs)
		// //(err, ShouldBeNil)
		// fmt.Printf("fsBody RawMessage String: %s\n\n", fsBody)
		// fst := common.FhirSystem{}
		// json.Unmarshal(fsBody, &fst)
		//So(err, ShouldBeNil)
		//fmt.Printf("\n\ntest:106 -- FhirSystem = %s\n", spew.Sdump(fhirSystem))
		cp := common.ConnectorPayload{}
		cp.FhirSystem = fhirSystem
		//cp := common.ConnectorPayload{}
		cc := common.ConnectorConfig{}
		cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
		cc.Name = "uc_cerner"
		cc.Version = "local_test"
		cc.CacheUrl = "http://localhost:30201"
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
		cps, err := json.Marshal(cp)
		rc := io.NopCloser(strings.NewReader(string(cps)))
		//req.Body, err = json.Marshal([]byte(fs))
		//req.Body = rc
		//req, err := http.NewRequest("GET", "/api/rest/v1/Find/Patient?family=smart", rc)
		req, err := http.NewRequest("GET", "/api/rest/v1/Find/DocumentReference?patient=12724066", rc)
		So(err, ShouldBeNil)
		// values := req.URL.Query()
		// values.Add("resource", "Patient")
		// req.URL.RawQuery = values.Encode()
		// vars := map[string]string{
		// 	"resource": "Patient",
		// }
		//req = mux.SetURLVars(req, vars)
		// fmt.Printf("test request: ")
		// spew.Dump(req)
		//So(err, ShouldBeNil)
		w := httptest.NewRecorder()
		os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
		dur := time.Duration(300) * time.Second
		jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
		So(err, ShouldBeNil)
		So(jwt, ShouldNotBeNil)
		req.Header.Set("Authorization", jwt)

		fmt.Printf("\nCalling Router\n")
		StartTime := time.Now()
		NewRouter().ServeHTTP(w, req)
		fmt.Printf("\n\n\n\n\nResults: %s\n", w.Result().Status)
		fmt.Printf("########################## Wait for Background Elapsed Time = %s\n\n\n\n\n", time.Since(StartTime))
		time.Sleep(20 * time.Second)
	})
}
func TestFindCernerResource(t *testing.T) {
	Convey("Subject: Find all Resources matching Filter", t, func() {
		godotenv.Load("./.env.uc_fhir4_test")
		//os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")

		//os.Setenv("COMPANY", "demo")

		// conf, err := service.InitCore("uc_fhir4", "local_test", "test")
		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)
		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)
		// _, err := service.Initialize()
		// //, err := service.InitCore("uc_core","test", "test")
		// if err != nil {
		// 	t.Fatalf("InitCore failed: %s", err.Error())
		// }

		Convey("Given a valid patient Family", func() {
			resource := "Patient"
			fmt.Printf("\n\nGiven a valid Family Name\n")
			//os.Setenv("ACCESS_SECRET", "12345678901234567890123456789012")
			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
			dur := time.Duration(300) * time.Second
			jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
			So(err, ShouldBeNil)
			So(jwt, ShouldNotBeNil)
			//w := httptest.NewRecorder()
			fmt.Printf("Creating Find Request\n")
			req, _ := http.NewRequest("GET", "http://192.1t68.1.117:30201/api/rest/v1/Patient?family=smart", nil)
			//req, _ := http.NewRequest("GET", "api/rest/v1/Patient/123456", nil)
			//req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/api/rest/v1/Documents?patient=1&_count=2", nil)
			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
			//req.Header.Add("tracing-id", "123")
			// vars := map[string]string{
			// 	"resource": "Patient",
			// }
			// req = mux.SetURLVars(req, vars)
			req.Header.Set("Authorization", jwt)
			req.Header.Set("facility", "demo")
			req.Header.Set("Fhir_Version", "r4")
			req.RequestURI = "api/rest/v1/Patient?family=smart"
			fs := `{       
				"fhir_system": {
				"id": "62f1c5dab3070d0b40e7aac1",
				"facilityId": "62e89b57e2da183de83c27a2",
				"facilityName": "Mercy Redding",
				"displayName": "Cerner Open",
				"description": "Medical Records from 2013-Current",
				"fhirVersion": "r4",
				"authUrl": "",
				"ucUrl": "http://192.168.1.117:30300/62f1c5dab3070d0b40e7aac1",
				"fhirUrl": "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d",
				"insert": "false",
				"identifiers": [
					{
						"Name": "Mrn",
						"Value": "urn:oid:2.16.840.1.113883.6.1000|"
					},
					{
						"Name": "Community",
						"Value": "urn:oid:2.16.840.1.113883.3.787.0.0|"
					},
					{
						"Name": "Military",
						"Value": "urn:oid:2.16.840.1.113883.3.42.10001.100001.12|"
					},
					{
						"Name": "Millennium",
						"Value": "https://fhir.cerner.com/ec2458f2-1e24-41c8-b71b-0e701af7583d/codeSet/4|"
					}
				],
				"facilityCode": ""
			}`
			rc := io.NopCloser(strings.NewReader(fs))
			//req.Body, err = json.Marshal([]byte(fs))
			req.Body = rc
			fmt.Printf("Call findResource\n")
			client := &http.Client{}
			resp, err := client.Do(req)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, 200)
			defer resp.Body.Close()
			byte, err := ioutil.ReadAll(resp.Body)
			//bundle := fhir.Bundle{}
			bundle, err := fhir4.UnmarshalBundle(byte)
			So(err, ShouldBeNil)
			So(bundle.ResourceType, ShouldEqual, resource)

			// // if err != nil {
			// // 	return 0, nil, nil, err
			// // }
			// //patientId, err := GetPatientFromBundle(resource, &bundle)

			// if err != nil {
			// 	return 0, nil, nil, err
			// }
			// // if err != nil {
			// // 	log.Println("CacheResourceBundleAndEntries:108  --  Error FindResource Request: ", err.Error())
			// // } else {
			// // 	log.Println("CacheResourceBundleAndEntries:106  --  FindResource Successful")
			// // }
			// findResource(w, req)
			// resp := w.Result()
			// So(resp.StatusCode, ShouldEqual, http.StatusOK)
			// //var bundle []fhir.Bundle

			// err = json.NewDecoder(w.Body).Decode(&bundle)
			// So(err, ShouldBeNil)
			// fmt.Printf("bundle: %s\n", spew.Sdump(bundle))

		})
	})
}

func TestPostFhirPatient(t *testing.T) {
	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
		godotenv.Load("./.env.uc_fhir4_test")
		//os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")

		//os.Setenv("COMPANY", "demo")

		conf, err := service.InitCore("uc_fhir4", "local_test", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		// _, err := service.Initialize()
		// //, err := service.InitCore("uc_core","test", "test")
		// if err != nil {
		// 	t.Fatalf("InitCore failed: %s", err.Error())
		// }

		Convey("Given a valid patient Family", func() {
			fmt.Printf("\n\nGiven a valid DocumentId\n")
			os.Setenv("ACCESS_SECRET", "12345678901234567890123456789012")
			dur := time.Duration(300) * time.Second
			jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
			So(err, ShouldBeNil)
			So(jwt, ShouldNotBeNil)
			w := httptest.NewRecorder()
			fmt.Printf("Creating Search Request\n")
			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Patient?family=smart&given=sandy", nil)
			//req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/api/rest/v1/Documents?patient=1&_count=2", nil)
			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
			//req.Header.Add("tracing-id", "123")
			// vars := map[string]string{
			// 	"doc_id" : "40441",
			// }
			//req = mux.SetURLVars(req, vars)
			req.Header.Set("Authorization", jwt)
			req.Header.Set("facility", "demo")
			req.Header.Set("Fhir_Version", "r4")
			fmt.Printf("")

			findResource(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			var bundle []fhir4.Bundle

			err = json.NewDecoder(w.Body).Decode(&bundle)
			So(err, ShouldBeNil)
			fmt.Printf("bundle: %s\n", spew.Sdump(bundle))

		})
	})
}

func TestFhirDocumentForPatient(t *testing.T) {
	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
		godotenv.Load("./.env.uc_fhir4_test")
		//os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")

		//os.Setenv("COMPANY", "demo")

		conf, err := service.InitCore("uc_fhir4", "local_test", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		// _, err := service.Initialize()
		// //, err := service.InitCore("uc_core","test", "test")
		// if err != nil {
		// 	t.Fatalf("InitCore failed: %s", err.Error())
		// }

		Convey("Given a valid patient Family", func() {
			fmt.Printf("\n\nGiven a valid DocumentId\n")
			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
			dur := time.Duration(300) * time.Second
			jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
			So(err, ShouldBeNil)
			So(jwt, ShouldNotBeNil)
			w := httptest.NewRecorder()
			fmt.Printf("Creating Search Request\n")
			cps := CreateTestFileCloser()
			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Find/DocumentReference?patient=12748336", cps)
			//req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/api/rest/v1/Documents?patient=1&_count=2", nil)
			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
			//req.Header.Add("tracing-id", "123")
			// vars := map[string]string{
			// 	"doc_id" : "40441",
			// }
			//req = mux.SetURLVars(req, vars)

			req.Header.Set("Authorization", jwt)
			req.Header.Set("facility", "demo")
			req.Header.Set("Fhir_Version", "r4")
			fmt.Printf("")

			findResource(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			//var bundle []fhir4.Bundle
			var resResp common.ResourceResponse

			err = json.NewDecoder(w.Body).Decode(&resResp)
			So(err, ShouldBeNil)
			fmt.Printf("ResResp: %s\n", spew.Sdump(resResp))

		})
	})
}

func TestFhirEncountersForPatient(t *testing.T) {
	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
		godotenv.Load("./.env.uc_fhir4_test")
		conf, err := service.InitCore("uc_fhir4", "local_test", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)

		Convey("Given a valid patient Family", func() {
			fmt.Printf("\n\nGiven a valid PatientId\n")
			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
			dur := time.Duration(300) * time.Second
			jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
			So(err, ShouldBeNil)
			So(jwt, ShouldNotBeNil)
			w := httptest.NewRecorder()
			fmt.Printf("Creating Search Request\n")
			cps := CreateTestFileCloser()
			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Find/Encounter?patient=12748336", cps)
			req.Header.Set("Authorization", jwt)
			req.Header.Set("facility", "demo")
			req.Header.Set("Fhir_Version", "r4")
			fmt.Printf("")

			findResource(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			var resResp common.ResourceResponse
			err = json.NewDecoder(w.Body).Decode(&resResp)
			So(err, ShouldBeNil)
			fmt.Printf("ResResp: %s\n", spew.Sdump(resResp))
		})
	})
}

func TestFhirProceduresForPatient(t *testing.T) {
	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
		godotenv.Load("./.env.uc_fhir4_test")
		conf, err := service.InitCore("uc_fhir4", "local_test", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)

		Convey("Given a valid patient Family", func() {
			fmt.Printf("\n\nGiven a valid PatientId\n")
			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
			dur := time.Duration(300) * time.Second
			jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
			So(err, ShouldBeNil)
			So(jwt, ShouldNotBeNil)
			w := httptest.NewRecorder()
			fmt.Printf("Creating Search Request\n")
			cps := CreateTestFileCloser()
			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Find/Procedure?patient=12748336", cps)
			req.Header.Set("Authorization", jwt)
			req.Header.Set("facility", "demo")
			req.Header.Set("Fhir_Version", "r4")
			fmt.Printf("")

			findResource(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			var resResp common.ResourceResponse
			err = json.NewDecoder(w.Body).Decode(&resResp)
			So(err, ShouldBeNil)
			fmt.Printf("ResResp: %s\n", spew.Sdump(resResp))
		})
	})
}
func CreateTestFhirSystem() *common.FhirSystem {
	fs := `{
		"id": "62f1c5dab3070d0b40e7aac1",
		"facilityId": "62e89b57e2da183de83c27a2",
		"facilityName": "Mercy Redding",
		"displayName": "Cerner Open",
		"description": "Medical Records from 2013-Current",
		"fhirVersion": "r4",
		"authUrl": "",
		"ucUrl": "http://192.168.1.117:30300/62f1c5dab3070d0b40e7aac1",
		"fhirUrl": "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d",
		"insert": "false",
		"identifiers": [
			{
				"Name": "Mrn",
				"Value": "urn:oid:2.16.840.1.113883.6.1000|"
			},
			{
				"Name": "Community",
				"Value": "urn:oid:2.16.840.1.113883.3.787.0.0|"
			},
			{
				"Name": "Military",
				"Value": "urn:oid:2.16.840.1.113883.3.42.10001.100001.12|"
			},
			{
				"Name": "Millennium",
				"Value": "https://fhir.cerner.com/ec2458f2-1e24-41c8-b71b-0e701af7583d/codeSet/4|"
			}
		],
		"facilityCode": "",
		"serviceName" : "uc_cerner",
		"returnBundle" : "true",
		"connector" : "uc_cerner:local_test"
	}`
	fsByte := []byte(fs)
	fhirSystem := &common.FhirSystem{}
	err := json.Unmarshal(fsByte, fhirSystem)
	if err != nil {
		fmt.Printf("CreateTestFhirSystem:423  --  Unmarshal fhirSystem err: %s\n", err.Error())
		return nil
	}
	return fhirSystem
}
func CreateTestFileCloser() io.ReadCloser {
	fhirSystem := CreateTestFhirSystem()
	cp := common.ConnectorPayload{}
	cp.FhirSystem = fhirSystem
	cp.FhirSystem = fhirSystem
	//cp := common.ConnectorPayload{}
	cc := common.ConnectorConfig{}
	cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
	cc.Name = "uc_cerner"
	cc.Version = "local_test"
	cc.CacheUrl = "http://192.168.1.117:30201"
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
	cps, err := json.Marshal(cp)
	if err != nil {
		fmt.Printf("CreateTestFileCloser:456  --  Marshal cp failed: %s\n", err.Error())
	}
	rc := io.NopCloser(strings.NewReader(string(cps)))
	return rc
}
