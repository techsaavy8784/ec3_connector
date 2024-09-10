package services

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	//"strings"
	//"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	timeout = 15
)

var (
	//err error
	body []byte
)

// RetData is the mapped json of the request
type RetData map[string]interface{}

// Connection is a FHIR connection
type Connection struct {
	BaseURL string
	client  *http.Client
}

//TODO: Consider including the FhirSystem in the Connection
// New creates a new connection
func New(baseurl string) *Connection {
	return &Connection{
		BaseURL: baseurl,
		client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   time.Duration(timeout*120) * time.Second,
					KeepAlive: time.Duration(timeout*120) * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   time.Duration(timeout) * time.Second,
				ResponseHeaderTimeout: time.Duration(timeout) * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
}

// Query sends a query to the base url
func (c *Connection) Query(q string) ([]byte, error) {
	fmt.Printf("\n\n\n\nQuery:55  --  BaseUrl: %s  -  Query param: %s\n\n\n\n", c.BaseURL, q)
	if q == "" {
		return nil, fmt.Errorf("Query: query parameter missing")
	}
	url := fmt.Sprintf("%s/%s", c.BaseURL, q)
	//fmt.Printf("fhir4_query:60  --  c.BaseUrl = %s\n", c.BaseURL)

	return c.GetFhir(url)
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// req.Header.Add("Accept", "application/json+fhir")
	// //fmt.Println("Calling the GET request")
	// resp, err := c.client.Do(req)
	// if err != nil {
	// 	log.Errorf(" !!!fhir query returned err: %s\n", err)
	// 	return nil, err
	// }
	// //fmt.Printf("resp: %s\n", spew.Sdump(resp))
	// //defer resp.Body.Close()
	// if resp.StatusCode < 200 || resp.StatusCode > 299 {
	// 	err = fmt.Errorf("%d|%s", resp.StatusCode, string(body))
	// 	return nil, err
	// }
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Query Error: %v\n", err)
	// 	return nil, err
	// }

	// //fmt.Printf("fhir length of Body: %d\n", len(body))
	// return body, nil
}

// func (c *Connection) GetById(id string)([]byte, error ) {

// }

func (c *Connection) GetFhir(url string) ([]byte, error) {
	log.Infof("GetFhir:95 URL Requested: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("GetFhir:98  --  !!!NewRequest failed: %s\n", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json+fhir")
	//fmt.Printf("getFhir:102  --  req: %s\n", spew.Sdump(req))
	resp, err := c.client.Do(req)
	if err != nil {
		log.Errorf("GetFhir:105  --  !!!fhir query returned err: %s\n", err)
		return nil, err
	}
	//fmt.Printf("GetFhir:108  --  resp = %s\n", spew.Sdump(resp))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Errorf("GetFhir:110  --  returned error of %d - %s\n", resp.StatusCode, resp.Status)
		err = fmt.Errorf("%d|fhir:105 %s", resp.StatusCode, resp.Status)
		//log.Errorf("%s", err.Error())
		return nil, err
	}
	// data := DocumentReference{}
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	fmt.Printf("NewDecoder error: %s\n", err.Error())
	// }
	// fmt.Printf("NewDecoder: %s\n\n", spew.Sdump(data))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadBody Error:119 %s\n", err.Error())
		return nil, err
	}
	fmt.Printf("GetFhir:122 returning no error and length of data: %d\n", len(body))
	return body, nil
}

// func (c *Connection)PatientNextPage(url string) {
// 	bytes, err := c.GetFhir(url)
// }
