package common

import (
	//"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/schema"
	"github.com/oleiade/reflections"
	"net/http"
	//"reflect"
	//"strconv"
	"strings"
)

type PatientSearchParams struct {
	Count      SearchParam   `json:"count" schema:"_count"`
	OffSet     SearchParam   `json:"offset" schema:"_offset"`
	Order      SearchParam   `json:"order" schema:"_order"`
	Sort       SearchParam   `json:"sort" schema:"_sort"`
	Page       SearchParam   `json:"page" schema:"_page"`
	Id         SearchParam   `json:"id" schema:"id"`
	MRN        SearchParam   `json:"mrn" schema:"mrn"`
	SSN        SearchParam   `json:"ssn" schema:"ssn"`
	Identifier SearchParam   `json:"identifier" schema:"identifier"`
	Gender     SearchParam   `json:"gender" schema:"gender"`
	BirthDate  SearchParam   `json:"birthdate" schema:"birthdate"`
	Family     SearchParam   `json:"family" schema:"family"`
	Given      SearchParam   `json:"given" schema:"given"`
	Facility   SearchParam   `json:"facility" schema:"facility"`
	Active     SearchParam   `json:"active" schema:"active"`
	DOB        []SearchParam `json:"dob" schema:"dob"`
	BaseUrl    string        `json:"base_url"`
	RequestURI string        `json:"request_uri"`
	Limit      uint32        `json:"limit"`
	Skip       uint32        `json:"skip"`
}

type PatientFhirSearchParams struct {
	Count      SearchParam `json:"count" schema:"_count"`
	OffSet     SearchParam `json:"offset" schema:"_offset"`
	Order      SearchParam `json:"order" schema:"_order"`
	Sort       SearchParam `json:"sort" schema:"_sort"`
	Page       SearchParam `json:"page" schema:"_page"`
	Id         SearchParam `json:"id" schema:"_id"`
	MRN        SearchParam `json:"mrn" schema:"mrn"`
	SSN        SearchParam `json:"ssn" schema:"ssn"`
	Identifier SearchParam `json:"identifier" schema:"identifier"`
	Gender     SearchParam `json:"gender" schema:"gender"`
	BirthDate  SearchParam `json:"birthdate" schema:"birthdate"`
	Name       SearchParam `json:"name" schema:"name"`
	Family     SearchParam `json:"family" schema:"family"`
	Given      SearchParam `json:"given" schema:"given"`
	Phone      SearchParam `json:"phone" schema:"phone"`
	Email      SearchParam `json:"email" schema:"email"`
	PostalCode SearchParam `json:"address-postalcode" schema:"address-postalcode"`

	Facility   SearchParam   `json:"facility" schema:"facility"`
	Active     SearchParam   `json:"active" schema:"active"`
	DOB        []SearchParam `json:"dob" schema:"dob"`
	BaseUrl    string        `json:"base_url"`
	RequestURI string        `json:"request_uri"`
	Limit      uint32        `json:"limit"`
	Skip       uint32        `json:"skip"`
}

type DocumentSearchParams struct {
	Count      SearchParam   `json:"count" schema:"_count"`
	OffSet     SearchParam   `json:"offset" schema:"_offset"`
	Order      SearchParam   `json:"order" schema:"_order"`
	Sort       SearchParam   `json:"sort" schema:"_sort"`
	Page       SearchParam   `json:"page" schema:"_page"`
	Id         SearchParam   `json:"id" schema:"id"`
	Patient    SearchParam   `json:"patient" schema:"patient"`
	Subject    SearchParam   `json:"subject" schema:"subject"`
	Encounter  SearchParam   `json:"encounter" schema:"encounter"`
	Created    []SearchParam `json:"created" schema:"created"` // need to handle two of these
	Facility   SearchParam   `json:"facility" schema:"facility"`
	BaseUrl    string        `json:"base_url"`
	RequestURI string        `json:"request_uri"`
	Limit      uint32        `json:"limit"`
	Skip       uint32        `json:"skip"`
}

type SearchParam struct {
	Schema   string
	Modifier string
	Value    string
}

var decoder = schema.NewDecoder()

func FhirPatientSearch(r *http.Request) (*PatientFhirSearchParams, error) {
	err := r.ParseForm()
	if err != nil {
		// Handle error
	}

	var pat PatientFhirSearchParams

	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&pat, r.PostForm)
	if err != nil {
		return nil, err
	}
	fmt.Printf("PatFilter: %s\n", spew.Sdump(pat))
	return &pat, err
}

func patientSearchParams(r *http.Request) (*PatientSearchParams, error) {
	var pspTags map[string]string
	tagFields := make(map[string]string)

	//buildFieldsByTagMap("schema", *psp)
	//facility := "demo"
	fmt.Printf("searchPatient called\n")
	if err := r.ParseForm(); err != nil {
		err = fmt.Errorf("Error parsing query: %s", err.Error())
		return nil, err
	}
	psp := new(PatientSearchParams)

	//TODO: Include the facility as part of the base url not as a parameter
	fmt.Printf("tls: %v\n", r.TLS)
	protocol := "http://"
	psp.BaseUrl = fmt.Sprintf("%s%s/api/rest/v1", protocol, r.Host)
	psp.RequestURI = r.RequestURI
	psp.Facility.Value = strings.Trim(psp.Facility.Value, " ")
	if psp.Facility.Value == "" {
		err := fmt.Errorf("Faciity is required")
		return nil, err
	}

	pspTags, _ = reflections.Tags(psp, "schema")
	for k, v := range pspTags {
		//	fmt.Printf("key: = %s;  value: %s\n", k, v)
		tagFields[v] = k
	}
	fmt.Printf("\ntagFields: %s\n\n", spew.Sdump(tagFields))

	// var decoder = schema.NewDecoder()
	// decoder.IgnoreUnknownKeys(true)
	fmt.Printf("query: %s\n", r.URL.RawQuery)
	qryParams := strings.Split(r.URL.RawQuery, "&")
	for _, param := range qryParams {
		fmt.Printf("patientSearchParam:98  -- param: %s\n", param)
		//keyValue := strings.Split(parm, "=") //split key and value
		//value := keyValue[1]
		//fmt.Printf("parts : %v\n", keyValue[0] )  // key is elem 0 value is elem 1
		//keyMod := strings.Split(keyValue[0], ":") // separate the modifier from the key if any
		//key := keyMod[0]
		//mod := ""
		// if len(keyMod) == 2 { //There is a modifier
		// 	mod = keyMod[1]
		// }
		// param := db.SearchParam{}
		// param.Modifier = mod
		// param.Value = strings.Trim(value, " ")
		// param.Schema = key
		//fmt.Printf("Key: %s,  mod: %s  value: %s  Field: %s\n", key, mod, value, spew.Sdump(param))
		//fieldName := tagFields[key]
		//fmt.Printf("Setting Field Data\n")
		//err = reflections.SetField(psp, fieldName, param)
	}

	if psp.Count.Value == "" { // set ount to default 20
		psp.Count.Value = "20"
		psp.Count.Schema = "_count"
	}
	// count, err := strconv.ParseUint(psp.Count.Value, 10, 32)
	// if err != nil {
	// 	err = fmt.Errorf("invalid _count Err: %s", err.Error())
	// 	return nil, err
	// }
	return psp, nil
}
