package services

import (
	//"context"
	"encoding/json"

	//"encoding/json"
	//"errors"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	fhir4 "github.com/dhf0820/fhir4"
	common "github.com/dhf0820/uc_common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Request a specific resource by id
func GetResource(fhirSystem *common.FhirSystem, resourceName, resourceId string) (json.RawMessage, error) {
	//startTime := time.Now()
	fmt.Printf("GetResource:23 - fhirConnector: %s\n\n", spew.Sdump(fhirSystem))
	url := fmt.Sprintf("%s/%s%s", fhirSystem.FhirUrl, resourceName, resourceId)
	fmt.Printf("GetResource: 25 final Query: %s\n", url)
	c := New(fhirSystem.FhirUrl)

	bytes, err := c.GetFhir(url)
	if err != nil {
		return nil, err
	}
	err = nil

	return bytes, nil
	// var res interface{}
	// //var binary *fhir4.Binary
	// //var vsResource *fhir4.VsResource
	// switch resourceName {
	// case "Binary":
	// 	res, err = fhir4.UnmarshalBinary(bytes)
	// case "Condition":
	// 	res, err = fhir4.UnmarshalCondition(bytes)
	// case "DocumentReference":
	// 	res, err = fhir4.UnmarshalDocumentReference(bytes)
	// 	if err != nil {
	// 		fmt.Printf("UnmarshalDocumentReference direct to DocumentReference err = %v\n", err)
	// 	}
	// case "Observation":
	// 	res, err = fhir4.UnmarshalObservation(bytes)
	// case "Patient":
	// 	fmt.Printf("Returning ByteArray")
	// 	return bytes, nil
	// 	pat, err := fhir4.UnmarshalPatient(bytes)
	// 	if err != nil {
	// 		fmt.Printf("UnmarshalPatient direct to patient err = %v\n", err)
	// 	}
	// 	bytes, err := json.Marshal(&pat)
	// 	if err != nil {
	// 		fmt.Printf("MarshalPatient to RawMessage err = %v\n", err)
	// 	}
	// 	rm := string(bytes)
	// 	fmt.Println(rm)

	// default:
	// 	err := fmt.Errorf("unknown type of Resource: [%s]", resourceName)
	// 	return nil, err
	// }
	// //fmt.Printf("GetResource is returning err: %v\n", err)
	// fmt.Printf("  QueryTime: %s\n", time.Since(startTime))
	// return bytes, err
}

//Search for Resources matching url filters or id
func FindResource(connPayLoad *common.ConnectorPayload, resource, userId, query, token string) (int64, *fhir4.Bundle, *common.CacheHeader, error) {
	page := 1
	connConfig := connPayLoad.ConnectorConfig
	fmt.Printf("FindResource:77 -- connConfig = %s\n", spew.Sdump(connConfig))
	fhirSystem := connPayLoad.FhirSystem
	fmt.Printf("\n\n@@@@@@FindResource:79  --  Resource: %s\n", resource)
	fmt.Printf("FindResource:80  --  UserId: %s\n", userId)
	fmt.Printf("FindResource:81  --  Query: %s\n", query)
	fmt.Printf("FindResource:82  --  fhirSystemId: %s\n", fhirSystem.ID.Hex())
	fmt.Printf("FindResource:83  --  FhirSystem: %s\n", spew.Sdump(fhirSystem))
	fmt.Printf("FindResource:84  -- Page: %d\n", page)
	fmt.Printf("FindResource:85  --  ConnectorConfig: %s\n", spew.Sdump(connConfig))
	// NOT SURE THIS TODO NEEDS TO BE DONE. FindResource uses bundle and returns after the first bundle is returned
	//
	//TODO: Process the query in the background filling the resourceCache and BundleCache. Assign a cacheId on the call
	//Once background is started wait in a loop checking the ResourceCache Status using the assigned cacheId until either
	// Have count documents or status is finished.
	// check every 10 seconds.  Should be a FhirSystem variable value to avoid code change
	c := New(fhirSystem.FhirUrl)
	startTime := time.Now()
	byte, err := c.Query(query) // Perform the actul query of the fhir server
	if err != nil {
		return 0, nil, nil, err
	}
	header := &common.CacheHeader{}
	header.FhirSystem = fhirSystem
	header.ResourceType = resource
	header.UserId = userId
	header.PageId = page
	queryId := primitive.NewObjectID().Hex()
	header.QueryId = queryId
	fmt.Printf("FindResource:105  --  connConfig: %s\n", spew.Sdump(connConfig))
	header.CacheBase = fmt.Sprintf("%s/%s", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
	//header.ResourceCacheBase = fmt.Sprintf("%s/%s/%s/BundleTransaction", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
	header.GetBundleCacheBase = fmt.Sprintf("%s/%s/BundleTransaction", header.CacheBase, header.FhirSystem.ID.Hex())
	header.GetResourceCacheBase = fmt.Sprintf("%s/%s/CachePage", header.CacheBase, header.FhirSystem.ID.Hex())

	cacheBundle := common.CacheBundle{}
	cacheBundle.PageId = header.PageId
	cacheBundle.Header = header
	cacheBundle.ID = primitive.NewObjectID()
	//fmt.Printf("\n\n\n\n$$$ FindResource:110 calling CacheResourceBundleAndEntries (without bundle) - %s \n", spew.Sdump(cacheBundle))

	//Cache the first bundle(page)
	fmt.Printf("\n\n FindResource:118  --  Query %s for %ss took %s\n\n\n\n\n\n\n\n\n", fhirSystem.DisplayName, resource, time.Since(startTime))
	fmt.Printf("FindResource:119 --  UnmarshalBundle\n\n")
	bundle := fhir4.Bundle{}
	bundle, err = fhir4.UnmarshalBundle(byte)
	if err != nil {
		return 0, nil, nil, err
	}
	cacheBundle.Bundle = &bundle
	startTime = time.Now()
	pg, err := CacheResourceBundleAndEntries(&cacheBundle, token, page)
	fmt.Printf("FindResource:127 CacheResource returned %d %ss in page: %d for %s  took %s\n", len(cacheBundle.Bundle.Entry), resource, page, fhirSystem.DisplayName, time.Since(startTime))
	if err != nil {
		//return err and done
		return int64(pg + 1), &bundle, cacheBundle.Header, err
	}
	fmt.Printf("FindResource:133  --  links: %s\n", spew.Sdump(bundle.Link))
	//Follow the bundle links to retrieve all bundles(pages) in the query response
	nextURL := GetNextResourceUrl(bundle.Link)
	if nextURL == "" {
		msg := fmt.Sprintf("FindResource:132 -- GetNext%sUrl initialy No Next - One page only ", resource)
		fmt.Println(msg)
		_, err := TotalCacheForQuery(cacheBundle.QueryId)
		cacheBundle.Header.PageId = pg
		//page++
		return int64(pg), &bundle, cacheBundle.Header, err
	}
	page++
	go c.GetNextResource(header, nextURL, resource, token, page)
	fmt.Printf("FindResource:146  --  Page 1 total time: %s\n", time.Since(startTime))
	// There is one full page and possibley more. Respond with two aso they user will create two page buttons and update every
	// 10 secnds.
	return int64(page), &bundle, cacheBundle.Header, err
}

func GetNextResourceUrl(link []fhir4.BundleLink) string {
	for _, lnk := range link {
		if lnk.Relation == "next" {
			fmt.Printf("$$$$  GetNextResourceUrl:146  --  There is  next page to get\n")
			return lnk.Url
		}
	}
	return ""
}

//GetNextResource: fetches the resource at provided url, processes it and checks if more to call.
func (c *Connection) GetNextResource(header *common.CacheHeader, url, resource, token string, page int) {
	fmt.Printf("\n\n\n\n####################  GetNextResource page: %d   ###############\n", page)
	//fmt.Printf("GetNextResource:155  --  resource: %s\n", resource) //spew.Sdump(header))
	//Call Remote FHIR server for the resource bundle
	startTime := time.Now()
	bytes, err := c.GetFhir(url)
	if err != nil {
		msg := fmt.Sprintf("GetNextResource:170  --  error: %s", err.Error())
		fmt.Println(msg)
		return
	}
	fmt.Printf("GetNextResource:174  --  Query Next Set from %s of %s time: %s\n", header.FhirSystem.DisplayName, header.ResourceType, time.Since(startTime))
	fmt.Printf("GetNextResource:175  --  UnmarshalBundle\n")
	bundle, err := fhir4.UnmarshalBundle(bytes)
	if err != nil {
		msg := fmt.Sprintf("GetNextResource:178 unmarshal : %s", err.Error())
		//fmt.Printf(msg)
		fmt.Println(msg)
		return
	}

	//unMarshalResource(resource, bundle)
	header.PageId += 1
	tn := time.Now()
	header.CreatedAt = &tn
	cacheBundle := common.CacheBundle{}
	cacheBundle.ID = primitive.NewObjectID()
	cacheBundle.Header = header
	cacheBundle.Bundle = &bundle
	fmt.Printf("GetNextResource:192  -- Calling CacheResourceBundleAndEntries\n")
	pg, err := CacheResourceBundleAndEntries(&cacheBundle, token, page)
	if err != nil {
		log.Errorf("GetNextResource:195 returned err; %s\n", err.Error())
		return
		//return int64(pg + 1), &bundle, cacheBundle.Header, err
	}

	fmt.Printf("GetNextResource:200  --  GetNextResourceUrl\n")
	nextURL := GetNextResourceUrl(bundle.Link)
	if nextURL == "" {
		msg := fmt.Sprintf("GetNextResourceUrl Last page had %d Resources processed ", len(bundle.Entry))
		// fmt.Println(msg)
		fmt.Printf("GetNextResource:205 --  Should return --  %s\n", msg)
		return
	} else {
		fmt.Printf("GetNextResource:208 -- is being called in the background\n")
		go c.GetNextResource(header, nextURL, resource, token, pg+1)
		fmt.Printf("GetNextResource:210 -- Returned\n")
	}
	fmt.Printf("GetNextResource:203 is returning\n")
}

func GetHeaderInfoFromBundle(resource string, hdr *common.CacheHeader, bundle *fhir4.Bundle) (string, string, error) {
	fmt.Printf("\n\n\n################## GetHederInfoFromBundle:208  --  for Resource: [%s]\n\n", resource)
	//resHeader := common.ResourceHeader{}
	switch resource {
	case "Patient":
		fsIdentifiers := hdr.FhirSystem.Identifiers
		res, err := fhir4.UnmarshalPatient(bundle.Entry[0].Resource)
		if err != nil {
			log.Errorf("%s unmarshal : %s", resource, err.Error())
			return "", "", err
		}
		hdr.PatientId = *res.Id
		hdr.ResourceId = *res.Id
		GetPatientIdentifier(res.Identifier, fsIdentifiers, "mrn")
		fmt.Printf("\nGetHeaderInfoFromBundle:221 -- Patient: %s\n", spew.Sdump((res)))
		fmt.Printf("\nGetHeaderInfoFromBundle:222")
		res1, err := fhir4.UnmarshalPatient(bundle.Entry[1].Resource)
		if err != nil {
			log.Errorf("%s unmarshal : %s", resource, err.Error())
			return "", "", err
		}
		hdr.PatientId = *res.Id
		hdr.ResourceId = *res.Id
		fmt.Printf("\nGetHeaderInfoFromBundle:230 -- Patient: %s\n", spew.Sdump((res1)))
		return *res.Id, *res.Id, nil
	case "DocumentReference":
		//log.Printf("GetHederInfoFromBundle312: --  Raw data: %s\n", string(bundle.Entry[0].Resource))
		res, err := fhir4.UnmarshalDocumentReference(bundle.Entry[0].Resource)
		if err != nil {
			log.Errorf("%s unmarshal : %s", resource, err.Error())
			return "", "", err
		}
		//fmt.Printf("\n###DocumentReference: %s\n", spew.Sdump(res))
		parts := strings.Split(*res.Subject.Reference, "/")
		hdr.PatientId = parts[1]
		hdr.ResourceId = *res.Id
		fmt.Printf("GetHeaderInfoFromBundle:243 -- DocumentReference: %s\n", spew.Sdump((res)))
		return parts[1], *res.Id, nil
	case "DiagnosticReport":
		res, err := fhir4.UnmarshalDiagnosticReport(bundle.Entry[0].Resource)
		if err != nil {
			log.Errorf("%s unmarshal : %s", resource, err.Error())
			return "", "", err
		}
		//fmt.Printf("\n###DiagnosticReport: %s\n", spew.Sdump(res))
		parts := strings.Split(*res.Subject.Reference, "/")
		hdr.PatientId = parts[1]
		hdr.ResourceId = *res.Id
		fmt.Printf("GetHeaderInfoFromBundle:255 -- DiagnosticReport: %s\n", spew.Sdump((res)))
		return parts[1], "", nil
	case "Observation":
		res, err := fhir4.UnmarshalObservation(bundle.Entry[0].Resource)
		if err != nil {
			log.Errorf("%s unmarshal : %s", resource, err.Error())
			return "", "", err
		}
		//fmt.Printf("\n###Observation: %s\n", spew.Sdump(res))
		parts := strings.Split(*res.Subject.Reference, "/")
		hdr.PatientId = parts[1]
		hdr.ResourceId = *res.Id
		fmt.Printf("GetHeaderInfoFromBundle:267 -- Observation:%s\n", spew.Sdump((res)))
		return parts[1], "", nil
	}
	return "", "", nil
}

func GetCacheStatus(ucUrl, queryId string) int {
	// coreURL := ucUrl + "/BundleTransaction"
	// client := &http.Client{}
	// //fmt.Printf("Send Status to: [%s]\n", statusURL)
	// fmt.Printf("CacheResourceBundleAndEntries:99  --  Using CoreUrl: %s\n", coreURL)
	// req, _ := http.NewRequest("GET", coreURL, bytes.NewBuffer(cacheBundle))
	// //req, _ := http.NewRequest("POST", coreURL, bytes.NewBuffer(cacheBundle))
	// //r, _ := http.NewRequest("POST", coreURL, nil)
	// fmt.Printf("\nCacheResourceBundleEntries:339  --  Req: %s\n\n\n", spew.Sdump(req))
	// req.Header.Set("Accept", "application/json")
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", token)
	// fmt.Printf("\nCacheResourceBundleEntries:107  --  Calling core: %s\n", coreURL)
	// ///fmt.Printf("Using Token: %s\n", token)
	// _, err = client.Do(req)
	return 0
}
