package services

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	//core "github.com/dhf0820/core_service/connect_core"
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	//t.Parallel()
	fmt.Printf("---TestInitConnector\n")
	Convey("Initializes the fhir4 service getting configuration from Core", t, func() {
		//os.Setenv("ENV_DELIVERY_TEST", "/Users/dhf/work/roi/services/delivery_service/config/config.json")
		os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1")
		conf, err := Initialize()
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		fmt.Printf("cnf: %s\n", spew.Sdump(conf))
	})
}

// func TestGetServiceConfig(t *testing.T) {
// 	//t.Parallel()
// 	fmt.Printf("---TestGetServiceConfig\n")
// 	Convey("Retrieves the Delivery configuration from core", t, func() {
// 		//os.Setenv("ENV_DELIVERY_TEST", "/Users/dhf/work/roi/services/delivery_service/config/config.json")
// 		os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1")
// 		conf, err := GetServiceConfig("uc_fhir4", "local_test", "test") //GetConfig("delivery", "test")

// 		//conf, err := Initialize()
// 		So(err, ShouldBeNil)
// 		So(conf, ShouldNotBeNil)
// 		//So(conf.Name, ShouldEqual, "delivery")
// 		//fmt.Printf("\ncfg: %s\n\n", spew.Sdump(conf))
// 		//for i, me := range conf.MyEndpoints {
// 		//	fmt.Printf("MyConnection: %d - %s\n", i, spew.Sdump(me))
// 		//}
// 	})
// }
