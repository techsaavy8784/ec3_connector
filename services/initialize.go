package services

import (
	//"github.com/dhf0820/fhir4"
	//"encoding/json"
	//"fmt"
	//"github.com/davecgh/go-spew/spew"
	common "github.com/dhf0820/uc_common"
	// log "github.com/sirupsen/logrus"
	// "io/ioutil"
	// "net/http"
	// "os"
	// "strings"
)

var (
	//Conf        *common.ServiceConfig
	Mongo       *MongoDB
	Company     string
	DbConnector *common.DataConnector
)

// type ConfigResp struct {
// 	Status  int                   `json:"status"`
// 	Message string                `json:"message"`
// 	Config  *common.ServiceConfig `json:"config"`
// }

// func Initialize() (*common.ServiceConfig, error) {
// 	var err error
// 	fmt.Printf("Initiallizing FhirCa\n")
// 	if os.Getenv("SERVICE_NAME") == "" {
// 		os.Setenv("SERVICE_NAME", "uc_fhir4")
// 	}
// 	if os.Getenv("SERVICE_VERSION") == "" {
// 		os.Setenv("SERVICE_VERSION", "local_test")
// 	}

// 	if os.Getenv("SERVICE_COMPANY") == "" {
// 		os.Setenv("SERVICE_COMPANY", "test")
// 	}
// 	Company = strings.ToLower(os.Getenv("SERVICE_COMPANY"))
// 	fmt.Printf("SERVICE_COMPANY:41 - [%s]\n\n", Company)
// 	// PatientData := strings.ToLower(os.Getenv("PATIENT_DATA"))
// 	// if strings.Trim(PatientData, " ") == "" {
// 	// 	PatientData = "postgres"
// 	// }
// 	Conf, err = GetServiceConfig(strings.ToLower(os.Getenv("SERVICE_NAME")),
// 		strings.ToLower(os.Getenv("SERVICE_VERSION")), Company)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not retieve Configuration : %v", err)
// 	}

// 	//fmt.Printf("\n----config: %s]\n", spew.Sdump(Conf))
// 	Mongo, err = OpenMongoDB()

// 	// OpenCaDB(Conf.DataConnectors)
// 	// setEndPoints()
// 	// setupQueryMidifiers()
// 	fmt.Printf("Initilized fhir4:latest\n\n")
// 	return Conf, err
// }

// func GetServiceConfig(name, version, company string) (*common.ServiceConfig, error) {
// 	var cfg ConfigResp
// 	var err error
// 	//var unmarshalErr *json.UnmarshalTypeError
// 	var bdy []byte
// 	//cfg = ConfigResp{}
// 	log.Infof("GetServiceConfig-55 name: %s, version: %s, company: %s",
// 		name, version, company)
// 	//coreName := strings.ReplaceAll(os.Getenv("CORE_NAME_PORT"), " ","")
// 	api := os.Getenv("API")
// 	log.Infof("API: [%s]\n", api)
// 	configAddr := os.Getenv("CONFIG_ADDRESS")
// 	fmt.Printf("GetSereviceConfig:74 ConfigAddress: %s\n", configAddr)
// 	log.Infof(fmt.Sprintf("%s/config?name=%s&version=%s&company=%s", configAddr, name, version, company))
// 	if api == "" || api == "RESTFUL" {
// 		log.Infof("api is blank\n")
// 		//

// 		url := fmt.Sprintf("%s/config?name=%s&version=%s&company=%s", configAddr, name, version, company)
// 		log.Infof("GetServiceConfig:81 - Config url: %s", url)
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			log.Errorf("Get config returned error: %v", err)
// 			return nil, err
// 		}
// 		fmt.Printf("status : %d\n", resp.StatusCode)
// 		defer resp.Body.Close()
// 		//cfg = mod.ServiceConfig{}
// 		bdy, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fmt.Printf("raw string: %s\n", string(bdy))
// 		err = json.Unmarshal(bdy, &cfg)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fmt.Printf("Config JSON: %s\n", spew.Sdump(cfg))

// 	} else {
// 		return nil, fmt.Errorf("%s is not supported yet", api)
// 	}
// 	err = nil
// 	Conf = cfg.Config
// 	// fmt.Printf("Number of dataConnectors = %d\n", len(Conf.DataConnectors))
// 	// fmt.Printf("CallingGetDb with %s\n", spew.Sdump(Conf.DataConnectors))
// 	// DbConnector, err := common.GetDatabaseByName(Conf.DataConnectors, "mongo")
// 	// if err != nil {
// 	// 	fmt.Printf("GetDatabase error: %s\n", err.Error())
// 	// }
// 	DbConnector = Conf.DataConnectors[0]
// 	fmt.Printf("DbConnector: %s\n", spew.Sdump(DbConnector))
// 	Company = DbConnector.Database
// 	//fmt.Printf("CFG: = %s\n", spew.Sdump(cfg.Config))
// 	//fmt.Printf("\n\n###Conf: = %s\n", spew.Sdump(Conf))

// 	return cfg.Config, err
// }

// func GetConfig() *common.ServiceConfig {
// 	return Conf
// }
