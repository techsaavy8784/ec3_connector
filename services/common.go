package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	//common "github.com/dhf0820/uc_common"
	// db "github.com/dhf0820/cadatabase"
	// fhir "github.com/dhf0820/fhirR2go"
	"github.com/gorilla/mux"
	//"github.com/gorilla/schema"
	//"github.com/dhf0820/ids_model/common"
	// "github.com/oleiade/reflections"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	// "strconv"
	"strings"
	"time"
)

var currentUrl string
var protocol string
var requestURI string
var baseUrl string

func GetDeploymentFacility(r *http.Request) string {
	params := mux.Vars(r)
	fmt.Printf("\n$$$ params: %v\n", params)
	//release := params["release"]
	// cust_key := params["cust_key"]
	// if cust_key == "" {
	// 	log.Error("No Customer specified, default to demo")
	// 	return "demo"
	// }
	facility := GetConfig().Customer.Facility
	if facility == "" {
		facility = "demo"
	}
	fmt.Printf("Using facility %s\n", facility)
	return facility //for now
}

func GetCurrentURL(r *http.Request) string {
	//var config *ServiceConfig
	params := mux.Vars(r)

	//fmt.Printf("Request: %s\n", spew.Sdump(r))
	proto := strings.Split(r.Proto, "/")[0]
	proto = strings.ToLower((proto))
	cust_key := params["cust_key"]
	release := params["release"]

	base := fmt.Sprintf("%s://%s/%s/%s/api/rest/v1/", proto, r.Host, release, cust_key)
	fmt.Printf("CurrentUrl: %s\n", base)

	//config := GetConfig()

	customer := &Customer{}
	fmt.Printf("Customer: %s\n\n", *customer)
	//addr := fmt.Sprintf(r.P)
	return base
}

func GetFHIRVersion(r *http.Request) string {
	//params := mux.Vars(r)
	fhirVersion := r.Header.Get("Fhir-Version")
	// release := params["release"]
	// if release == "" {
	// 	release = "R4"
	// }
	fmt.Printf("GetFHIRVersion:73 - %s\n\n", fhirVersion)
	return fhirVersion
}

func GetFhirId(r *http.Request) string {
	log.Printf("GetFhir:78  --  r = %s\n", spew.Sdump(r))
	params := mux.Vars(r)
	log.Printf("GetFhir:80  --  params = %s\n", params)
	fhirId := params["fhirId"]
	if fhirId == "" {
		fhirId = r.Header.Get("Fhir-System")
	}
	fmt.Printf("GetFhirId:85 - fhirSystemId = %s\n", fhirId)
	return fhirId
	// if params["fhir_id"] != "" {
	// 	return params["fhir_id"]
	// }else {
	// 	return r.Header.Get("fhirSystem")
	// }
}

func GetFHIRResource(r *http.Request) string {
	params := mux.Vars(r)
	resource := params["resource"]
	if resource == "" {
		uriParts := strings.Split(r.RequestURI, "v1/")
		fmt.Printf("\nuriParts: %v\n", uriParts)
		uriParts1 := strings.Split(uriParts[1], "/")
		resource = uriParts1[0]
	}
	return resource
}

// func ParseSearchParamDate(searchParam *SearchParam) (mod string, dt string, err error) {
// 	//var err error
// 	fmt.Printf("\n\n\n\n$$$$\n")
// 	fmt.Printf("ParseSearchParamDate:82 -  %s\n", spew.Sdump(searchParam))
// 	rawDate := strings.Trim(searchParam.Value, " ")
// 	if rawDate == "" {
// 		return "", "", fmt.Errorf("date can not be blank")
// 	}
// 	datePrefix := rawDate[0:1]
// 	fmt.Printf("datePrefix [%s]\n", datePrefix)
// 	if datePrefix == "1" || datePrefix == "2" {
// 		mod = "eq"
// 		dt = rawDate
// 		fmt.Printf("No Prefix in original date: %s\n", dt)
// 	} else {
// 		mod = rawDate[0:2]
// 		mod = strings.ToLower(mod)

// 		fmt.Printf("mod:97 %s\n", mod)
// 		fmt.Printf("SearchDate: %s\n", rawDate[2:])
// 		_, ok := DateMods[mod]
// 		if !ok {
// 			return "", "", fmt.Errorf("invalid date modifier: [%s]", mod)
// 		}
// 		dt = rawDate[2:]
// 		fmt.Printf("SearchParamDate:487 - Validate: %s\n", dt)
// 	}
// 	_, err = ValidateDate(dt)
// 	if err != nil {
// 		mod = ""
// 		dt = ""
// 		err = fmt.Errorf("invalid date: %s ", dt)
// 	}

// 	return
// }

// func ValidateDate(dt string) (time.Time, error) {
// 	layout := "2006-01-02"
// 	fmt.Printf("Validating [%s] is a valid date\n", dt)
// 	tdt, err := time.Parse(layout, dt)
// 	if err != nil {
// 		fmt.Printf("Time parse: %s -  error: %s\n", dt, err)

// 		return tdt, err
// 	}
// 	fmt.Printf("\ntdt : %s\n", tdt)
// 	return tdt, nil
// }

// func LookupDateModifier(mod string) string {
// 	fmt.Printf("Looking up Modifier: [%s]\n", mod)
// 	fmt.Printf("\nmodifiers: %s\n", spew.Sdump(DateMods))
// 	val, ok := DateMods[mod]
// 	if !ok {
// 		fmt.Printf("invalid Modifier: [%s]\n", mod)
// 		return ""
// 	}
// 	return val
// }
// func LookupSqlDateModifier(mod string) string {
// 	fmt.Printf("Looking up SQLModifier: [%s]\n", mod)
// 	fmt.Printf("\nSqlModifiers: %s\n", spew.Sdump(SqlMods))
// 	val, ok := SqlMods[mod]
// 	if !ok {
// 		fmt.Printf("invalid SqlModifier: [%s]\n", mod)
// 		return ""
// 	}
// 	return val
// }

type Client struct {
	Source     string `json:"source"`
	User       string `json:"user"`
	SourceId   string `json:"source_id"`
	ExternalId string `json:"external_id"`
}

type Connector struct {
	Id   string `bson:"id", json:"id"`
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
	Data KVData `bson:"data" json:"data`
}

type ConnectorConfig struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name            string             `json:"name"`
	Version         string             `json:"version"`
	ListenAddress   *ListenAddress     `json:"listen_address,omitempty" bson:"listen_address,omitempty"`
	EndPoints       []*EndPoint        `json:"endpoints,omitempty"` //All System endpoints this connector needs to talk to
	Pulls           string             `json:"pulls,omitempty"`     //Queue, Poll for work, Wait for work
	DestinationInfo []*KVData          `json:"destination_info" bson:"destination_info"`
	Priority        string             `json:"priority,omitempty" bson:"priority,omitempty"`
	Enabled         string             `json:"enabled,omitempty"` //true/false
	Fields          []*Field           `json:"fields,omitempty"`
	URL             string             `json:"url"`
	Data            []*KVData          `json:"data"`
	//MyEndPoints     []*EndPoint        `json:"my_end_points" bson:"my_endpoints"`
}

//
type ConnectorResponse struct {
	Status    int              `json:"status"`
	Message   string           `json:"message"`
	Connector *ConnectorConfig `json:"connector"`
}

/* ConnectAddress defines the external service used by the connector to actually delivery the release.
For example, smtp mail server, configured mail api, configured fax api, EMR connection details
*/
type ConnectAddress struct {
	Name          string    `json:"name"`     // The name of the function do deliver smtp2Go, SendGrid,FxService
	Protocol      string    `json:"protocol"` //GRPC, HTTP, AMQP
	Address       string    `json:"address"`
	Authorization string    `json:"authorization"`
	Data          []*KVData `json:"data,omitempty"` // Any additional the destination may need configured
}

type ConnectInfo struct {
	//ID    string `json:"id" bson:"id,omitempty"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Customer struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Facility string `json:"facility"`
	//UserId 		string 	`json:"user_id"`
}

type DataConnector struct {
	DbName     string    `json:"dbName" bson:"db_name"`
	Server     string    `json:"server"`
	User       string    `json:"user"`
	Password   string    `json:"password"`
	Database   string    `json:"database"`
	Collection string    `json:"collection"`
	Fields     []*KVData `json:"fields"`
}

type EndPoint struct { // Replaces BaseUrl
	Name        string `json:"name"` //internal name
	Label       string `json:"label"`
	Scope       string `json:"scope,omitempty" bson:"scope"`
	Protocol    string `json:"protocol" bson:"protocol"` // grpc or amqp
	Address     string `json:"address" bson:"address"`   //How do I get to this service
	Port        string `json:"port"`
	Credentials string `json:"credentials" bson:"credentials"`
	CertName    string `json:"certname" bson:"cert_name"`
	TLSMode     string `json:"tlsmode" bson:"tls_mode"`
	DeployMode  string `json:"deploymode" bson:"deploy_mode"` // RestService, GRPCService, messageQueue
	QueueName   string `json:"queue_name" bson:"queue_name"`
}

type Field struct {
	Name         string `json:"name"`
	Label        string `json:"label"`
	Default      string `json:"default"`
	Value        string `json:"value"`
	DisplayValue string `json:"display_value" bson:"display_value"`
	Required     string `json:"required omitempty"`
	UserVisible  string `json:"user_visible omitempty" bson:"user_visible"`
	IsNameValue  string `json:"is_name_value" bson:"is_name_value"`
	Sensitive    string `json:"sensitive"`
}

type KVData struct {
	Name  string
	Value string
}

type ListenAddress struct {
	Name          string    `json:"name"`
	Primary       string    `json:"primary"`
	Pulls         string    `json:"pulls,omitempty"`          //true/false string
	BaseURL       string    `json:"base_url" bson:"base_url"` // First part is the protocol http://localhost:port/api/v1
	Authorization string    `json:"authorization"`
	Data          []*KVData `json:"fields"`
}

type Option struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Value    string             `json:"value"`
	Required string             `json:"required"`
	Module   string             `json:"module,omitempty"`
}

type ServiceConfig struct {
	ID               primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Customer         Customer           `json:"customer"`
	Name             string             `json:"name"`
	Version          string             `json:"version"`
	DataConnector    *DataConnector     `json:"data_connector"`
	DataConnectors   []*DataConnector   `json:"data_connectors"`
	Services         []*ServiceScope    `json:"services" bson:"services"` //Services used by this Connector
	MyEndPoints      []*EndPoint        `json:"my_endpoints"`
	ServiceEndPoints []*EndPoint        `json:"service_endpoints" bson:"service_endpoints"`
	ConnectInfo      []*KVData          `json:"connect_info" bson:"connect_info"`
	Data             []*KVData          `json:"data" bson:"data"`
	Connectors       []ConnectorConfig  `json:"connectors" bson:"connectors"`
	CallBacks        []*KVData          `json:"call_backs" bson:"call_backs"`
}

type ServiceScope struct {
	Name  string `json:"name" bson:"name"`
	Scope string `json:"scope" bson:"scope"` // min, norm, max
}

type Status struct {
	State      string    `json:"state" bson:"state"` // "submitted", "pending", "queued", "inprocess", "delivered", "error", "failed"
	StatusTime time.Time `json:"status_time" bson:"status_time"`
	Comment    string    `json:"comment"`
}

type StatusReport struct {
	StatusType string    `json:"status_type" bson:"status_type"` //update, critical
	Status     string    `json:"status" bson:"status"`           // "submitted", "pending", "queued", "inprocess", "delivered", "error", "failed"
	Timestamp  time.Time `json:"timestamp" bson:"timestamp"`
	Nanotime   int64     `json:"nanotime" bson:"nanotime"`
	Comment    string    `json:"comment" bson:"comment"`
}

type StatusResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//#############################  Methods  ##########

func GenerateUUID() string {
	id := uuid.New()
	id.MarshalBinary()
	return id.String()
}

func GetDataByName(data []*KVData, name string) string {
	for _, d := range data {
		if d.Name == name {
			return d.Value
		}
	}
	return ""
}

func GetFieldByName(flds []*Field, name string) (*Field, error) {
	//fmt.Printf("GetFieldByName looking for %s\n", name)
	//fmt.Printf("  in: %s\n", spew.Sdump(flds))
	for _, fld := range flds {
		//fmt.Printf("looking at [%s] for [%s]\n", fld.Name, name)
		if fld.Name == name {
			return fld, nil
		}
	}
	return nil, fmt.Errorf("Field %s was not found", name)
}

func GetDataConnectorByName(dcs []*DataConnector, name string) (*DataConnector, error) {
	for _, dc := range dcs {
		if dc.DbName == name {
			return dc, nil
		}
	}
	return nil, fmt.Errorf("Field %s was not found", name)
}

func GetKVData(hash []*KVData, key string) string {
	for _, hv := range hash {
		if hv.Name == key {
			return hv.Value
		}
	}
	return ""
	//return "", fmt.Errorf("Key %s was Not Found", key)
}

func GetMyEndpoint(endPoints []*EndPoint, value string) *EndPoint {
	//endPoints := GetConfig().MyEndPoints
	fmt.Printf("GetMyEndPoint:387 --  %v", endPoints)
	for _, ep := range endPoints {
		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
		if ep.Name == value {
			fmt.Printf("Found Endpoint: %s\n", spew.Sdump(ep))
			return ep
		}
	}
	return nil
}

func GetMyEndpoints(cfg *ServiceConfig) []*EndPoint {
	endPoints := cfg.MyEndPoints
	fmt.Printf("Core Endpoints: %s", spew.Sdump(endPoints))
	return endPoints
}

func GetServiceEndpoint(endPoints []*EndPoint, value string) *EndPoint {
	for _, ep := range endPoints {
		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
		if ep.Name == value {
			//fmt.Printf("Found Endpoint: %s\n", ep.Name)
			return ep
		}
	}
	return nil
}

func SendStatus(statusType string, status *Status, url string) {
	SendStatusReport(statusType, status.State, status.Comment, url)
}

func SendStatusReport(statusType string, status string, comment string, statusURL string) {
	fmt.Printf("/////////////////// SendStatusReport to %s  //////////////\n", statusURL)
	stat := StatusReport{}
	stat.Status = status
	stat.StatusType = statusType
	stat.Comment = comment
	stat.Timestamp = time.Now().UTC()
	stat.Nanotime = time.Now().UnixNano()
	fmt.Printf("\n#### Send StatusReport: %s\n", spew.Sdump(stat))
	bstr, err := json.Marshal(stat)
	if err != nil {
		log.Println("Error marshaling status into json:", err.Error())
		return
	}
	client := &http.Client{}
	fmt.Printf("Send Status to: [%s]\n", statusURL)
	req, _ := http.NewRequest("POST", statusURL, bytes.NewBuffer(bstr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	//defer resp.Body.Close()
	if err != nil {
		log.Println("Error sending log message:", err.Error())

	}
}

func StrPtr(str string) *string {
	return &str
}
func IntPtr(intVal int) *int {
	return &intVal
}
func Int64Ptr(intVal int64) *int64 {
	return &intVal
}
func Int32Ptr(intVal int32) *int32 {
	return &intVal
}

func BoolPtr(boolVal bool) *bool {
	return &boolVal
}

//func LogMessage(payload *Payload, log_type string, status string, message string) {
//	url := payload.Config.log_url
//	var msg Message
//
//	msg.Delivery_id = payload.Delivery_id
//	msg.Log_type = log_type (basic or detail)
//	msg.Status = status  (error, success, info)
//	msg.Connector = Connector
//	msg.Message = message
//	msg.Timestamp = time.Now().UTC()
//	msg.Nanotime = time.Now().UnixNano()
//
//	bstr, err := json.Marshal(msg)
//	if err != nil {
//		log.Println("Error marshaling log message into json:", err.Error())
//		return
//	}
//	client := &http.Client{}
//	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bstr))
//	req.Header.Set("Accept", "application/json")
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := client.Do(req)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Println("Error sending log message:", err.Error())
//		return
//	}
//	return
//}

//if err := r.ParseForm(); err != nil {
//	log.Println("Error parsing query: ", err.Error())
//	err := WriteGenericResponse(w, 400, "Error parsing query: "+err.Error())
//	if err != nil {
//		log.Println("Error writing response: ", err.Error())
//	}
//	return
//}
//
//filter := new(ServiceFilter)
//var decoder = schema.NewDecoder()
//decoder.IgnoreUnknownKeys(true)
//
//if err := decoder.Decode(filter, r.Form); err != nil {
//	log.Println("Error decoding query: ", err.Error())
//	err := WriteGenericResponse(w, 400, "Error decoding query: "+err.Error())
//	if err != nil {
//		log.Println("Error writing response: ", err.Error())
//		return
//	}
//	return
//}
