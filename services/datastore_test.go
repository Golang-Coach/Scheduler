package services

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/Golang-Coach/Scheduler/mocks"
	"github.com/Golang-Coach/Scheduler/models"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func TestDataStore(t *testing.T) {

	Convey("Should add repository information", t, func() {
		collection := new(mocks.ICollection)
		repositoryInfo := models.RepositoryInfo{}
		collection.On("Insert", repositoryInfo).Return(nil)
		target := NewDataStore(collection)
		err := target.AddPackage(repositoryInfo)
		So(err, ShouldBeNil)
	})

	Convey("Should update repository information", t, func() {
		collection := new(mocks.ICollection)
		id := bson.ObjectId("AA")
		repositoryInfo := &models.RepositoryInfo{ID: id}
		collection.On("Update", bson.M{"_id": id}, repositoryInfo).Return(nil)
		target := NewDataStore(collection)
		err := target.UpdatePackage(repositoryInfo)
		So(err, ShouldBeNil)
	})

	Convey("Should find repository by query", t, func() {
		collection := new(mocks.ICollection)
		id := bson.ObjectId("AA")
		query := &mgo.Query{}
		collection.On("Find", bson.M{"_id": id}).Return(query)
		target := NewDataStore(collection)
		result := target.FindPackage(bson.M{"_id": id})
		So(result, ShouldEqual, query)
	})

	Convey("Should find repository by query and limit information", t, func() {
		collection := new(mocks.ICollection)
		id := bson.ObjectId("AA")
		query := &mgo.Query{}
		collection.On("Find", bson.M{"_id": id}).Return(query)
		target := NewDataStore(collection)
		result := target.FindPackageWithinLimit(bson.M{"_id": id}, 10, 10)
		So(result, ShouldEqual, query)
	})

}