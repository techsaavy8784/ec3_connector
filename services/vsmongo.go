package services

import (
	"context"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	common "github.com/dhf0820/uc_common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"time"
)

type MongoDB struct {
	Client       *mongo.Client
	DatabaseName string
	URL          string
	Database     *mongo.Database
	Session      mongo.Session
	Collection   *mongo.Collection
}

var DB MongoDB
var mongoClient *mongo.Client
var dbConnector *common.DataConnector

//var DbConnector *DataConnector
var insertResult *mongo.InsertOneResult

func OpenDBUrl(dbURL string) *MongoDB {
	var err error
	//svcConfig := GetConfig()
	//if svcConfig == nil {
	//	fmt.Printf("\n---$$$Config is not initialized\n\n")
	//}
	startTime := time.Now()
	uri := dbURL
	fmt.Printf("Opening database:39  %s\n", uri)
	//uri := dbURL + databaseName
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	// client, error := vsmongo.NewClient(options.Client().ApplyURI("ur_Database_uri"))
	// error = client.Connect(ctx)

	// //Checking the connection
	// error = client.Ping(context.TODO(), nil)
	// fmt.Println("Database connected")

	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	fmt.Printf("vsmongo:56  - Connecting to mongo\n")
	if mongoClient, err = mongo.Connect(ctx, opts); err != nil {
		msg := fmt.Sprintf("vsmmongo:57  mongo.Connect error: %s\n", err.Error())
		log.Error(msg)
		return nil
	}
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		//fmt.Printf("Database did not connect:62 %v\n", err)
		log.Errorf("Database did not connect:63 %s", err.Error())
		return nil
	}
	fmt.Printf("Company(DatabaseName: %s\n", os.Getenv("COMPANY"))

	DB.Client = mongoClient
	DB.DatabaseName = Company //.Getenv("COMPANY") //DbConnector.Database  //databaseName
	DB.Database = mongoClient.Database(DB.DatabaseName)
	DB.URL = dbURL
	fmt.Printf("Database: %s\n", DB.DatabaseName)
	fmt.Println("Database connected")
	//fmt.Printf("Client: %s\n", spew.Sdump(client))

	DB.Collection = DB.Client.Database(DB.DatabaseName).Collection(GetDbField("collection"))
	fmt.Printf("DBOpen-77 took %d ms\n", time.Since(startTime).Milliseconds())
	return &DB
}

func OpenMongoDB() (*MongoDB, error) {
	var err error
	//svcConfig := GetConfig()
	//if svcConfig == nil {
	//	fmt.Printf("\n---$$$Config is not initialized\n\n")
	//}
	startTime := time.Now()
	//DbConnector = svcConfig.DataConnector
	//dbURL := DbConnector.Server
	// DbConnector, err := common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	// if err != nil {
	// 	return nil, err
	// }
	dbURL := DbConnector.Server
	//dbURL := DBUrl() //os.Getenv("CORE_DB")
	uri := dbURL
	fmt.Printf("Opening database:97 uri:[%s]\n", uri)
	//uri := dbURL + databaseName
	//ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	//defer cancel()

	// client, error := vsmongo.NewClient(options.Client().ApplyURI("ur_Database_uri"))
	// error = client.Connect(ctx)

	// //Checking the connection
	// error = client.Ping(context.TODO(), nil)
	// fmt.Println("Database connected")

	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
	clientOptions.SetMaxPoolSize(5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Printf("vsmongo:114 -- Using new connect routine from atlas\n")
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		msg := fmt.Sprintf("vsMongo:121 Mongo.Connect error: %s\n", err.Error())
		log.Fatal(msg)
	}

	// if mongoClient, err = mongo.Connect(ctx, clientOptions); err != nil {
	// 	msg := fmt.Sprintf("Mongo.Connect error: %s\n", err.Error())
	// 	log.Error(msg)
	// 	return nil
	// }
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		//fmt.Printf("Database did not connect: %v\n", err)
		log.Errorf("Database:131 did not connect by ping:129 %v", err)
		return nil, err
	}

	DB.Client = mongoClient
	DB.DatabaseName = DbConnector.Database
	//DB.DatabaseName = os.Getenv("COMPANY") //DbConnector.Database  //databaseName
	//DB.DatabaseName = "test"
	fmt.Printf("DatabaseName:140 -  %s\n", DB.DatabaseName)
	DB.Database = mongoClient.Database(DB.DatabaseName)
	DB.URL = dbURL
	fmt.Printf("vsmongo:142 -- Database: %s\n", DB.DatabaseName)
	fmt.Println("Database connected")
	//fmt.Printf("Client: %s\n", spew.Sdump(client))

	DB.Collection = DB.Client.Database(DB.DatabaseName).Collection(GetDbField("collection"))
	fmt.Printf("vsmongo:147 - DBOpen -- took %d ms\n", time.Since(startTime).Milliseconds())
	return &DB, err
}

func DBUrl() string {

	//cacheDB := os.Getenv("CACHE_DB")
	var err error
	fmt.Printf("vsmongo.DBUrl:157 - GetDatabaseByName\n")
	dbConnector, err = common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	if err != nil {
		return ""
	}
	//fmt.Printf("vsmongo.DBUrl:162 - GotDataConnector: %s\n", spew.Sdump(dbConnector))
	return dbConnector.Server
}

// ConnectToDB starts a new database connection and returns a reference to it
func ConnectToMongoDB() (*MongoDB, error) {
	// DbConnector = GetConfig().DataConnector
	// databaseName := DbConnector.Database
	// url := DbConnector.Server
	// coreDB := os.Getenv("CORE_DB")
	// if coreDB == "" {
	// 	coreDB = "LOCAL_DB"
	// }
	// fmt.Printf("Using database: %s\n", coreDB)
	url := DBUrl()
	if url == "" {
		log.Panic("coreDB is not defined. Should contain the name of the actual Database to use\n")
	}

	fmt.Printf("Use DB URL: %s\n", url)
	//databaseName := os.Getenv("COMPANY")
	//databaseName := Company
	databaseName := dbConnector.Database
	fmt.Printf("Using DB: [%s]\n", databaseName)
	//if url == "" {
	//	url = settings.DbURL()
	//}
	//fmt.Printf("Mongo URL: %s\n", url)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(url)

	options.SetMaxPoolSize(DbPoolSize())
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	DB.Client = client
	DB.DatabaseName = databaseName
	DB.Database = client.Database(DB.DatabaseName)
	DB.URL = url
	return &DB, nil
}

func DbPoolSize() uint64 {
	var poolSize uint64
	poolSizeString := GetDbField("poolsize")
	if poolSizeString == "" {
		poolSizeString = "100"
	}
	poolSizeInt64, err := strconv.ParseInt(poolSizeString, 10, 64)
	if err == nil {
		poolSize = uint64(poolSizeInt64)
	} else {
		poolSize = 100
	}
	return poolSize
}
func Current() (*MongoDB, error) {
	if DB.Client != nil {
		return &DB, nil
	}
	_, err := ConnectToMongoDB()
	//client, err := Open("")
	return &DB, err
}

func (db *MongoDB) Close() error {
	err := db.Client.Disconnect(context.TODO())
	return err
}

func GetCollection(collection string) (*mongo.Collection, error) {
	if collection == "" {
		collection = CollectionName()
		fmt.Printf("Using default Collection: %s\n", collection)
	}
	db, err := Current() //"mongodb://admin:Sacj0nhat1@cat.vertisoft.com:27017")
	if err != nil {
		fmt.Printf("Current DB returned error: %s\n", err)
		log.Fatal(err)
		//return nil, err
	}
	client := db.Client
	coll := client.Database(DB.DatabaseName).Collection(collection)
	fmt.Printf("Changed to Collection: %s\n", collection)
	return coll, nil
}

func CollectionName() string {
	return "srv_config"
}

func GetDbField(key string) string {
	return ""
	// //LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
	// flds := mod.DataConnector.Fields
	// for _, fld := range flds {
	// 	switch {
	// 	case fld.Name == key:
	// 		return fld.Value
	// 	}
	// }
	// return ""

}

// IsDup returns whether err informs of a duplicate key error because
// a primary key index or a secondary unique index already has an entry
// with the given value.
func IsDup(err error) bool {
	if wes, ok := err.(mongo.WriteException); ok {
		for i := range wes.WriteErrors {
			if wes.WriteErrors[i].Code == 11000 || wes.WriteErrors[i].Code == 11001 || wes.WriteErrors[i].Code == 12582 || wes.WriteErrors[i].Code == 16460 {
				return true
			}
		}
	}
	return false
}
