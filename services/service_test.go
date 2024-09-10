package services

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	//core "github.com/dhf0820/core_service/connect_core"
	"os"
	"testing"
	"time"
)

// func TestInitialize(t *testing.T) {
// 	//t.Parallel()
// 	fmt.Printf("---TestInitConnector\n")
// 	Convey("Initializes the CA Connector system getting configuration from Core", t, func() {
// 		//os.Setenv("ENV_DELIVERY_TEST", "/Users/dhf/work/roi/services/delivery_service/config/config.json")
// 		os.Setenv("CONFIG_ADDRESS", "localhost:9200")
// 		conf, err := Initialize()
// 		So(err, ShouldBeNil)
// 		So(conf, ShouldNotBeNil)
// 	})
// }

func TestGetServiceConfig(t *testing.T) {
	//t.Parallel()
	fmt.Printf("---TestGetServiceConfig\n")
	Convey("Retrieves the Delivery configuration from core", t, func() {
		//os.Setenv("ENV_DELIVERY_TEST", "/Users/dhf/work/roi/services/delivery_service/config/config.json")
		//os.Setenv("CONFIG_ADDRESS", "localhost:9200")
		os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1")
		conf, err := GetServiceConfig("uc_fhir4", "local_test", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		// conf, err = Initialize()
		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)
		//So(conf.Name, ShouldEqual, "delivery")
		//fmt.Printf("\ncfg: %s\n\n", spew.Sdump(conf))
		//for i, me := range conf.MyEndpoints {
		//	fmt.Printf("MyConnection: %d - %s\n", i, spew.Sdump(me))
		//}
		fmt.Printf("System config")
	})
}

func TestGetFhirConnector(t *testing.T) {
	//t.Parallel()
	fmt.Printf("---TestGetFhirConnector\n")
	os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1")
	startTime := time.Now()
	_, err := GetServiceConfig("uc_fhir4", "local_test", "demo") //GetConfig("delivery", "test")
	if err != nil {
		t.Errorf("GetservicConfig:52 - err: %s", err.Error())
		t.Fail()
	}
	fmt.Printf("\nService config done in %s/n", time.Since(startTime))
	//fmt.Printf("Config: %s\n", spew.Sdump(conf))
	Convey("Retrieves FhirConnector", t, func() {
		// startTime := time.Now()
		// os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1")
		// conf, err := GetServiceConfig("uc_fhir4", "local_test", "demo") //GetConfig("delivery", "test")
		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)
		// fmt.Printf("Service config done in %s/n", time.Since(startTime))

		Convey("Retrieves FhirConnector", func() {
			fmt.Printf("\n\n\nGet First FHIR Conector\n")
			fhirConnector, err := GetFhirSystem("62d0ad3c9d0119afff9978b3")
			So(err, ShouldBeNil)
			So(fhirConnector, ShouldNotBeNil)
		})
		Convey("Retrieves Second FhirConnector", func() {
			fmt.Printf("\n\n\nGet Second FHIR Conector\n")
			fhirConnector, err := GetFhirSystem("62d0ad3c9d0119afff9978b3")
			So(err, ShouldBeNil)
			So(fhirConnector, ShouldNotBeNil)
		})
		Convey("Retrieves third FhirConnector", func() {
			fmt.Printf("\n\n\nGet Third FHIR Conector\n")
			fhirConnector, err := GetFhirSystem("62d0ad3c9d0119afff9978b3")
			So(err, ShouldBeNil)
			So(fhirConnector, ShouldNotBeNil)
		})
	})
}

// {
//     "_id" : ObjectId("62d0ac2e0ca10fa533966d13"),
//     "name" : "demo.Demo",
//     "display_name" : "Demo",
//     "url" : "docker1.ihids.com:20102/r4/62d0ac2e0ca10fa533966d13/api/rest/v1/",
//     "endpoints" : null,
//     "fhir_info" : null,
//     "enabled" : "true",
//     "fhir_fields" : null,
//     "fields" : null,
//     "data" : null,
//     "created_at" : ISODate("2022-07-14T23:52:14.587+0000"),
//     "updated_at" : ISODate("2022-07-14T23:52:14.587+0000"),
//     "actualURL" : "http://docker1.ihids.com:19400/r4/628d2c8a9f29b7032e9a154e/api/rest/v1/"
// }
//
//func TestConnectToCore(t *testing.T) {
//	fmt.Printf("\n\nTestGetConfig\n")
//	Convey("Connects to core via GRPC", t, func() {
//		os.Setenv("CONFIG_ADDRESS", "localhost:9200")
//		client, err := core.Connect()
//		//client, err := ConnectToCore()
//		So(err, ShouldBeNil)
//		So(client, ShouldNotBeNil)
//		//req := pb.ConfigRequest{
//		//	Name: "delivery",
//		//	Version: "test",
//		//}
//		conf, err := core.GetServiceConfig("delivery", "test")    //GetConfig("delivery", "test")
//		//resp, err := client.GetServiceConfig(context.Background(), &req)
//		So(err, ShouldBeNil)
//		So(conf, ShouldNotBeNil)
//		//So(resp, ShouldNotBeNil)
//		//conf := toServiceConfig(resp.GetServiceConfig())
//		//fmt.Printf("Received: %s\n", spew.Sdump(conf))
//	})
//
//}
