package services

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	//"github.com/dhf0820/fhir4"
	cm "github.com/dhf0820/ca3_connector/common"
	"github.com/dhf0820/fhir4"
	fhir "github.com/dhf0820/fhir4"
	common "github.com/dhf0820/uc_common"
	"github.com/gorilla/mux"

	//"os"
	//"strconv"
	"strings"
	"time"

	token "github.com/dhf0820/token"
)

//####################################### Response Writers Functions #######################################
func WriteFhirOperationOutcome(w http.ResponseWriter, status int, resp *fhir.OperationOutcome) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteSaveResponse(w http.ResponseWriter, status int, resp *common.SaveResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteFhirResource(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}
func WriteFhirResourceBundle(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteFhirBundle(w http.ResponseWriter, status int, resp *fhir4.Bundle) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteFhirResponse(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

//################################### FHIR Responses ####################################

func CreateOperationOutcome(code fhir.IssueType, severity fhir.IssueSeverity, details *string) *fhir.OperationOutcome {
	fmt.Printf("Error Message : %s\n", *details)
	s := *details
	outcome := fhir.OperationOutcome{}
	issue := fhir.OperationOutcomeIssue{}
	issue.Code = code
	issue.Severity = severity
	issue.Details = &fhir.CodeableConcept{}
	issue.Details.Text = &s
	outcome.Issue = append(outcome.Issue, issue)
	return &outcome
}

//####################################### Route Handlers #######################################

// // Header will have the fhir services token
// // Header may or url will have the id of the FhirConnector
// func searchPatient(w http.ResponseWriter, r *http.Request) {
// 	// var pspTags map[string]string
// 	// tagFields := make(map[string]string)
// 	// var Limit int
// 	// var Skip int
// 	fmt.Printf("searchPatient:86 - Request: %s \n", spew.Sdump(r))
// 	//buildFieldsByTagMap("schema", *psp)
// 	//facility = "demo"
// 	userId := r.Header.Get("UserId")
// 	uri := r.RequestURI
// 	parts := strings.Split(uri, "v1/")
// 	uri = parts[1]
// 	fmt.Printf("searchPatient:93 - URI = %s\n", uri)
// 	resource := GetFHIRResource(r)

// 	fhirVersion := r.Header.Get("FhirVersion")
// 	if fhirVersion == "" {
// 		fhirVersion = "R4"
// 	}
// 	// if resource == "Patient" {
// 	// 	urlA, err := r.URL.Parse(r.RequestURI)
// 	// 	if err != nil {
// 	// 		err = fmt.Errorf("error parsing patient URI: %s", err.Error())
// 	// 		errMsg := err.Error()
// 	// 		fmt.Printf("findResource:102 - r.URL.Parse error = %s\n", errMsg)
// 	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 		return
// 	// 	}
// 	// 	fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
// 	// 	urlB := *urlA
// 	// 	uriValues := urlB.Query()
// 	// 	fmt.Printf("findResource:109 - uriValues= %v\n", uriValues)
// 	// 	ident := uriValues.Get("identifier")
// 	// 	if ident != "" { // There is identifier Search, use it
// 	// 		fmt.Printf("findResource:110 - using Identifier: %s to search\n", ident)
// 	// 	} else {
// 	// 		fmt.Printf("findResource:110 - using other search params: %v\n", uriValues)
// 	// 	}

// 	// }
// 	//fhirVersion := GetFHIRVersion(r)
// 	//cacheBaseURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
// 	if err := r.ParseForm(); err != nil {
// 		err = fmt.Errorf("error parsing query: %s", err.Error())
// 		errMsg := err.Error()
// 		fmt.Printf("searchPatient:126 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	FhirId := GetFhirId(r)
// 	fmt.Printf("findResource:132 - FhirKey - [%s]\n", FhirId)
// 	fhirSystem, err := GetFhirSystem(FhirId)

// 	if err != nil {
// 		fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
// 		err = fmt.Errorf("fhirSystem error:  %s", err.Error())
// 		errMsg := err.Error()
// 		fmt.Printf("searchPatient:138 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	fmt.Printf("searchPatient:142 -  %s/n", spew.Sdump(fhirSystem))

// 	if resource == "Patient" {
// 		urlA, err := r.URL.Parse(r.RequestURI)
// 		if err != nil {
// 			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("searchPatient:149 - r.URL.Parse error = %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		//fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
// 		urlB := *urlA
// 		uriValues := urlB.Query()
// 		fmt.Printf("searchPatient:156 - uriValues= %v\n", uriValues)
// 		idSearch := uriValues.Get("identifier")
// 		idValue := ""
// 		if idSearch != "" { // There is identifier Search, use it
// 			fmt.Printf("searchPatient:160 - using Identifier: %s to search\n", idSearch)
// 			ids := strings.Split(idSearch, "|")
// 			if len(ids) != 2 {
// 				err = fmt.Errorf("invalid identifier: %s", idSearch)
// 				errMsg := err.Error()
// 				fmt.Printf("searchPatient:165 - r.URL.Parse error = %s\n", errMsg)
// 				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 				return
// 			}
// 			idName := ids[0]
// 			idSearchValue := ids[1]
// 			idents := fhirSystem.Identifiers
// 			for _, id := range idents {
// 				if id.Name == idName {
// 					idValue = id.Value
// 					break
// 				}
// 			}
// 			if idValue == "" { //Not configured identifier
// 				err = fmt.Errorf("identifier type: %s is not configured", idName)
// 				errMsg := err.Error()
// 				fmt.Printf("searchPatient:181 - Identifiers = %s\n", errMsg)
// 				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 				return
// 			}
// 			uri = fmt.Sprintf("Patient?identifier=%s", idValue+idSearchValue)
// 			fmt.Printf("searchPatient:186 - New Identifier search Value: %s\n", uri)
// 			// uriValues.Del("identifier")
// 			// uriValues.Set("identifier", id)
// 			// //urlB.RawQuery = uriValues.Encode()
// 			// r.URL.RawQuery = uriValues.Encode()
// 			// //curUri := r.RequestURI
// 			// //urUriParts := strings.Split(curUri, "?")
// 			// r.RequestURI = uriValues.Encode()
// 			// fmt.Printf("\n\n$$$ searchResources: 188 - Updated request: %s\n\n", spew.Sdump(r))
// 		} else {
// 			fmt.Printf("searchPatient:196 - using other search params: %v\n", uriValues)
// 		}

// 	}
// 	var bundle *fhir.Bundle
// 	var header *common.CacheHeader
// 	//resourceId := r.Header.Get("Fhir-System")
// 	// params := mux.Vars(r)
// 	// fmt.Printf("findResource params:115 %v\n", params)

// 	//resource := strings.Split(uri, "?")[0]
// 	fmt.Printf("\nsearchPatient:207 - resource = %s  uri = %s\n", resource, uri)
// 	url := fhirSystem.FhirUrl + "/" + uri
// 	fmt.Printf("searchPatient:209 - calling %s \n", url)
// 	var totalPages int64
// 	// if resourceId != "" {
// 	// 	fmt.Printf("findResource:128 - Get %s with [%s]\n", resource, resourceId)
// 	// } else {
// 	fmt.Printf("searchPatient:214 Search %s with %s\n", url, r.RequestURI)
// 	totalPages, bundle, header, err = SearchPatient(fhirSystem, url, resource, userId, r.RequestURI)
// 	if err != nil {
// 		err = fmt.Errorf("fhirSearch url: %s error:  %s", url, err.Error())
// 		errMsg := err.Error()
// 		fmt.Printf("searchPatient:219 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
// 		return
// 	}
// 	//}
// 	fmt.Printf("searchPatient:224 - Get %s bundle successful\n", resource)
// 	fmt.Printf("searchPatient:225 - Number in page: %d\n", len(bundle.Entry))
// 	fmt.Printf("searchPatient:226 - PageNumber: %d\n", header.PageId)
// 	resp := ResourceResponse{}
// 	//hostParts := strings.Split(r.Host, ":")

// 	host := common.GetKVData(GetConfig().Data, "cacheHost")
// 	//host := os.Getenv("CORE_ADDRESS")
// 	fmt.Printf("searchPatient:232 - ##host: %s\n\n\n", host)
// 	header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)

// 	resp.Bundle = bundle
// 	resp.Resource = header.ResourceType
// 	resp.BundleId = *bundle.Id
// 	resp.ResourceType = resource
// 	resp.Status = 200
// 	resp.QueryId = header.QueryId
// 	resp.PageNumber = header.PageId
// 	resp.CountInPage = len(bundle.Entry)
// 	resp.TotalPages = totalPages
// 	resp.Header = header
// 	resp.Message = "Ok"
// 	fmt.Printf("searchPatient:246 - returning a resource bundle\n")
// 	WriteFhirResourceBundle(w, resp.Status, &resp)

// }
var Resource string
var JWToken string
var Payload *token.Payload

func DebbieTest(w http.ResponseWriter, r *http.Request) {
	_, err := cm.FhirPatientSearch(r)
	if err != nil {
		msg := err.Error()
		WriteFhirOperationOutcome(w, 200, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &msg))
		return
	}
	params := mux.Vars(r)
	fmt.Printf("DebbieTest324  --  Prmd: %v\n", params)
	Resource := ""
	uri := r.URL.RequestURI()
	fmt.Printf("DebbieTest:327  --  uri: %s\n", uri)
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}

	return
	parts := strings.Split(uri, "api/")
	p0 := parts[1]
	fmt.Printf("searchPatient:93 - URI-p0 = %s\n", p0)
	findParts := strings.Split(p0, "?")
	fmt.Printf("uriParts length = %d\n\n", len(findParts))
	if len(findParts) == 1 { // is a getResource
		getParts := strings.Split(p0, "/")
		fmt.Printf("getParts = %v\n", getParts)
		Resource = getParts[0]
	} else {
		Resource = findParts[0]
	}
	msg := fmt.Sprintf("Resource = %s\n", Resource)
	WriteFhirOperationOutcome(w, 200, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &msg))
}

func findResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Resource := params["resource"]
	//log.Printf("findResource:326  --  params = %s\n", params)
	fmt.Printf("findResource:346  --  being called for resource: [%s]\n", Resource)
	//fmt.Printf("findResource:347 - Request: ")
	//spew.Dump(r)
	// if true {
	// 	return
	// }
	body, err := ioutil.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		fmt.Printf("findResource:354  --  ReadAll FhirSystem error %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	connectorPayload := common.ConnectorPayload{}

	//fhirSystem := common.FhirSystem{}
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		fmt.Printf("\n365  --  unmarshal err = %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fhirSystem := connectorPayload.FhirSystem
	connectorConfig := connectorPayload.ConnectorConfig
	fmt.Printf("findResource:371  -- ConnectorPayload = %s\n", spew.Sdump(connectorPayload))
	JWToken = r.Header.Get("Authorization")
	Payload, status, err := token.ValidateToken(JWToken, "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	userId := Payload.UserId
	//Resource := r.Header.Get("Resource")
	uri := r.URL.RequestURI()
	fmt.Printf("findResource:383  --  uri: %s\n", uri)
	parts := strings.Split(uri, "/Find/")
	fmt.Printf("findResource:385  -- parts[1] = %v\n", parts[1])
	uri = parts[1]
	// p0 := parts[1]
	fmt.Printf("findResource:388  --  part0 = %s\n", parts[0])
	fmt.Printf("findResource:389  --  part1 = %s\n", parts[1])
	findParts := strings.Split(uri, "?")
	fmt.Printf("findResource:391  --  uriParts length = %d p0 = %s,  p1 = %s\n\n", len(findParts), findParts[0], findParts[1])
	if len(findParts) == 1 { // is a getResource
		getParts := strings.Split(uri, "/")
		fmt.Printf("findResource:394  --  getParts = %v\n", getParts)
		Resource = getParts[0]
		// } else {
		// 	Query = p0
		// 	Resource = findParts[0]
	} else {
		Resource = findParts[0]
	}

	//Resource = r.Header.Get("Resource")

	//buildFieldsByTagMap("schema", *psp)
	//facility = "demo"
	//userId := r.Header.Get("UserId")
	//fhirId := r.Header.Get("Fhir-System")
	// FhirId := GetFhirId(r)
	// fmt.Printf("findResource:275 - FhirKey - [%s]\n", FhirId)
	// fhirSystem, err := GetFhirSystem(FhirId)

	// if err != nil {
	// 	fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
	// 	err = fmt.Errorf("fhirSystem error:  %s", err.Error())
	// 	errMsg := err.Error()
	// 	fmt.Printf("findResource:282 - %s\n", errMsg)
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Errorf("findResource:374  --  ReadAll error %s\n", err.Error())
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("FindResource:379  --  bodyString = %s\n\n", string(body))
	// fhirSystem := common.FhirSystem{}
	// err = json.Unmarshal(body, &fhirSystem)
	// if err != nil {
	// 	fmt.Printf("\nFindResource:383  --  unmarshal err = %s\n", err.Error())
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("FindResource:388  -- FhirSystem = %s\n", spew.Sdump(fhirSystem))

	//TODO: Split on the Resource keeping the actual url variables and query params
	//uri := r.RequestURI
	// fmt.Printf("findResource:341  --  uri = %s\n", r.RequestURI)
	// u := r.URL
	// fmt.Printf("findResource:342  --  %s\n", u.RequestURI())
	// fmt.Printf("findResource:343  --  URL = %s\n", spew.Sdump(u))
	// uri := r.URL.RequestURI()
	// uriParts := strings.Split(uri, "/")
	// queryString := ""
	// i := 1
	// for _, part := range uriParts {
	// 	log.Printf("uri part: %d = %s\n", i, part)
	// 	i++
	// }
	// parts := strings.Split(uriParts[len(uriParts)-1], "?")
	// if len(parts) > 1 {
	// 	//url query and keep the query element[1]
	// 	//element[0] is the resource
	// 	queryString = parts[len(parts)-1]
	// } else {
	// 	// last element is the id of the query
	// 	queryString = parts[len(uriParts)-1]
	// }

	// log.Printf("QueryString = %s\n", queryString)

	// fmt.Printf("findResource:362 - URI = %s\n", uri)
	// resource := Resource
	//resource := GetFHIRResource(r)

	//fhirVersion := "R4"
	//fhirVersion := r.Header.Get("FhirVersion")
	// if fhirVersion == "" {
	// 	fhirVersion = "R4"
	// }
	// if resource == "Patient" {
	// 	urlA, err := r.URL.Parse(r.RequestURI)
	// 	if err != nil {
	// 		err = fmt.Errorf("error parsing patient URI: %s", err.Error())
	// 		errMsg := err.Error()
	// 		fmt.Printf("findResource:102 - r.URL.Parse error = %s\n", errMsg)
	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 		return
	// 	}
	// 	fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
	// 	urlB := *urlA
	// 	uriValues := urlB.Query()
	// 	fmt.Printf("findResource:109 - uriValues= %v\n", uriValues)
	// 	ident := uriValues.Get("identifier")
	// 	if ident != "" { // There is identifier Search, use it
	// 		fmt.Printf("findResource:110 - using Identifier: %s to search\n", ident)
	// 	} else {
	// 		fmt.Printf("findResource:110 - using other search params: %v\n", uriValues)
	// 	}
	// }
	//fhirVersion := GetFHIRVersion(r)
	//cacheBaseURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
	if err := r.ParseForm(); err != nil {
		err = fmt.Errorf("error parsing query: %s", err.Error())
		errMsg := err.Error()
		fmt.Printf("findResource:497 - %s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}

	//fmt.Printf("findResource:143 -  %s/n", spew.Sdump(fhirSystem))

	if Resource == "Patient" {
		urlA, err := r.URL.Parse(r.URL.RequestURI())
		if err != nil {
			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
			errMsg := err.Error()
			fmt.Printf("findResource:509 - r.URL.Parse error = %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		//fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
		urlB := *urlA
		uriValues := urlB.Query()
		fmt.Printf("\n\n\nfindResource:516 - uriValues= %v\n", uriValues)
		idSearch := uriValues.Get("identifier")
		idValue := ""
		if idSearch != "" { // There is identifier Search, use it
			fmt.Printf("findResource:520 - using Identifier: %s to search\n", idSearch)
			ids := strings.Split(idSearch, "|")
			if len(ids) != 2 {
				err = fmt.Errorf("invalid identifier: %s", idSearch)
				errMsg := err.Error()
				fmt.Printf("searchResource:525 - r.URL.Parse error = %s\n", errMsg)
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			idName := ids[0]
			idSearchValue := ids[1]
			idents := fhirSystem.Identifiers
			for _, id := range idents {
				if id.Name == idName {
					idValue = id.Value
					break
				}
			}
			if idValue == "" { //Not configured identifier
				err = fmt.Errorf("identifier type: %s is not configured", idName)
				errMsg := err.Error()
				fmt.Printf("findResource:541 - Identifiers = %s\n", errMsg)
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			uri = fmt.Sprintf("Patient?identifier=%s", idValue+idSearchValue)
			fmt.Printf("findResource:546 - New Identifier search Value: %s\n", uri)
			// uriValues.Del("identifier")
			// uriValues.Set("identifier", id)
			// //urlB.RawQuery = uriValues.Encode()
			// r.URL.RawQuery = uriValues.Encode()
			// //curUri := r.RequestURI
			// //urUriParts := strings.Split(curUri, "?")
			// r.RequestURI = uriValues.Encode()
			// fmt.Printf("\n\n$$$ searchResources: 188 - Updated request: %s\n\n", spew.Sdump(r))
		} else {
			fmt.Printf("findResource:556 - using other search params: %v\n", uriValues)
		}

	}
	var bundle *fhir.Bundle
	var header *common.CacheHeader
	//resourceId := r.Header.Get("Fhir-System")
	fmt.Printf("\nfindResource:564  - connectorPayload = %s\n", spew.Sdump(connectorPayload))
	resource := strings.Split(uri, "?")[0]
	fmt.Printf("\nfindResource:566 - resource = %s  uri = %s\n", resource, uri)
	url := fhirSystem.FhirUrl + "/" + uri
	fmt.Printf("findResource:568 - calling %s \n", url)
	var totalPages int64
	// if resourceId != "" {
	// 	fmt.Printf("findResource:128 - Get %s with [%s]\n", resource, resourceId)
	// } else {
	fmt.Printf("findResource:573  --  Search %s with %s\n", Resource, uri)
	startTime := time.Now()
	totalPages, bundle, header, err = FindResource(&connectorPayload, Resource, userId, uri, JWToken)
	if err != nil {
		err = fmt.Errorf("findResource:577  --  fhirSearch url: %s  --error:  %s\n", url, err.Error())
		errMsg := err.Error()
		fmt.Printf("%s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
		return
	}
	//}
	fmt.Printf("findResource:584 - Get %s bundle successful in %s\n", Resource, time.Since(startTime))
	fmt.Printf("findResource:585 - Total Pages: %d\n", totalPages)
	fmt.Printf("findResource:586 - Number in page: %d\n", len(bundle.Entry))
	fmt.Printf("findResource:587 - PageNumber: %d\n", header.PageId)
	fmt.Printf("findResource:588 - QueryId: %s\n", header.QueryId)

	resp := common.ResourceResponse{}
	//fmt.Printf("findResource:586 - Header: %s\n", spew.Sdump(header))
	host := fhirSystem.UcUrl
	//host := common.GetKVData(GetConfig().Data, "cacheHost")
	fmt.Printf("findResource:594 - ##host: %s\n\n\n", host)
	cacheBundleUrl := fmt.Sprintf("%s/%s/BundleTransaction", connectorConfig.CacheUrl, header.FhirSystem.ID.Hex())
	//header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)
	fmt.Printf("findResource:597  --  CacheUrl = %s\n", cacheBundleUrl)
	//resp.Resource = header.ResourceType
	resp.BundleId = *bundle.Id
	resp.ResourceType = Resource
	resp.Status = 200
	resp.QueryId = header.QueryId
	resp.PageNumber = header.PageId
	resp.CountInPage = len(bundle.Entry)
	resp.TotalPages = totalPages
	resp.Header = header
	resp.Message = "Ok"
	logTime := time.Now()
	fmt.Printf("findResource:609  --  resp without bundle: %s\n", spew.Sdump(resp))
	fmt.Printf("findResource:610  --  Time to log = %s\n\n", time.Since(logTime))
	resp.Bundle = bundle
	fmt.Printf("findResource:612  --  Number of entries in buldle: %d\n", len(bundle.Entry))
	fmt.Printf("findResource:609  --  QueryId: %s\n\n", header.QueryId)
	//fmt.Printf("findResource:614  --  Returning Bundle: %s\n", spew.Sdump(bundle))
	//WriteFhirResourceBundle(w, resp.Status, &resp)
	WriteFhirResponse(w, resp.Status, &resp)
}

func getResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resource := params["resource"]
	resourceId := params["resId"]
	log.Printf("getResource:595  --  params = %s\n", params)
	fmt.Printf(" param resource : %s   ResId: %s\n", resource, resourceId)

	fmt.Printf("getResource:598 - Request - ")
	spew.Dump(r)
	// Resource = r.Header.Get("Resource")
	// resource := Resource
	fmt.Printf("getResource:602  --  Resource = %s\n", Resource)
	if err := r.ParseForm(); err != nil {
		err = fmt.Errorf("error parsing query: %s", err.Error())
		errMsg := err.Error()
		fmt.Printf("getResurce:606  - %s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeInvalid, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	FhirId := GetFhirId(r)
	fhirSystem, err := GetFhirSystem(FhirId)
	if err != nil {
		err = fmt.Errorf("fhirConnetcor error:  %s", err.Error())
		errMsg := err.Error()
		fmt.Printf("getResource:615 -  GetFhirSystem err = %s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	// uriParts := strings.Split(r.RequestURI, "v1/")
	// fmt.Printf("\nuriParts: %v\n", uriParts)
	// uriParts1 := strings.Split(uriParts[1], "/")
	// resource = uriParts1[0]

	fmt.Printf("getResource:624 - fhirSystem to use : %s\n", spew.Sdump(fhirSystem))
	// params := mux.Vars(r)
	// fmt.Printf("getResource:274 - params  %v\n", params)
	// resourceId := params["id"]
	fmt.Printf("getResource:628 - Retrieving %s Record for id: [%s]\n", resource, resourceId)
	if resourceId == "" {
		//fmt.Printf(":180 %s with [%s]\n", resource, resourceId)
		err = fmt.Errorf("getResource:631  --  GetResource %s specific ID string is required", resource)
		errMsg := err.Error()
		fmt.Printf("%s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverity(fhir.IssueTypeInvalid), &errMsg))
	}
	//TODO: Handle Get Resource by specific ID.  All Resources including Binary.
	resp := common.ResourceResponse{}
	results, err := GetResource(fhirSystem, resource, resourceId)
	if err == nil {
		resp.Status = 200
		resp.Message = "Ok"
	} else {
		fmt.Printf("\n\nGetResource:643 -  returned err: %v\n\n\n", err)
		resp.Status = 400
		resp.Message = err.Error()
	}

	resp.ResourceType = resource
	resp.Resource.Resource = results
	fmt.Printf("getResource:650 - returning a %s resource\n", resource)
	WriteFhirResourceBundle(w, resp.Status, &resp)
}

// func getCachePage(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	//fmt.Printf("getCachePage:300 - %s \n", spew.Sdump(r))

// 	params := mux.Vars(r)
// 	fmt.Printf("getCachePage:484  -- %v\n", params)
// 	queryId := params["queryId"] // The id assigned to the query that created the cache
// 	pageNumber := params["pageNum"]
// 	//pageId := params["page_id"]
// 	fmt.Printf("Retrieving bundle for id: [%s]  Page: [%s]\n", queryId, pageNumber)
// 	FhirId := GetFhirId(r)
// 	fmt.Printf("getCachePage:490 - FhirKey - [%s]\n", FhirId)
// 	fhirSystem, err := GetFhirSystem(FhirId)
// 	if err != nil {
// 		fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
// 		err = fmt.Errorf("fhirSystem error:  %s", err.Error())
// 		errMsg := err.Error()
// 		fmt.Printf("getCachePage::496 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	fmt.Printf("getCachePage:500 -  %s/n", spew.Sdump(fhirSystem))
// 	if queryId == "" || pageNumber == "" {
// 		err = fmt.Errorf("GetCachedPage queryId: %s, pageNumber: %s -  error:  %s", queryId, pageNumber, "query_id and pageNumber are required")
// 		errMsg := err.Error()
// 		fmt.Printf("getCachePage:312 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	} else {
// 		fmt.Printf("Call GetCache for  queryId: %s  pageNumber: %s\n", queryId, pageNumber)
// 		pageId, err := strconv.Atoi(pageNumber)
// 		if err != nil {
// 			err = fmt.Errorf("PageNumber: [%s] is invalid %s", pageNumber, err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("getCachePage:321 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeInvalid, fhir.IssueSeverityFatal, &errMsg))

// 		}
// 		totalPages, bundle, header, err := GetCache(queryId, pageId)
// 		if err != nil {
// 			err = fmt.Errorf("GetCachePage QueryId: %s, page: %s -  error:  %s", queryId, pageNumber, err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("GetCachePage: 329 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		//parts := strings.Split(r.RequestURI, "v1/")
// 		resp := common.ResourceResponse{}
// 		resp.Bundle = bundle
// 		resp.ResourceType = header.ResourceType
// 		resp.BundleId = *bundle.Id
// 		resp.Status = 200
// 		resp.QueryId = header.QueryId
// 		resp.PageNumber = header.PageId
// 		resp.CountInPage = len(bundle.Entry)
// 		resp.TotalPages = totalPages
// 		resp.Header = header

// 		//host := common.GetKVData(GetConfig().Data, "cacheHost")
// 		//header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)
// 		header.CacheUrl = fmt.Sprintf("%s/Cache/%s/", fhirSystem.UcUrl, header.QueryId)
// 		//header.CacheUrl = fmt.Sprintf("%s/%sv1/Cache/%s/", r.Host, parts[0], header.QueryId)

// 		resp.Message = "Ok"
// 		fmt.Printf("getCachePage:351 - returning a cached %s bundle\n", header.ResourceType)
// 		WriteFhirResourceBundle(w, resp.Status, &resp)

// 	}
// }

// func getCacheStatus(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	fmt.Printf("getCacheStatus:335 - %s \n", spew.Sdump(r))
// 	params := mux.Vars(r)
// 	fmt.Printf("params:337 - %v\n", params)
// 	queryId := params["queryId"] // The id assigned to the query that created the cache
// 	fmt.Printf("Count how many pages of cache are in an ID\n")

// 	if queryId == "" {
// 		err = fmt.Errorf("GetCacheStatus queryId: %s -  error:  %s", queryId, "query_id is required")
// 		errMsg := err.Error()
// 		fmt.Printf("Handler:344 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	} else {
// 		fmt.Printf("Count Pages for queryId: %s\n", queryId)
// 		totalPages, err := TotalCacheForQuery(queryId)
// 		if err != nil {
// 			err = fmt.Errorf("GetCacheStatus QueryId: %s -  error:  %s", queryId, err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("GetCacheStatus: 353 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		parts := strings.Split(r.RequestURI, "v1/")
// 		resp := common.ResourceResponse{}

// 		//resp.ResourceType = header.ResourceType
// 		header := &common.CacheHeader{}
// 		resp.Status = 200
// 		resp.QueryId = queryId
// 		//resp.PageNumber = header.PageId
// 		//resp.CountInPage = len(bundle.Entry)
// 		resp.TotalPages = totalPages
// 		resp.Header = header
// 		resp.Header.QueryId = queryId
// 		host := common.GetKVData(GetConfig().Data, "cacheHost")
// 		resp.Header.CacheUrl = fmt.Sprintf("%s/%sv1/Cache/%s/", host, parts[0], header.QueryId)

// 		resp.Message = "Ok"
// 		//fmt.Printf("$$$:373 - returning  cached %s bundle\n", header.ResourceType)
// 		WriteFhirResourceBundle(w, resp.Status, &resp)
// 	}
// }

// func checkStatus(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	fmt.Printf("checkStatus:301 - %s \n", spew.Sdump(r))
// 	params := mux.Vars(r)
// 	fmt.Printf("params:303 - %v\n", params)
// 	queryId := params["queryId"] // The id assigned to the query that created the cache
// 	fmt.Printf("Count how many pages of cache are in an ID\n")

// 	if queryId == "" {
// 		err = fmt.Errorf("GetCacheStatus queryId: %s -  error:  %s", queryId, "query_id is required")
// 		errMsg := err.Error()
// 		fmt.Printf("Handler:344 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	} else {
// 		fmt.Printf("Count Pages for queryId: %s\n", queryId)
// 		totalPages, err := TotalCacheForQuery(queryId)
// 		if err != nil {
// 			err = fmt.Errorf("GetCacheStatus QueryId: %s -  error:  %s", queryId, err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("GetCacheStatus: 353 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		parts := strings.Split(r.RequestURI, "v1/")
// 		resp := common.ResourceResponse{}

// 		//resp.ResourceType = header.ResourceType
// 		header := &common.CacheHeader{}
// 		resp.Status = 200
// 		resp.QueryId = queryId
// 		//resp.PageNumber = header.PageId
// 		//resp.CountInPage = len(bundle.Entry)
// 		resp.TotalPages = totalPages
// 		resp.Header = header
// 		resp.Header.QueryId = queryId
// 		host := common.GetKVData(GetConfig().Data, "cacheHost")
// 		resp.Header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)

// 		resp.Message = "Ok"
// 		//fmt.Printf("$$$:373 - returning  cached %s bundle\n", header.ResourceType)
// 		WriteFhirResourceBundle(w, resp.Status, &resp)
// 	}
// }
