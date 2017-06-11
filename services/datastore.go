package services

import (
	"gopkg.in/mgo.v2"
	"github.com/Golang-Coach/Scheduler/models"
	"gopkg.in/mgo.v2/bson"
)

type ICollection interface {
	Insert(...interface {}) error
	Update(selector interface{}, update interface{}) error
	Find(query interface{}) *mgo.Query
}

type IDataStore interface {
	AddPackage(pack models.RepositoryInfo) error
	UpdatePackage(pack *models.RepositoryInfo) error
	FindPackage(query interface{}) *mgo.Query
	FindPackageWithinLimit(query interface{}, skip int, limit int) *mgo.Query
}

type DataStore struct {
	collection ICollection
}

func NewDataStore(collection ICollection) IDataStore{
	return DataStore{
		collection:collection,
	}
}

func (store DataStore) AddPackage(repositoryInfo models.RepositoryInfo) error {
	// insert Document in collection
	return store.collection.Insert(repositoryInfo)
}

func (store DataStore) UpdatePackage(repositoryInfo *models.RepositoryInfo) error {
	// update Document in collection
	return store.collection.Update(bson.M{"_id": repositoryInfo.ID}, repositoryInfo)
}

func (store DataStore) FindPackage(query interface{}) *mgo.Query {
	// find package with limit
	return store.collection.Find(query)
}

func (store DataStore) FindPackageWithinLimit(query interface{}, skip int, limit int) *mgo.Query {
	// find package with limit
	result := store.collection.Find(query)
	if limit > 0 {
		return result.Limit(limit)
	}

	if skip > 0 {
		return result.Skip(skip)
	}
	return result
}



