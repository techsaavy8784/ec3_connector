package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	//"github.com/dhf0820/fhir4"
	fhir "github.com/dhf0820/fhir4"
	token "github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"

	//"github.com/gorilla/mux"
	"log"
	"net/http"

	//"os"
	//"strconv"
	"strings"
)

//####################################### Response Writers Functions #######################################
//################################### FHIR Responses ####################################
//####################################### Route Handlers #######################################

func getDocRef(w http.ResponseWriter, r *http.Request) {
	Resource := "DocumentReference"
	fmt.Printf("getDocRef:29 - Request: %s \n", spew.Sdump(r))

	//buildFieldsByTagMap("schema", *psp)

	// Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	// if err != nil {
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// userId := Payload.UserId
	// fhirId := GetFhirId(r)
	// fhirSystem, err := GetFhirSystem(fhirId)
	// if err != nil {
	// 	log.Printf("searchPatient:50  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// 	err = errors.New("invalid FHIR URL")
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }

	Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	userId := Payload.UserId
	log.Printf("getDocRef:59  --  UserId: %s\n", userId)

	fhirId := GetFhirId(r)
	fhirSystem, err := GetFhirSystem(fhirId)
	if err != nil {
		log.Printf("getDocRef:63  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
		err = errors.New("invalid FHIR URL")
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	urlA, err := r.URL.Parse(r.RequestURI)
	if err != nil {
		err = fmt.Errorf("error parsing DocRef URI: %s", err.Error())
		errMsg := err.Error()
		fmt.Printf("getDocRef:73 - r.URL.Parse error = %s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("getDocRef:77 - r.URL.Parse = %v\n", urlA)
	urlB := *urlA
	uriValues := urlB.Query()
	fmt.Printf("getDocRef:80 - uriValues= %v\n", uriValues)

	uri := r.RequestURI
	log.Printf("uri = %s\n", uri)
	parts := strings.Split(uri, Resource)
	uri = parts[1]
	log.Printf("getDocRef:86 - URI = %s\n", uri)
	//patient := fhir.Patient{}
	resource, err := GetResource(fhirSystem, Resource, uri)
	resp := common.ResourceResponse{}
	if err != nil {
		resp.Status = 400
		resp.Message = err.Error()
	} else {
		resp.Status = 200
		resp.Message = "Ok"
	}
	// var docRef fhir.DocumentReference
	// docRef = resource.(fhir.DocumentReference)
	resp.Resource.Resource = resource
	// var res []interface{}
	// res = append(res, &resource)
	// resp.Resources = res
	resp.ResourceType = Resource
	//resp.ResourceId = *docRef.Id
	log.Printf("\nGetDocRef:105  --  resp: %s\n", spew.Sdump(resp))
	WriteFhirResourceBundle(w, resp.Status, &resp)
}

// searchPatient uses the fhirId url parameter to determin the FhirSystem to use
func searchDocRef(w http.ResponseWriter, r *http.Request) {
	Resource := "DocumentReference"
	body, err := ioutil.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		fmt.Printf("findResource:354  --  ReadAll FhirSystem error %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	connectorPayload := common.ConnectorPayload{}
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		err = fmt.Errorf("searchDocRef:121  --  unmarshal err = %s", err.Error())
		fmt.Println(err)
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fhirSystem := connectorPayload.FhirSystem
	connConfig := connectorPayload.ConnectorConfig
	JWToken := r.Header.Get("Authorization")
	Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	userId := Payload.UserId
	log.Printf("searchDocRef:119  --  UserId: %s\n", userId)
	// fhirId := GetFhirId(r)
	// fhirSystem, err := GetFhirSystem(fhirId)
	// if err != nil {
	// 	log.Printf("searchDocRef:123  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// 	err = errors.New("invalid FHIR URL")
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	uri := r.RequestURI
	log.Printf("uri = %s\n", uri)
	parts := strings.Split(uri, Resource)
	uri = parts[1]
	log.Printf("\nsearchDocRef:151 - URI = %s\n", uri)

	urlA, err := r.URL.Parse(r.RequestURI)
	if err != nil {
		err = fmt.Errorf("error parsing DocRef URI: %s", err.Error())
		errMsg := err.Error()
		fmt.Printf("searchDocRef:157 - r.URL.Parse error = %s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("searchDocRef:161 - r.URL.Parse = %v\n", urlA)
	urlB := *urlA
	uriValues := urlB.Query()
	fmt.Printf("searchDocRef:164 - uriValues= %v\n", uriValues)

	log.Printf("\n\nResource Is DocumentReference\n\n")
	//urlA, err := r.URL.Parse(r.RequestURI)
	if err != nil {
		err = fmt.Errorf("error parsing patient URI: %s", err.Error())
		errMsg := err.Error()
		fmt.Printf("searchDocRef:171 - r.URL.Parse error = %s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("searchDocRef:175 - r.URL.Parse = %v\n", urlA)
	//urlB := *urlA
	//uriValues := urlB.Query()
	fmt.Printf("searchDocRef:160 - uriValues= %v\n", uriValues)
	idSearch := uriValues.Get("identifier")
	idValue := ""
	if idSearch != "" { // There is identifier Search, use it
		fmt.Printf("searchDocRef:182 - using Identifier: %s to search\n", idSearch)
		ids := strings.Split(idSearch, "|")
		if len(ids) != 2 {
			err = fmt.Errorf("invalid identifier: %s", idSearch)
			errMsg := err.Error()
			fmt.Printf("searchDocRef:187 - r.URL.Parse error = %s\n", errMsg)
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
		//
		if idValue == "" { //Not configured identifier
			err = fmt.Errorf("identifier type: %s is not configured", idName)
			errMsg := err.Error()
			fmt.Printf("searchDocRef:204 - Identifiers = %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		uri = fmt.Sprintf("?identifier=%s", idValue+idSearchValue)
		fmt.Printf("searchDocRef:209 - New Identifier search Value: %s\n", uri)
	} else {
		fmt.Printf("searchDocRef:211 - using other search params: %v\n", uriValues)
	}
	var bundle *fhir.Bundle
	var header *common.CacheHeader
	fmt.Printf("\nsearchDocRef:215 - resource = %s  uri = %s\n", Resource, uri)
	url := fmt.Sprintf("%s/%s%s", fhirSystem.FhirUrl, Resource, uri) //" + "/" + uri
	var totalPages int64
	uri = "/" + Resource + uri
	fmt.Printf("searchDorRef:219 FindResource %s\n", uri)
	totalPages, bundle, header, err = FindResource(&connectorPayload, Resource, userId, uri, JWToken)
	if err != nil {
		err = fmt.Errorf("searchDocRef:205 --  fhirSearch url: %s error:  %s", url, err.Error())
		errMsg := err.Error()
		fmt.Printf("searchDocRef:1207 - %s\n", errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
		return
	}
	if bundle == nil {
		log.Printf("searchDocRef:212  --  bundle is nil")
	} else {
		log.Printf("searchDocRef:214  --  bundle is not nil \n")
	}

	fmt.Printf("searchDocRef:216 - Get %s bundle successful\n", Resource)
	fmt.Printf("searchDocRef:217 - Number in page: %d\n", len(bundle.Entry))
	fmt.Printf("searchDocRef:218 - PageNumber: %d\n", header.PageId)
	resp := common.ResourceResponse{}
	cacheBase := fmt.Sprintf("%s/%s", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
	cacheBundleURL := cacheBase + "/BundleTransaction"
	log.Printf("\n\n\n\n\n$$$ $$$ searchDocRef:214  --  CacheBundleUrl = %s\n", cacheBundleURL)
	header.FhirId = fhirSystem.ID.String()
	header.UserId = userId
	resp.Bundle = bundle
	resp.Resource.Resource = bundle.Entry[0].Resource //header.ResourceType
	resp.BundleId = *bundle.Id
	resp.ResourceType = Resource
	resp.Status = 200
	resp.QueryId = header.QueryId
	resp.PageNumber = header.PageId
	resp.CountInPage = len(bundle.Entry)
	resp.TotalPages = totalPages
	resp.Header = header
	resp.Message = "Ok"
	//fmt.Printf("searchPatient:228 - returning a resource bundle: %s\n", spew.Sdump(resp))
	WriteFhirResourceBundle(w, resp.Status, &resp)
	//WriteFhirBundle(w, resp.Status, bundle)

}
