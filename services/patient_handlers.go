package services

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	//"github.com/dhf0820/fhir4"
	fhir "github.com/dhf0820/fhir4"
	token "github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"os"
	//"strconv"
	//"strings"
)

//####################################### Response Writers Functions #######################################

//################################### FHIR Responses ####################################

//####################################### Route Handlers #######################################
// getPatient - By patientId returning one single patient matching the ID
// otherwise return OperationOutcome for NotFound
func getPatient(w http.ResponseWriter, r *http.Request) {
	Resource := "Patient"
	fmt.Printf("getPatient:33 - \n")

	//buildFieldsByTagMap("schema", *psp)

	Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	userId := Payload.UserId
	fmt.Printf("getPatient:44  --  userId: %s\n", userId)
	defer r.Body.Close()
	params := mux.Vars(r)
	fmt.Printf("GetPatient:48  --  Params: %v\n", params)
	id := params["id"]
	//uri := r.URL.RequestURI()
	fmt.Printf("getPatient:50  --  raw: %s\n", r.URL.RawQuery)
	fmt.Printf("GetPatient:51  --  query values: %v\n", r.URL.Query())
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}
	//id := r.URL.Query().Get("id")
	patient, err := GetPatient(id)
	if err != nil {
		err = fmt.Errorf("getPatient:57  --  GetPatient error: %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("getPatient:57  --  Returning Patient: %s\n", spew.Sdump(patient))
	resp := common.ResourceResponse{}
	resp.Resource.Patient = *patient
	resp.Patient = *patient
	resp.Status = 200
	resp.Message = "Ok"
	resp.ResourceType = Resource
	WriteFhirResource(w, 200, &resp)

	// //defer resp.Body.Close()
	// //cfg = mod.ServiceConfig{}
	// fmt.Printf("Reading Body\n")
	// body, err = ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Printf("getPatient:53  --  ReadAllBody : error: %s\n", err.Error())
	// 	//err = errors.New("invalid FHIR URL")
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// connectPayload := common.ConnectorPayload{}
	// fmt.Printf("raw string: %s\n", string(body))
	// fmt.Printf("GetPatient:61  --  Unmarshal ConnectorPayload\n")
	// err = json.Unmarshal(body, &connectPayload)
	// if err != nil {
	// 	log.Printf("getPatient:64  --  Unmarshal connectPayload error: %s\n", err.Error())
	// 	//err = errors.New("invalid FHIR URL")
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// params := mux.Vars(r)
	// id := params["id"]
	// if id != "" {
	// 	fmt.Printf("GetPatient:73  -- Specific Query: %s\n", id)
	// 	//patient, err := GetPatient(id)
	// }

	// //fmt.Printf("getPatient:70 -- connectPayload: %s\n", spew.Sdump(connectPayload))
	// // fhirId := GetFhirId(r)
	// // fhirSystem, err := GetFhirSystem(fhirId)
	// // if err != nil {
	// // 	log.Printf("searchPatient:50  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// // 	err = errors.New("invalid FHIR URL")
	// // 	errMsg := err.Error()
	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// // 	return
	// // }

	// //  Separate connector for each emr Vendor.  CernerConnector, EpicConnector, CAConnector,...
	// // // handles query andsaveeither via fhir or direct API (AllScripts, Athena)

	// // Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	// // if err != nil {
	// // 	errMsg := err.Error()
	// // 	fmt.Printf("getPatient:55  --  Err: %s\n", errMsg)
	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// // 	return
	// // }
	// // userId := Payload.UserId
	// // log.Printf("getPatient:59  --  UserId: %s\n", userId)
	// // fhirId := GetFhirId(r)
	// // fhirSystem, err := GetFhirSystem(fhirId)
	// // if err != nil {
	// // 	log.Printf("getPatient:63  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// // 	err = errors.New("invalid FHIR URL")
	// // 	errMsg := err.Error()
	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// // 	return
	// // }
	// fmt.Printf("getPatient:109  --  Request: [%s]\n", r.RequestURI)
	// urlA, err := r.URL.Parse(r.RequestURI)
	// if err != nil {
	// 	err = fmt.Errorf("error parsing patient URI: %s", err.Error())
	// 	errMsg := err.Error()
	// 	fmt.Printf("getPatient:114 - r.URL.Parse error = %s\n", errMsg)
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("getPatient:118 - r.URL.Parse = %v\n", urlA)
	// urlB := *urlA
	// uriValues := urlB.Query()
	// fmt.Printf("getPatient:121 - uriValues= %v\n", uriValues)

	// uri := r.RequestURI
	// log.Printf("uri = %s\n", uri)
	// parts := strings.Split(uri, Resource)
	// uri = parts[1]
	// log.Printf("getPatient:127 - URI = %s\n", uri)
	// //patient := fhir.Patient{}
	// resource, err := GetResource(connectPayload.FhirSystem, Resource, uri)
	// resp := common.ResourceResponse{}
	// if err != nil {
	// 	resp.Status = 400
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Status = 200
	// 	resp.Message = "Ok"
	// }
	// //var patient fhir.Patient
	// //patient := resource.(fhir.Patient)
	// resp.Resource.Resource = resource
	// // var res []interface{}
	// // res = append(res, &resource)
	// // resp.Resources = res
	// resp.ResourceType = Resource
	// //resp.ResourceId = *patient.Id
	// //log.Printf("\nGetPatient:139  --  resp: %s\n", spew.Sdump(resp))

	// WriteFhirResource(w, resp.Status, &resp)
}

//postPatient: Stores the fhir patient payload in the url {Fhir-System} specified fhirSystem.
func savePatient(w http.ResponseWriter, r *http.Request) {
	//Resource := "Patient"

	//fmt.Printf("postPatient:148 - Post: %s \n", spew.Sdump(r))

	Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	if err != nil {
		errMsg := err.Error()
		fmt.Printf("postPatient:160  - ValidateToken err = %s\n", errMsg)
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	//userId := Payload.UserId
	log.Printf("savePatient:165  --  User Name: %s\n", Payload.FullName)
	body, err := ioutil.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		fmt.Printf("savePatient:168  --  ReadAll FhirSystem error %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	//b := string(body)
	//fmt.Printf("SavePatient:167  Body: %s\n", b)
	conPayload := common.ConnectorPayload{}
	//fhirSystem := common.FhirSystem{}
	err = json.Unmarshal(body, &conPayload)
	if err != nil {
		fmt.Printf("\nSavePatient:179  --  unmarshal err = %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	patId, patMrn, text, err := SavePatient("", conPayload.SavePayload.SrcPatient)
	fmt.Printf("savePatient:209  --  Patient.Text: %s,  id: %s,  MRN: %s\n", text, patId, patMrn)
	resp := &common.SaveResponse{}
	resp.Id = patId
	resp.Text = text
	resp.Mrn = patMrn
	WriteSaveResponse(w, 200, resp)
	// fhirId := GetFhirId(r)                   // Get the Fhir-System ID portion of the URL
	// fhirSystem, err := GetFhirSystem(fhirId) // Get the actual FhirSystem Configuration
	// if err != nil {
	// 	log.Printf("postPatient:162  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// 	err = errors.New("url contains Invalid FHIR identifier: " + fhirId)
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// urlA, err := r.URL.Parse(r.RequestURI)
	// if err != nil {
	// 	err = fmt.Errorf("error parsing patient URI: [%s]  error:%s", r.RequestURI, err.Error())
	// 	errMsg := err.Error()
	// 	fmt.Printf("postPatient:172 - r.URL.Parse error = %s\n", errMsg)
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("postPatient:176 - r.URL.Parse = %v\n", urlA)
	// urlB := *urlA
	// uriValues := urlB.Query()
	// fmt.Printf("postPatient:179 - uriValues= %v\n", uriValues)

	// uri := r.RequestURI
	// log.Printf("uri = %s\n", uri)
	// parts := strings.Split(uri, Resource)
	// uri = parts[1]
	// log.Printf("postPatient:185 - URI = %s\n", uri)
	// //patient := fhir.Patient{}
	//WriteFhirResource(w, 200, resp)
	// errMsg := "SavePatient to CA not implemented"
	// WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// return
	// resource, err := GetResource(fhirSystem, Resource, uri)
	// resp := common.ResourceResponse{}
	// if err != nil {
	// 	resp.Status = 400
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Status = 200
	// 	resp.Message = "Ok"
	// }
	// var patient fhir.Patient
	// //patient = resource.(fhir.Patient)
	// resp.Resource.Resource = resource
	// // var res []interface{}
	// // res = append(res, &resource)
	// // resp.Resources = res
	// resp.ResourceType = Resource
	// resp.ResourceId = *patient.Id
	// log.Printf("\nGetPatient:204  --  resp: %s\n", spew.Sdump(resp))
	// WriteFhirResourceBundle(w, resp.Status, &resp)
}

// searchPatient uses the fhirId url parameter to determin the FhirSystem to use
func searchPatient(w http.ResponseWriter, r *http.Request) {
	// var pspTags map[string]string
	// tagFields := make(map[string]string)
	// var Limit int
	// var Skip int
	//Resource := "Patient"
	body, err := ioutil.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		fmt.Printf("findResource:217  --  ReadAll FhirSystem error %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	connectorPayload := common.ConnectorPayload{}
	//fhirSystem := common.FhirSystem{}
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		fmt.Printf("\nfindResource:226  --  unmarshal err = %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	//fhirSystem := connectorPayload.FhirSystem
	//connConfig := connectorPayload.ConnectorConfig
	//buildFieldsByTagMap("schema", *psp)
	JWToken = r.Header.Get("Authorization")
	//fmt.Printf("searchPatient:219 - JWToken: %s\n", JWToken)
	Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	if err != nil {
		errMsg := err.Error()
		fmt.Printf("searchPatient:239  --  ValidateToken err: %s\n", errMsg)
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	//fhirId := fhirSystem.ID.String()
	userId := Payload.UserId
	log.Printf("searchPatient:300  --  UserId: %s\n", userId)

	fmt.Printf("SearchPatient:302  --  raw: %s\n", r.URL.RawQuery)
	fmt.Printf("SearchPatient:303  --  query values: %v\n", r.URL.Query())
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}

	/*
		// fhirId := GetFhirId(r)
		// fhirSystem, err := GetFhirSystem(fhirId)
		// if err != nil {
		// 	log.Printf("searchPatient:232  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
		// 	err = errors.New("invalid FHIR URL")
		// 	errMsg := err.Error()
		// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		// 	return
		// }
		uri := r.RequestURI
		log.Printf("searchPatient:314  --  r.RequestURI = %s\n", uri)
		parts := strings.Split(uri, Resource)
		uri = parts[1]
		log.Printf("\nsearchPatient:260 - URI = %s\n", uri)

		urlA, err := r.URL.Parse(r.RequestURI)
		if err != nil {
			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
			errMsg := err.Error()
			fmt.Printf("searchPatient:266 - r.URL.Parse error = %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		fmt.Printf("searchPatient:270 - r.URL.Parse = %v\n", urlA)
		urlB := *urlA
		uriValues := urlB.Query()
		fmt.Printf("searchPatient:273 - uriValues= %v\n", uriValues)
		//ident := uriValues.Get("identifier")
		// if ident != "" { // There is identifier Search, use it
		// 	fmt.Printf("searchPatient:102 - using Identifier: %s to search\n", ident)
		// } else {
		// 	fmt.Printf("searchPatient:104 - using other search params: %v\n", uriValues)
		// }

		// //}
		// //fhirVersion := GetFHIRVersion(r)
		// //cacheBaseURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
		// if err := r.ParseForm(); err != nil {
		// 	err = fmt.Errorf("error parsing query: %s", err.Error())
		// 	errMsg := err.Error()
		// 	fmt.Printf("searchPatient:113 - %s\n", errMsg)
		// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		// 	return
		// }
		// FhirId := GetFhirId(r)
		// fmt.Printf("searchPatient:79 - FhirKey - [%s]\n", FhirId)
		// fhirSystem, err := GetFhirSystem(FhirId)
		// if err != nil {
		// 	fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
		// 	err = fmt.Errorf("fhirSystem error:  %s", err.Error())
		// 	errMsg := err.Error()
		// 	fmt.Printf("searchPatient:86 - %s\n", errMsg)
		// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
		// 	return
		// }
		//fmt.Printf("searchPatient:90 -  %s/n", spew.Sdump(fhirSystem))

		// if Resource == "Patient" {
		log.Printf("\n\nsearchPatient:305  --  Resource Is Patient\n\n")
		//urlA, err := r.URL.Parse(r.RequestURI)
		if err != nil {
			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
			errMsg := err.Error()
			fmt.Printf("searchPatient:310 - r.URL.Parse error = %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		fmt.Printf("searchPatient:314 - r.URL.Parse = %v\n", urlA)
		//urlB := *urlA
		//uriValues := urlB.Query()
		fmt.Printf("searchPatient:317 - uriValues= %v\n", uriValues)
		idSearch := uriValues.Get("identifier")
		idValue := ""
		if idSearch != "" { // There is identifier Search, use it
			fmt.Printf("searchPatient:321- using Identifier: %s to search\n", idSearch)
			ids := strings.Split(idSearch, "|")
			if len(ids) != 2 {
				err = fmt.Errorf("invalid identifier: %s", idSearch)
				errMsg := err.Error()
				fmt.Printf("searchPatient:326 - r.URL.Parse error = %s\n", errMsg)
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			idName := ids[0]
			idSearchValue := ids[1]
			idents := fhirSystem.Identifiers
			for _, id := range idents {
				fmt.Printf("searchPatient:334  --  Looking at %s = %s\n", id.Name, idName)
				if id.Name == idName {
					idValue = id.Value
					break
				}
			}
			//
			if idValue == "" { //Not configured identifier
				err = fmt.Errorf("identifier type: %s is not configured", idName)
				errMsg := err.Error()
				fmt.Printf("searchPatient:344 - Identifiers = %s\n", errMsg)
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			uri = fmt.Sprintf("?identifier=%s", idValue+idSearchValue)
			fmt.Printf("searchPatient:349 - New Identifier search Value: %s\n", uri)
		} else {
			fmt.Printf("searchPatient:351 - using other search params: %v\n", uriValues)
		}
		var bundle *fhir.Bundle
		var header *common.CacheHeader
		fmt.Printf("\nsearchPatient:355 - resource = %s  uri = %s\n", Resource, uri)
		url := fmt.Sprintf("%s/%s%s", fhirSystem.FhirUrl, Resource, uri) //" + "/" + uri
		fmt.Printf("searchPatient:357 - calling %s \n", url)
		var totalPages int64
		fmt.Printf("searchPatient:359 Search %s\n", url)
		uri = "/" + Resource + uri
		totalPages, bundle, header, err = FindResource(&connectorPayload, Resource, userId, uri, JWToken)
		if err != nil {
			err = fmt.Errorf("searchPatient:363 --  fhirSearch url: %s error:  %s", url, err.Error())
			errMsg := err.Error()
			fmt.Printf("searchPatient:365 - %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
			return
		}
		if bundle == nil {
			log.Printf("searchPatient:370  --  bundle is nil")
		} else {
			log.Printf("searchPatient:372  --  bundle is not nil \n")
		}
		fmt.Printf("searchPatient:374 - Get %s bundle successful\n", Resource)
		fmt.Printf("searchPatient:375 - Number in page: %d\n", len(bundle.Entry))
		fmt.Printf("searchPatient:376 - PageNumber: %d\n", header.PageId)
		resp := common.ResourceResponse{}
		header.CacheBase = fmt.Sprintf("%s/%s/BundleTransaction", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
		log.Printf("\n\nsearchPatient:379  --  CacheUrl = %s\n", header.CacheBase)
		header.FhirId = fhirId
		header.UserId = userId
		resp.Bundle = bundle
		resp.Resource.Resource = bundle.Entry[0].Resource
		resp.BundleId = *bundle.Id
		resp.ResourceType = Resource
		resp.Status = 200
		resp.QueryId = header.QueryId
		resp.PageNumber = header.PageId
		if bundle.Entry == nil {
			err = fmt.Errorf("searchPatient:390 --  fhirSearch url: %s error:  %s", url, "Bundle.Entry is nil")
			errMsg := err.Error()
			fmt.Printf("searchPatient:392 - %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
			return
		}
		resp.CountInPage = len(bundle.Entry)
		resp.TotalPages = totalPages
		resp.Header = header
		resp.Message = "Ok"
		//fmt.Printf("searchPatient:400 - returning a resource bundle: %s\n", spew.Sdump(resp))
		WriteFhirResourceBundle(w, resp.Status, &resp)
		//WriteFhirBundle(w, resp.Status, bundle)
	*/
}

//func searchPatient(w http.ResponseWriter, r *http.Request) {
// 	// var pspTags map[string]string
// 	// tagFields := make(map[string]string)
// 	// var Limit int
// 	// var Skip int
// 	fmt.Printf("Request: %s \n", spew.Sdump(r))
// 	//buildFieldsByTagMap("schema", *psp)
// 	//facility = "demo"
// 	resource := GetFHIRResource(r)
// 	fmt.Printf("search%s called with %s\n", resource, r.URL.RawQuery)
// 	if err := r.ParseForm(); err != nil {
// 		err = fmt.Errorf("error parsing query: %s", err.Error())
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	params := mux.Vars(r)
// 	fmt.Printf("params: %v\n", params)
// 	resourceId := params["id"]
// 	fmt.Printf("Retrieving Patient Record for id: %s\n", resourceId)
// 	// psp := new(PatientSearchParams)
// 	// psp.Limit = Limit
// 	// psp.Skip = Skip
// 	// psp.CurrentFacility = GetDeploymentFacility(r)
// 	// psp.BaseUrl = GetCurrentURL(r)
// 	//FhirVersion := GetFHIRVersion(r)
// 	FhirId := GetFhirId(r)
// 	_, err := GetFhirConnector(FhirId)
// 	if err != nil {
// 		err = fmt.Errorf("fhirConnecor error:  %s", err.Error())
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	}

// 	//bundle, err := FindResource(fhirConnector, resource, r.URL.RawQuery)

// }

// func WriteFhirOperationOutcome(w http.ResponseWriter, status int, resp *fhir.OperationOutcome) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

// func WriteFhirPatientBundle(w http.ResponseWriter, status int, resp *ResourceResponse) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }
