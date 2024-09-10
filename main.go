package main

import (
	"fmt"

	//"github.com/davecgh/go-spew/spew"
	//common "github.com/dhf0820/uc_common"
	"github.com/joho/godotenv"

	//"net/http"
	"os"

	//	replace github.com/dhf0820/uc_core => ../uc_core
	//h "github.com/dhf0820/Fhir4Service/handlers"
	service "github.com/dhf0820/ec3_connector/services"
	//log "github.com/sirupsen/logrus"
	// mod "github.com/dhf0820/uc_common"
	// "strings"
	//"google.golang.org/grpc/credentials"
)

func main() {
	version := "221222_1"
	os.Setenv("CodeVersion", version)
	fmt.Printf("\n\n----Starting CernerConnector %s\n\n", version)
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Printf("No environment set up using: .env.uc_cerner\n ")
		err := godotenv.Load(".env.uc_cerner")
		if err != nil {
			fmt.Printf("Main: Error getting environment: %v\n", err)
		}
	} else {
		//os.Setenv("ENV_CORE", "/Users/dhf/work/roi/services/core_service/config/config.json")
		// service.InitCore("test")
		// service.OpenDB()
		fmt.Printf("main:37  --  Calling service.Initialize\n")
		service.Initialize()
		fmt.Printf("\n\n\nmain:39  --  Calling service Start\n")
		service.Start() //Should Not Return
		//Start()

		//fmt.Printf("Start returned\n")
		// fmt.Printf("\n\n---Start restful handler\n")
		// eps := service.Conf.MyEndPoints
		// restEp := common.GetMyEndpoint(eps, "fhir4")
		// //restEp := eps[0] //service.GetMyEndPoint("restful_core")
		// fmt.Printf("Restful EndPoint:39 - %s\n", spew.Sdump(restEp))
		// //restAddress := restEp.Address

		// restAddress := fmt.Sprintf("%s:%s", restEp.Address, restEp.Port)
		// fmt.Printf("listening on:43 -  %s\n", restAddress)
		// router := h.NewRouter()
		// fmt.Printf("----listening for restful requests at %s", restAddress)
		// fmt.Printf("Calling ListenAndServe\n")
		// mainErr := http.ListenAndServe(restAddress, router)
		// if mainErr != nil {
		// 	log.Errorf("Rest Startup error: %v", mainErr)
		// }
		fmt.Printf("Main should never stop\n")
	}
}

// func Start() {
// 	//ar opts = []grpc.ServerOption{}
// 	// service.RunEnv = os.Getenv("CONFIG_VERSION")
// 	// service.Company = os.Getenv("COMPANY")
// 	//fmt.Printf("---Start Called: run_env: [%s]  company: [%s]\n", service.RunEnv, service.Company)

// 	//fmt.Printf("32<<20 = %d   32<<24 = %d\n", 32<<20, 32<<22)
// 	configAddress := os.Getenv("CONFIG_AADDRESS")
// 	if configAddress == "" {
// 		os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1")
// 	}

// 	conf, err := service.Initialize() //TODO: get the env value from flag
// 	if err != nil {
// 		fmt.Printf("Main InitCore failed: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Start InitCore: %s\n", spew.Sdump(conf))
// 	//service.OpenDB()
// 	cfg := service.GetConfig()
// 	fmt.Printf("\n---cfg: %s\n", spew.Sdump(cfg))
// 	//ep := GetMyEndpoint("core")
// 	eps := common.GetMyEndpoints(cfg)
// 	//fmt.Printf("GRPC EndPoint: %s\n", spew.Sdump(ep))

// 	for _, ep := range eps {

// 		if ep.Protocol == "http" {
// 			fmt.Printf("\n--Staring restful service\n")
// 			fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(ep))
// 			//restAddress := restEp.Address
// 			//ep.Port = "29900"
// 			restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
// 			router := h.NewRouter()

// 			fmt.Printf("\n\n$$$ Core is listening for restful requests at %s:%s\n\n", "0.0.0.0", ep.Port)
// 			mainErr := http.ListenAndServe(restAddress, router)
// 			if mainErr != nil {
// 				log.Errorf("Rest Startup error: %v", mainErr)
// 			}
// 		} else {
// 			log.Errorf("start:94 - Invalid Protocol: %s\n", ep.Protocol)
// 		}
// 		fmt.Printf("--Start restful service return and should not have")

// 	}
// }
