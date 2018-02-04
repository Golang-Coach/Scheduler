package services_test

import (
	"testing"
	"github.com/Golang-Coach/Scheduler/mocks"
	"github.com/Golang-Coach/Scheduler/models"
	"github.com/Golang-Coach/Scheduler/services"
	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDataStore(t *testing.T) {

	Convey("Should add repository information", t, func() {
		collection := new(mocks.ICollection)
		dataStore := new(mocks.IDataStore)
		dataStore.On("GetCollection").Return(collection)
		repositoryInfo := models.RepositoryInfo{}
		collection.On("Insert", repositoryInfo).Return(nil)
		target := services.NewRepositoryStore(dataStore)
		err := target.AddPackage(repositoryInfo)
		So(err, ShouldBeNil)
	})

	Convey("Should update repository information", t, func() {
		collection := new(mocks.ICollection)
		dataStore := new(mocks.IDataStore)
		dataStore.On("GetCollection").Return(collection)
		id := bson.ObjectId("AA")
		repositoryInfo := &models.RepositoryInfo{ID: id}
		collection.On("Update", bson.M{"_id": id}, repositoryInfo).Return(nil)
		target := services.NewRepositoryStore(dataStore)
		err := target.UpdatePackage(repositoryInfo)
		So(err, ShouldBeNil)
	})
}
