package services

import (
	"github.com/Golang-Coach/Scheduler/models"
	"github.com/globalsign/mgo/bson"
	"github.com/Golang-Coach/Scheduler/interfaces"
)

type RepositoryStore struct {
	dataStore interfaces.IDataStore
}

func NewRepositoryStore(dataStore interfaces.IDataStore) interfaces.IRepositoryStore {
	return RepositoryStore{
		dataStore,
	}
}

func (s RepositoryStore) AddPackage(repositoryInfo models.RepositoryInfo) error {
	// insert Document in collection
	return s.dataStore.GetCollection().Insert(repositoryInfo)
}

func (s RepositoryStore) UpdatePackage(repositoryInfo *models.RepositoryInfo) error {
	// update Document in collection
	return s.dataStore.GetCollection().Update(bson.M{"_id": repositoryInfo.ID}, repositoryInfo)
}

func (s RepositoryStore) FindPackage(query interface{}) (*models.RepositoryInfo, error) {
	// find package with limit
	repositoryInfo := &models.RepositoryInfo{}
	err := s.dataStore.GetCollection().Find(query).All(repositoryInfo)
	return repositoryInfo, err
}

func (s RepositoryStore) FindPackageWithinLimit(query interface{}, skip int, limit int) (*[]models.RepositoryInfo, error) {
	// find package with limit
	repositoryInfos := &[]models.RepositoryInfo{}
	result := s.dataStore.GetCollection().Find(query)
	if limit > 0 {
		result = result.Limit(limit)
	}

	if skip > 0 {
		result = result.Skip(skip)
	}
	err := result.All(repositoryInfos)
	return repositoryInfos, err
}
