package services

import (
	//"bytes"
	"context"
	//"encoding/json"

	//"errors"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	//"github.com/dhf0820/fhir4"
	fhir "github.com/dhf0820/fhir4"
	//"github.com/samply/golang-fhir-models/fhir-models/fhir"
	common "github.com/dhf0820/uc_common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"io/ioutil"
	//"net/http"
	//"os"
	//"strings"
)

type Interface interface{}
type PostPatientPayload struct {
	MRN     string       `json:"mrn"`
	Patient fhir.Patient `json:"patient"`
}

//This is ChartArchive Fhir Interface to save a patient in Mongo.
//returns id, mrn, text
func SavePatient(mrn string, patient *fhir.Patient) (string, string, string, error) {
	fmt.Printf("SavePatient:34  --  patient: %s\n", spew.Sdump(patient))
	id := primitive.NewObjectID().Hex()
	ident := CreateIdentifier(id)
	fmt.Printf("SavePatient:38 --  New Identifier: %s\n", spew.Sdump(ident))
	patident := patient.Identifier
	patident = append(patident, ident)
	patient.Identifier = patident
	fmt.Printf("SavePatient:42 --  New Identifiers: %s\n", spew.Sdump(patient.Identifier))
	collection, err := GetCollection("Patients")
	if err != nil {
		return "", "", "", err
	}

	result, err := collection.InsertOne(context.TODO(), patient)
	if err != nil {
		err = fmt.Errorf("savePatient:45  --  insert Patient InsertOne failed: %v", err.Error())
		return "", "", "", err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		GetMrn(patient, "http://terminology.hl7.org/CodeSystem/v2-0203", "OurMrn")
		return oid.Hex(), *ident.Value, patient.Text.Div, err
	} else {
		err := fmt.Errorf("Invalid objectId")
		return "", "", "", err
	}
}

func CreateMRN(id string) string {
	return string(id[len(id)-6:])
}

func CreateIdentifier(id string) fhir.Identifier {
	layout := "2006-01-02T15:04:05.000Z"
	ident := fhir.Identifier{}
	ident.Id = StrPtr(primitive.NewObjectID().Hex())
	ident.Use = nil
	cc := fhir.CodeableConcept{}
	code := fhir.Coding{}
	code.System = StrPtr("https://fhir.vertisoft.com/6329112852f3616990e2f763/codeSet/4")
	//code.System = StrPtr("http://terminology.hl7.org/CodeSystem/v2-0203")
	code.Code = StrPtr("OurMrn")
	code.Display = StrPtr("Medical Record Number")
	code.UserSelected = BoolPtr(false)

	fmt.Printf("\nCreateIdentifier:72  --  ident : %s\n\n", spew.Sdump(ident))
	//coding := []fhir.
	cc.Coding = append(cc.Coding, code)
	ident.Type = &cc
	ident.Type.Text = StrPtr("OurMRN")
	ident.Value = StrPtr(CreateMRN(id))
	currentTime := time.Now()
	ident.Period = &fhir.Period{}
	ident.Period.Start = StrPtr(currentTime.Format(layout))
	fmt.Printf("\nCreateIdentifier:81  --  ident : %s\n\n", spew.Sdump(ident))
	return ident
}

//This is Generic Fhir Interface to save a patient

// func (c *Connection) SavePatient(mrn string, patient *fhir.Patient) (*fhir.Patient, error) {

// 	if mrn == "" { // For now use the provided MRN, if not there error //Generate a new MRN and insert into Identifiers.
// 		return nil, errors.New("new UNIQUE MRN for the patient must be specified")
// 	}
// 	if patient == nil {
// 		return nil, errors.New("FHIR (R4) patient must be provided")
// 	}
// 	patient.Id = StrPtr(primitive.NewObjectID().Hex())
// 	patient.Meta = &fhir.Meta{}
// 	patient.Meta.VersionId = StrPtr("1")
// 	patient.Meta.LastUpdated = StrPtr(time.Now().Format("2006-01-02T15:04:05 0000Z"))

// 	ident := fhir.Identifier{}
// 	id := primitive.NewObjectID().Hex()
// 	ident.Id = &id
// 	// idUse := fhir.IdentifierUse.Code(fhir.IdentifierUseUsual)
// 	// fhir.IdentifierUseUsual
// 	//idUse := fhir.IdentifierUseUsual
// 	code := fhir.IdentifierUseUsual
// 	ident.Use = &code
// 	ident.Value = &mrn
// 	ident.Type = &fhir.CodeableConcept{}
// 	ident.Type.Coding = []fhir.Coding{}
// 	coding := fhir.Coding{}
// 	coding.System = StrPtr("http://terminology.hl7.org/CodeSystem/v2-0203")
// 	coding.Code = StrPtr("MR")
// 	coding.Display = StrPtr("Medical record number")
// 	coding.UserSelected = BoolPtr(false)
// 	ident.Type.Coding = append(ident.Type.Coding, coding)
// 	ident.Type.Text = StrPtr("MRN")
// 	//ident.Period
// 	ident.System = StrPtr("http://terminology.hl7.org/CodeSystem/v2-0203") //TODO: Replace with our own.
// 	ident.Value = &mrn
// 	//TODO: add _value Extension  for Rendered Value
// 	patient.Identifier = []fhir.Identifier{}
// 	patient.Identifier = append(patient.Identifier, ident)
// 	fmt.Printf("\npatient: %s\n\n", spew.Sdump(patient))
// 	client := &http.Client{}
// 	fmt.Printf("Save Fhir Patient to: [%s]\n", fhirSystemURL)
// 	bstr, err := json.Marshal(patient)
// 	req, err := http.NewRequest("POST", fhirSystemURL, bytes.NewBuffer(bstr))
// 	if err != nil {
// 		fmt.Printf("NewRequest error: %s\n", err.Error())
// 	}
// 	req.Header.Set("Accept", "application/json+fhir")
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		log.Println("Error Posting new Patient:", err.Error())
// 		return nil, err
// 	}
// 	//fmt.Printf("length of ressponse Body = %d\n", len(resp.Body) )
// 	defer resp.Body.Close()
// 	fmt.Printf("resp.StatusCode = %d - %s\n", resp.StatusCode, resp.Status)
// 	// body, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Printf("Query Error: %v\n", err)
// 	// 	return nil, err
// 	// }

// 	//fmt.Printf("PostPatient response: %s\n", spew.Sdump(resp))
// 	return patient, nil
// }

func GetPatient(patId string) (*fhir.Patient, error) {
	fmt.Printf("GetPatient:152  -- retrieving a patient by id: %s\n", patId)

	filter := bson.D{{"id", patId}}
	collection, _ := GetCollection("Patients")
	pat := &fhir.Patient{}
	fmt.Printf("GetPatient:157  --  Calling FindOne with Filter: %v\n", filter)
	err := collection.FindOne(context.TODO(), filter).Decode(pat) // See if the user already has a session
	if err != nil {
		fmt.Printf("GetPatient:158  -- FindOne error: %s\n", err.Error())
		return nil, err
	}
	//fmt.Printf("GetPatient:158  -- FindOne Patient: %s\n", spew.Sdump(pat))
	return pat, err

	// qry := fmt.Sprintf("Patient/%s", patId)
	// log.Infof("Final url to query: %s\n", qry)
	// startTime := time.Now()
	// bytes, err := c.Query(qry)
	// log.Infof("Query time: %s", time.Since(startTime))

	// if err != nil {
	// 	return nil, fmt.Errorf("Query %s failed: %s", qry, err.Error())
	// }
	// patient := fhir.Patient{}
	// err = json.Unmarshal(bytes, &patient)
	// if err != nil {
	// 	return nil, err
	// }
	// return &patient, err
}

func PatientSearch(fhirSystem *common.FhirSystem, query, resource, token string) (*fhir.Bundle, error) {
	// fhirID, err := primitive.ObjectIDFromHex(fhirId)
	// if err != nil {
	// 	return nil, err
	// }
	log.Infof("queryString: %s\n", query)
	qry := fmt.Sprintf("Patient?%s", query)
	log.Infof("Final url to query: %s\n", qry)
	// startTime := time.Now()
	// //b, err := c.Query(fmt.Sprintf("/Patient?%s", query))
	// bytes, err := c.Query(qry)
	// log.Infof("Query time: %s", time.Since(startTime))
	bundle := &fhir.Bundle{}
	/*
		if err != nil {

			return nil, fmt.Errorf("Query %s failed: %s", query, err.Error())
		}

		//fmt.Printf("\n\n\n@@@ RAW Patient: %s\n\n\n", pretty.Pretty(b))
		// prettyJSON, err := json.MarshalIndent(b, "", "    ")
		// if err != nil {
		// 	fmt.Printf("MarshalIndent failed: %s\n", err.Error())
		// 	return nil, err
		// }

		startTime = time.Now()
		bundle := &fhir.Bundle{}
		//data := PatientResult{}
		if err := json.Unmarshal(bytes, &bundle); err != nil {
			return nil, fmt.Errorf("PatientSearch ummarshal : %s", err.Error())
		}
		log.Infof("Unmarshal time: %s", time.Since(startTime))
		//fmt.Printf("Response: %s\n", spew.Sdump(bundle))
		//resourceCache := common.ResourceCache

		for _, entry := range bundle.Entry {
			resourceCache := common.ResourceCache{}
			resourceJson := entry.Resource
			patient := fhir.Patient{}
			json.Unmarshal(resourceJson, &patient)
			resourceCache.Resource = entry.Resource
			resourceCache.ResourceType = "Patient"
			fmt.Printf("PatientSearch:160  --  PatientId = %s\n", *patient.Id)

		}
		header := &common.CacheHeader{}
		header.FhirSystem = fhirSystem
		cacheBundle := common.CacheBundle{}
		cacheBundle.ID = primitive.NewObjectID() //Each cach bundle gets a new header. The queryId ties all pages together.

		header.FhirId = fhirSystem.ID.Hex()            // Uniquely identifies the real url fo the fhir server
		header.QueryId = primitive.NewObjectID().Hex() //Does not change on each page
		header.PatientId = ""                          // Not used for patient cache sine each entry is a different patient
		header.ResourceType = "Patient"
		tn := time.Now()
		header.CreatedAt = &tn

		cacheBundle.Header = header
		cacheBundle.Bundle = bundle
		cacheBundle.Header.PageId = 1

		//TODO: Call Core CacheResources to cachhe the resources(patients)
		fmt.Printf("PatientSearch:179 calling Insert %d Patients for now\n", len(cacheBundle.Bundle.Entry))

		err = Insert(context.Background(), &cacheBundle, token)
		if err != nil {
			msg := fmt.Sprintf("CacheInsert initial error %s", err.Error())
			fmt.Println(msg)
			log.Error(msg)
			return nil, errors.New(msg)
		}
		nextURL := GetNextResourceUrl(bundle.Link)
		if nextURL == "" {
			msg := fmt.Sprintf("GetNextResourceUrl initial No Next ")
			// fmt.Println(msg)
			log.Warn(msg)
			//return nil, errors.New(msg)
			return bundle, nil
		}
		go c.GetNextResource(header, nextURL, resource, token)
	*/
	return bundle, nil

}

// func GetNextResourceUrl(link []fhir.BundleLink) string {
// 	for _, lnk := range link {
// 		if lnk.Relation == "next" {
// 			return lnk.Url
// 		}
// 	}
// 	return ""
// }
// func (c *Connection) GetNextResource(header *common.CacheHeader, url, token string) {
// 	startTime := time.Now()
// 	bytes, err := c.GetFhir(url)
// 	fmt.Printf("Query Next Set time: %s\n", time.Since(startTime))
// 	if err != nil {
// 		msg := fmt.Sprintf("c.GetFhir error: %s", err.Error())
// 		fmt.Println(msg)
// 		log.Error(msg)
// 		return
// 	}
// 	bundle := &fhir.Bundle{}

// 	if err := json.Unmarshal(bytes, bundle); err != nil {
// 		msg := fmt.Sprintf("PatientSearch next unmarshal : %s", err.Error())
// 		log.Error(msg)
// 		fmt.Println(msg)
// 		return
// 	}
// 	header.PageId += 1
// 	tn := time.Now()
// 	header.CreatedAt = &tn
// 	cacheBundle := common.CacheBundle{}
// 	cacheBundle.ID = primitive.NewObjectID()
// 	cacheBundle.Header = header
// 	cacheBundle.Bundle = bundle

// 	err = Insert(context.Background(), &cacheBundle)
// 	if err != nil {
// 		msg := fmt.Sprintf("CacheInsert error %s", err.Error())
// 		fmt.Println(msg)
// 		log.Error(msg)
// 	}
// 	fmt.Printf("Link: %s\n", spew.Sdump(bundle.Link))
// 	nextURL := GetNextResourceUrl(bundle.Link)
// 	if nextURL == "" {
// 		msg := fmt.Sprintf("GetNextResourceUrl Last page had %d Resources processed ", len(bundle.Entry))
// 		// fmt.Println(msg)
// 		log.Warn(msg)
// 		fmt.Printf("GetNext Resources should return\n")
// 		return
// 	} else {
// 		fmt.Printf("GetNextResources is being called in the background\n")
// 		go c.GetNextResources(header, nextURL, token)
// 		fmt.Printf("GetNextResources was called in the background\n")
// 	}
// 	fmt.Printf("GetNext Resource is returning\n")
// 	return
// }

func GetMrn(pat *fhir.Patient, system string, code string) string {
	idents := pat.Identifier
	for _, ident := range idents {
		if *ident.Type.Text == code {
			value := *ident.Value
			fmt.Printf("GetMrn:337  --  MRN Code : %s = %s\n", code, value)
			return value
		}
	}
	return ""
}
