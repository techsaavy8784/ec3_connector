package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"io/ioutil"
	"strings"

	common "github.com/dhf0820/uc_common"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/dhf0820/ids_model/common"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()
var (
	// DateMods      map[string]string
	// SqlMods       map[string]string
	// StringMods    map[string]string
	// NumericMods   map[string]string
	Conf   *common.ServiceConfig //*config
	CoreEp *common.EndPoint
	CaEp   *common.EndPoint
	//DbConnector 	*common.DataConnector
	serverAddress string
	baseAddress   string
	port          string
	tlsMode       string
	deployMode    string
	lis           net.Listener
	PatientData   string
	//*messaging.MessagingClient
	//GWConfig *m.DeliveryConfig
)

type ConfigResp struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Config  common.ServiceConfig `json:"config"`
}

func Start() {
	//var err error
	//var opts = []grpc.ServerOption{}
	//run_env := os.Getenv("SERVICE_VERSION")
	//fmt.Printf("Start Called: [%s]\n", run_env)

	//cfg := GetConfig()

	//TODO: Get Config from Core and DO NOT use mongo dirctly
	fmt.Printf("\n\n\nStart:63  --  cfg: %s\n\n\n\n", spew.Sdump(Conf))
	fmt.Printf("Start:64  ca_3 EndPoint:\n %s", spew.Sdump(Conf.MyEndPoints))

	fmt.Printf("Start:67  --  calling GetMyEndpoint with 'uc_ca3'\n")
	ep := common.GetMyEndpoint(Conf.MyEndPoints, "uc_ca3")
	//OpenCaDB(Conf.DataConnector)
	if ep == nil {
		fmt.Printf("Start:67   EndPoint 'uc_ca3' was not found/n %s", spew.Sdump(Conf.MyEndPoints))
	}
	fmt.Printf("ConnectorCerner EndPoint: %s\n", spew.Sdump(ep))

	baseAddress = ep.Address
	fmt.Printf("CONNECTOR_ADDRESS - %s\n", baseAddress)
	port := ep.Port

	fmt.Printf("CONNECTOR_PORT - %s\n", port)
	deployMode := ep.DeployMode
	//deployMode := os.Getenv("CORE_DEPLOYMODE")
	//deployMode := strings.ToUpper(ep.DeployMode)
	//fmt.Printf("\nStarting [%s] connection\n\n", deployMode)

	if deployMode == "REST" {
		fmt.Printf("Start restful handler\n")

		//fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(ep))
		//restAddress := restEp.Address

		restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
		fmt.Printf("restAddress: %s\n", restAddress)
		router := NewRouter()

		fmt.Printf("\n\n%s listening for restful requests at %s\n\n", ep.Name, restAddress)
		err := http.ListenAndServe(restAddress, router)
		fmt.Printf("This should not happen err = %s\n", err.Error())
		//Processing = false

		//TODO: Start a worker for each facility that has their own EMR  e.g. Mercy
		//Worker("demo")
		//fmt.Printf("Worker returned\n")
		// mainErr := http.ListenAndServe(restAddress, router)
		// if mainErr != nil {
		// 	logrus.Errorf("Rest Startup error: %v", mainErr)
		// }
	} else {
		fmt.Printf("\nOnly restful is supported\n\n")
	}
	//switch deployMode {
	//case "DOCKER":
	//	serverAddress = fmt.Sprintf("0.0.0.0:%s", port)
	//case "LOCAL":
	//	serverAddress = fmt.Sprintf("localhost:%s", port)
	//	logrus.Infof("Using LOCAL mode to address [%s]\n", serverAddress)
	//case "K8S": // This may change to just the port.
	//	serverAddress = fmt.Sprintf(":%s", port)
	//}
	//
	//fmt.Printf("Deploying to %s server  - %s\n", deployMode, serverAddress)
	//lis, err = net.Listen("tcp", serverAddress)
	//if err != nil {
	//	log.Fatalf("Listen Failed: %s\n", err)
	//}
	//s := grpc.NewServer(opts...)
	//
	//fmt.Printf("Using CoreServiceServer [%s]\n", serverAddress)
	//connectorServiceServer := NewConnectorServiceServer()
	//corePB.RegisterCoreServiceServer(s, coreServiceServer) //&releaseServiceServer{})
	//reflection.Register(s)
	////fmt.Printf("Starting Server port: %s\n", serverAddress)
	//go s.Serve(lis)  // Run as goroutine so main can start the http handler
	////restEp := GetMyEndpoint("core")
	////fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(restEp))
	//////restAddress := restEp.Address
	////
	////restAddress := fmt.Sprintf("%s:%s", restEp.Address, restEp.Port)
	////router := h.NewRouter()
	////log.Infof("listening for restful requests at %s", restAddress)
	////mainErr := http.ListenAndServe(restAddress, router)
	////if mainErr != nil {
	////	log.Errorf("Rest Startup error: %v", mainErr)
	////}
}

func Initialize() (*common.ServiceConfig, error) {
	var err error
	fmt.Printf("Initiallizing cerner\n")
	if os.Getenv("SERVICE_NAME") == "" {
		os.Setenv("SERVICE_NAME", "uc_cerner")
	}
	if os.Getenv("SERVICE_VERSION") == "" {
		os.Setenv("SERVICE_VERSION", "local_test")
	}

	if os.Getenv("SERVICE_COMPANY") == "" {
		os.Setenv("SERVICE_COMPANY", "test")
	}

	fmt.Printf("SERVICE_COMPANY:139 - [%s]\n\n", os.Getenv("SERVICE_COMPANY"))
	// PatientData := strings.ToLower(os.Getenv("PATIENT_DATA"))
	// if strings.Trim(PatientData, " ") == "" {
	// 	PatientData = "postgres"
	// }
	Conf, err = GetServiceConfig(strings.ToLower(os.Getenv("SERVICE_NAME")), strings.ToLower(os.Getenv("SERVICE_VERSION")),
		strings.ToLower(os.Getenv("SERVICE_COMPANY")))
	if err != nil {
		return nil, fmt.Errorf("could not retieve Configuration : %v", err)
	}
	//DbConnector, err = common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	//fmt.Printf("\n----config: %s]\n", spew.Sdump(Conf))

	// OpenCaDB(Conf.DataConnectors)

	setEndPoints()
	fmt.Printf("Initilized ca_3Service\n\n")
	return Conf, err
}

func GetServiceConfig(name, version, company string) (*common.ServiceConfig, error) {
	fmt.Printf("GetServiceConfig:171   Starting\n")
	//_, _ = GetCollection("sys_config")
	var cfg ConfigResp
	var err error
	//var unmarshalErr *json.UnmarshalTypeError
	var bdy []byte
	cfg = ConfigResp{}
	fmt.Printf("GetServiceConfig:187  --  name: %s, version: %s, company: %s\n", name, version, company)
	//coreName := strings.ReplaceAll(os.Getenv("CORE_NAME_PORT"), " ","")
	api := os.Getenv("API")
	//Log.Infof("API: [%s]\n", api)
	configAddr := os.Getenv("CONFIG_ADDRESS")
	// configAddr = "http://docker1.ihids.com:19101/api/rest/v1"
	fmt.Printf("GetServiceConfig:193  --  ConfigAddress: %s\n", configAddr)
	Log.Printf("GetServiceConfig:194  --  %s/config?name=%s&version=%s&company=%s\n", configAddr, name, version, company)
	if api == "" || api == "RESTFUL" {
		fmt.Printf("GetServiceConfig:196  --  api is blank\n")
		//

		url := fmt.Sprintf("%s/config?name=%s&version=%s&company=%s", configAddr, name, version, company)
		//url := fmt.Sprintf("http://localhost:19100/api/v1/config?name=%s&version=%s&company=%s", name, version, company)
		fmt.Printf("GetServiceConfig:201  --  core url: %s\n", url)
		startTime := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			Log.Errorf("Get config returned error: %s", err.Error())
			return nil, err
		}
		fmt.Printf("  GetServiceConfig: 208 - Elapsed Time: %s\n", time.Since(startTime))
		fmt.Printf("status : %d\n", resp.StatusCode)
		defer resp.Body.Close()
		//cfg = mod.ServiceConfig{}
		bdy, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("raw string: %s\n", string(bdy))
		err = json.Unmarshal(bdy, &cfg)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Config JSON: %s\n", spew.Sdump(cfg))
		fmt.Printf("  GetServiceConfig: 222 - Elapsed Time: %s\n", time.Since(startTime))
		fmt.Printf("GetServiceConfig: 223 - status : %d\n", resp.StatusCode)
		//_, err = GetCollection("sys_config")
	} else {
		return nil, errors.New("only rest is supported")
	}
	Conf = &cfg.Config
	//DatabaseConnector, err := common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	//fmt.Printf("CFG: = %s\n", spew.Sdump(cfg.Config))
	//fmt.Printf("\n\n###Conf: = %s\n", spew.Sdump(Conf))
	DbConnector, err = common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	fmt.Println("dbConnector:233  -  Opening MongoDb")
	OpenMongoDB()
	return &cfg.Config, err
}

//func GetServiceEndpoint(value string) *mod.EndPoint {
//	fmt.Printf("GetServiceEndpoint\n")
//	endPoints := GetConfig().ServiceEndpoints
//	for _, ep := range endPoints {
//		fmt.Printf("Looking at %s for %s\n", ep.Name, value)
//		if ep.Name == value {
//			fmt.Printf("Found Endpoint: %s\n", ep.Name)
//			return ep
//		}
//	}
//	fmt.Printf("--Endpoint not found\n")
//	return nil
//}

func GetConfig() *common.ServiceConfig {

	return Conf
}

func setEndPoints() {
	//fmt.Printf("\n--In setEndPoints\n")
	// CoreEp = GetServiceEndpoint(Conf.ServiceEndPoints, "core")
	// if CoreEp == nil {
	// 	Log.Errorf("---Core EndPoint was not found in configuration: %s\n", spew.Sdump(Conf))
	// } else {
	// 	//fmt.Printf("---CoreEP: %s\n", spew.Sdump(CoreEp))
	// }

	// CaEp = GetServiceEndpoint(Conf.ServiceEndPoints, "ca_api")
	// if CaEp == nil {
	// 	Log.Errorf("---ca_api EndPoint was not found in configuration: %s\n", spew.Sdump(Conf))
	// } else {
	// 	//fmt.Printf("---CaEP: %s\n", spew.Sdump(CaEp))
	// }
}

func GetConfigDataElement(name string) string {
	data := Conf.Data
	for _, elem := range data {
		if elem.Name == name {
			return elem.Value
		}
	}
	return ""

}

func GetFhirSystem(id string) (*common.FhirSystem, error) {
	//fmt.Printf("\n\n\nGetFhirSystem:277 - for %s\n\n", id)
	//fmt.Printf("GetFhirSystem:278 - GetCollection\n")
	startTime := time.Now()
	collection, err := GetCollection("fhirSystem")
	if err != nil {
		return nil, err
	}
	fmt.Printf("GetFhirSystem:285 - Elapsed time: %s\n", time.Since(startTime))
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("GetFhirSystem:287  -  invalid FhirId: [%s]", id)
	}
	query := bson.D{{"_id", oid}}
	fmt.Printf("Query: %v\n", query)
	//filter := bson.D{{"name", "demo.Cerner"}}
	filter := bson.D{{"_id", oid}}
	//filterM := bson.M{"_id": oid}
	fhirSystem := &common.FhirSystem{}
	startTime = time.Now()
	fmt.Printf("\n\n\nGetFhirSystem:296 -  FindOne fhirConnector: bson.D %v\n", filter)
	err = collection.FindOne(context.Background(), filter).Decode(fhirSystem)
	//fmt.Printf("GetFhirSystem:290 - Elapsed Time: %s\n", time.Since(startTime))
	// if err != nil {
	// 	fmt.Printf("   Now Calling GetFhirConnector FindOne SvcConfig: bson.M %v\n", filterM)
	// 	err = collection.FindOne(context.Background(), filterM).Decode(fhirConfig)
	// }
	if err != nil {
		Log.Errorf("GetFhirSystem:304 FindOne %v NotFound\n", filter)
		return nil, fmt.Errorf("GetFhirSystem:304  FindOne %v NotFound\n", filter)
	}
	return fhirSystem, nil
}
