package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/Golang-Coach/Scheduler/mocks"
	"gopkg.in/mgo.v2/bson"
	"github.com/Golang-Coach/Scheduler/models"
	"time"
	"github.com/pkg/errors"
)

func TestSchedule(t *testing.T) {

	Convey("Should update latest packages", t, func() {
		githubService := new(mocks.IGithub)
		dataStore := new(mocks.IDataStore)
		repositoryInfo := models.RepositoryInfo{
			RepoName: "react",
		}
		repos := []models.RepositoryInfo{repositoryInfo}
		githubService.On("GetUpdatedRepositoryInfo", &repositoryInfo).Return(&repositoryInfo, nil)
		dataStore.On("FindPackageWithinLimit", bson.M{}, 0, 500).Return(&repos, nil)
		dataStore.On("UpdatePackage", &repositoryInfo).Return(nil)
		Schedule(dataStore, githubService)
		So(githubService.AssertCalled(t, "GetUpdatedRepositoryInfo", &repositoryInfo), ShouldBeTrue)
		So(dataStore.AssertCalled(t, "FindPackageWithinLimit", bson.M{}, 0, 500), ShouldBeTrue)
		So(dataStore.AssertCalled(t, "UpdatePackage", &repositoryInfo), ShouldBeTrue)
	})

	Convey("Should not update latest package if there is not change in package", t, func() {
		githubService := new(mocks.IGithub)
		dataStore := new(mocks.IDataStore)
		repositoryInfo := models.RepositoryInfo{
			RepoName: "react",
		}
		repos := []models.RepositoryInfo{repositoryInfo}
		githubService.On("GetUpdatedRepositoryInfo", &repositoryInfo).Return(nil, nil)
		dataStore.On("FindPackageWithinLimit", bson.M{}, 0, 500).Return(&repos, nil)
		Schedule(dataStore, githubService)
		So(githubService.AssertCalled(t, "GetUpdatedRepositoryInfo", &repositoryInfo), ShouldBeTrue)
		So(dataStore.AssertCalled(t, "FindPackageWithinLimit", bson.M{}, 0, 500), ShouldBeTrue)
		So(dataStore.AssertNotCalled(t, "UpdatePackage", &repositoryInfo), ShouldBeTrue)
	})

	Convey("Should not update latest package because of timeout", t, func() {
		githubService := new(mocks.IGithub)
		dataStore := new(mocks.IDataStore)
		repositoryInfo := models.RepositoryInfo{
			RepoName: "react",
		}
		repos := []models.RepositoryInfo{repositoryInfo}
		githubService.On("GetUpdatedRepositoryInfo", &repositoryInfo).WaitUntil(time.After(6 * time.Second)).Return(&repositoryInfo, nil)
		dataStore.On("FindPackageWithinLimit", bson.M{}, 0, 500).Return(&repos, nil)
		Schedule(dataStore, githubService)
		So(githubService.AssertCalled(t, "GetUpdatedRepositoryInfo", &repositoryInfo), ShouldBeTrue)
		So(dataStore.AssertCalled(t, "FindPackageWithinLimit", bson.M{}, 0, 500), ShouldBeTrue)
		So(dataStore.AssertNotCalled(t, "UpdatePackage", &repositoryInfo), ShouldBeTrue)
	})

	Convey("Should not update package if there is an error", t, func() {
		githubService := new(mocks.IGithub)
		dataStore := new(mocks.IDataStore)
		repositoryInfo := models.RepositoryInfo{
			RepoName: "react",
		}
		repos := []models.RepositoryInfo{repositoryInfo}
		githubService.On("GetUpdatedRepositoryInfo", &repositoryInfo).Return(&repositoryInfo, errors.New("Some problem"))
		dataStore.On("FindPackageWithinLimit", bson.M{}, 0, 500).Return(&repos, nil)
		dataStore.On("UpdatePackage", &repositoryInfo).Return(nil)
		Schedule(dataStore, githubService)
		So(githubService.AssertCalled(t, "GetUpdatedRepositoryInfo", &repositoryInfo), ShouldBeTrue)
		So(dataStore.AssertCalled(t, "FindPackageWithinLimit", bson.M{}, 0, 500), ShouldBeTrue)
		So(dataStore.AssertNotCalled(t, "UpdatePackage", &repositoryInfo), ShouldBeTrue)
	})

	Convey("Should not process if failed to retrieve package information", t, func() {
		githubService := new(mocks.IGithub)
		dataStore := new(mocks.IDataStore)
		repositoryInfo := models.RepositoryInfo{
			RepoName: "react",
		}
		repos := []models.RepositoryInfo{repositoryInfo}
		dataStore.On("FindPackageWithinLimit", bson.M{}, 0, 500).Return(&repos, errors.New("Some problem"))
		Schedule(dataStore, githubService)
		So(githubService.AssertNotCalled(t, "GetUpdatedRepositoryInfo", &repositoryInfo), ShouldBeTrue)
		So(dataStore.AssertCalled(t, "FindPackageWithinLimit", bson.M{}, 0, 500), ShouldBeTrue)
		So(dataStore.AssertNotCalled(t, "UpdatePackage", &repositoryInfo), ShouldBeTrue)
	})

}
