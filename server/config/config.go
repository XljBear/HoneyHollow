package config

import (
	"os"
	"strconv"
)

var MongoDBName = ""
var MongoDBHost = ""
var MongoDBPort = ""
var MongoDBUser = ""
var MongoDBPassword = ""

var RedisHost = ""
var RedisPort = ""
var RedisPassword = ""
var RedisDBIndex = 0

var RunMode = ""

func initMongoDBConfig() {
	MongoDBName = os.Getenv("MongoDBName")
	if MongoDBName == "" {
		MongoDBName = "HoneyHollow"
	}

	MongoDBHost = os.Getenv("MongoDBHost")
	if MongoDBHost == "" {
		MongoDBHost = "127.0.0.1"
	}

	MongoDBPort = os.Getenv("MongoDBPort")
	if MongoDBPort == "" {
		MongoDBPort = "27017"
	}

	MongoDBUser = os.Getenv("MongoDBUser")

	MongoDBPassword = os.Getenv("MongoDBPassword")
}

func initSystemConfig() {
	RunMode = os.Getenv("RunMode")
	if RunMode == "" {
		RunMode = "prod"
	}
}

func intRedisConfig() {
	RedisHost = os.Getenv("RedisHost")
	if RedisHost == "" {
		RedisHost = "127.0.0.1"
	}

	RedisPort = os.Getenv("RedisPort")
	if RedisPort == "" {
		RedisPort = "6379"
	}

	RedisPassword = os.Getenv("RedisPassword")

	RedisDBIndexStr := os.Getenv("RedisDBIndex")
	if RedisDBIndexStr == "" {
		RedisDBIndex = 0
	} else {
		dbIndex, err := strconv.Atoi(RedisDBIndexStr)
		if err != nil || dbIndex < 0 {
			RedisDBIndex = 0
		} else {
			RedisDBIndex = dbIndex
		}
	}
}

func InitConfig() {
	initSystemConfig()
	initMongoDBConfig()
	intRedisConfig()
}
