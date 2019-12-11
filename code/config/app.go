package config

var MongoStoresCollection = "stores"
var ElasticStoresIndex = "stores"

type ElasticConfiguration struct {
	DbUrl    string
	UserName string
	Password string
}

type MongoConfiguration struct {
	DbUrl      string
	Port       string
	DbName     string
	UserName   string
	Password   string
	ConnString string
	MinPool    uint64
	MaxPool    uint64
}

func GetMongoConfig() *MongoConfiguration {
	return &MongoConfiguration{
		DbUrl:      "",
		Port:       "27017",
		DbName:     "",
		UserName:   "arep",
		Password:   "****",
		ConnString: "mongodb://arep:***@docdb-2019-*****.cluster-ctoeottntm9y.us-east-1.docdb.amazonaws.com:27017",
	}
}

func GetElasticConfig() *ElasticConfiguration {
	return &ElasticConfiguration{
		DbUrl:    "***",
		UserName: "elastic",
		Password: "****",
	}
}
