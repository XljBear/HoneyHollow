package models

import (
	"HoneyHollow/server/config"
	"HoneyHollow/server/dataDriven/mongodb"
	"github.com/qiniu/qmgo"
)

type CollectionBase struct {
	CollectionName string
}

func (cb *CollectionBase) GetCollectionName() string {
	return cb.CollectionName
}

func (cb *CollectionBase) GetCollection() *qmgo.Collection {
	c, err := mongodb.GetCollection(config.MongoDBName, cb.CollectionName)
	if err != nil {
		//TODO 细化异常Log逻辑
		panic(err)
	}
	return c
}
