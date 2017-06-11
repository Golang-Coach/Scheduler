package services

import (
	"github.com/Golang-Coach/Scheduler/models"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type ICollection interface {
	Insert(...interface{}) error
	Update(selector interface{}, update interface{}) error
	Find(query interface{}) *mgo.Query
}

type IDataStore interface {
	AddPackage(repositoryInfo models.RepositoryInfo) error
	UpdatePackage(pack *models.RepositoryInfo) error
	FindPackage(query interface{}) (*models.RepositoryInfo, error)
	FindPackageWithinLimit(query interface{}, skip int, limit int) (*[]models.RepositoryInfo, error)
}

type DataStore struct {
	collection ICollection
}

func NewDataStore(collection ICollection) IDataStore {
	return DataStore{
		collection: collection,
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

func (store DataStore) FindPackage(query interface{}) (*models.RepositoryInfo, error) {
	// find package with limit
	repositoryInfo := &models.RepositoryInfo{}
	err := store.collection.Find(query).All(repositoryInfo)
	return repositoryInfo, err;
}

func (store DataStore) FindPackageWithinLimit(query interface{}, skip int, limit int) (*[]models.RepositoryInfo, error) {
	// find package with limit
	repositoryInfos := &[]models.RepositoryInfo{}
	result := store.collection.Find(query)
	if limit > 0 {
		result = result.Limit(limit)
	}

	if skip > 0 {
		result = result.Skip(skip)
	}
	err :=  result.All(repositoryInfos)
	return repositoryInfos, err
}
