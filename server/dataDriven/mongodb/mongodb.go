package mongodb

import (
	"HoneyHollow/server/config"
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
)

var dbs map[string]*qmgo.Database

func GetDb(dbName string) (*qmgo.Database, error) {
	if dbs == nil {
		dbs = make(map[string]*qmgo.Database)
	} else {
		if db, exist := dbs[dbName]; exist {
			return db, nil
		}
	}
	ctx := context.Background()
	mHost := config.MongoDBHost
	mPort := config.MongoDBPort
	mUser := config.MongoDBUser
	mPassword := config.MongoDBPassword
	var client *qmgo.Client
	var mongoUri string
	if mUser == "" || mPassword == "" {
		mongoUri = fmt.Sprintf("mongodb://%s:%s", mHost, mPort)
	} else {
		mongoUri = fmt.Sprintf("mongodb://%s:%s@%s:%s", mUser, mPassword, mHost, mPort)
	}
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoUri})
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	dbs[dbName] = db
	return db, nil
}

func GetCollection(dbName string, collection string) (*qmgo.Collection, error) {
	db, err := GetDb(dbName)
	if err != nil {
		return nil, err
	}
	return db.Collection(collection), nil
}
