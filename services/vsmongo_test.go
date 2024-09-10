package services

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"os"

	. "github.com/smartystreets/goconvey/convey"

	//log "github.com/sirupsen/logrus"
	"testing"
	//"github.com/joho/godotenv"
)

func TestOpenMongoDB(t *testing.T) {
	//t.Parallel()
	//InitTest()

	fmt.Printf("\n\nTestOpenDB\n")
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey")
		os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1/")
		os.Setenv("SERVICE_NAME", "uc_fhir4")
		os.Setenv("SERVICE_VERSION", "local_test")
		os.Setenv("COMPANY", "test")
		// InitTest()
		// conf := GetConfig()
		config, err := GetServiceConfig("uc_fhir4", "local_test", "test")
		So(err, ShouldBeNil)
		So(config, ShouldNotBeNil)
		//So(conf, ShouldNotBeNil)
		mongo, err := OpenMongoDB()
		So(err, ShouldBeNil)
		So(mongo, ShouldNotBeNil)
		c, err := GetCollection("bundleCache")
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
	})
}

func TestConnectToMongoDB(t *testing.T) {
	//t.Parallel()
	//InitTest()
	//godotenv.Load(".env.core")
	fmt.Printf("\n\nTestConnectToDMongoB \n")
	Convey("Subject: Open the mongo DB", t, func() {
		//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
		fmt.Printf("\n\n--- Convey TestConnectToDB\n")
		os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:20100/api/rest/v1/")
		os.Setenv("SERVICE_NAME", "uc_fhir4")
		os.Setenv("SERVICE_VERSION", "local_test")
		os.Setenv("COMPANY", "test")
		// InitTest()
		// conf := GetConfig()
		config, err := GetServiceConfig("uc_fhir4", "local_test", "test")
		So(err, ShouldBeNil)
		So(config, ShouldNotBeNil)

		So(Company, ShouldEqual, "test")
		mongo, err := ConnectToMongoDB()
		So(mongo, ShouldNotBeNil)
		So(err, ShouldBeNil)
		c, err := GetCollection("bundleCache")
		So(err, ShouldBeNil)
		So(c, ShouldNotBeNil)
		// c, err := GetCollection("configs")
		// So(err, ShouldBeNil)
		// So(c, ShouldNotBeNil)
	})
}
