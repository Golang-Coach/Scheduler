package services

import (
	"github.com/Golang-Coach/Scheduler/mocks"
	"context"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/google/go-github/github"
	"testing"
	"errors"
	"encoding/base64"
)

func TestGithubAPI(t *testing.T) {
	Convey("GetRepositoryInfo", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		client := new(mocks.IClient)
		githubService := NewGithub(client, repositoryServices, backgroundContext)
		Convey("Should return repository information", func() {
			fullName := "facebook/react"
			starCount := 10
			repo := &Repository{
				Name:            &fullName,
				FullName:        &fullName,
				Description:     &fullName,
				ForksCount:      &starCount,
				StargazersCount: &starCount,

			}
			repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)
			repositoryInfo, _ := githubService.GetRepositoryInfo("golang-coach", "Lessons")
			So(repositoryInfo.ForksCount, ShouldEqual, starCount)
		})

		Convey("Should return error when failed to retrieve  repository information", func() {
			repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(nil, nil, errors.New("Error has been occurred"))
			_, err := githubService.GetRepositoryInfo("golang-coach", "Lessons")
			So(err, ShouldNotBeEmpty)
		})
	})

	Convey("GetReadMe", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		client := new(mocks.IClient)
		githubService := NewGithub(client, repositoryServices, backgroundContext)
		Convey("should get return repository readme information", func() {
			content := "ABC"
			encodedContent := base64.StdEncoding.EncodeToString([]byte(content))
			repositoryContent := &RepositoryContent{
				Content: &encodedContent,
			}
			repositoryServices.On("GetReadme", backgroundContext, "golang-coach", "Lessons", (*RepositoryContentGetOptions)(nil)).Return(repositoryContent, nil, nil)
			readme, _ := githubService.GetReadMe("golang-coach", "Lessons")
			So(readme, ShouldEqual, encodedContent)
		})

		Convey("Should return error when failed to retrieve  repository readme information", func() {
			repositoryServices.On("GetReadme", backgroundContext, "golang-coach", "Lessons",
				(*RepositoryContentGetOptions)(nil)).Return(nil, nil, errors.New("Error has been occurred"))
			_, err := githubService.GetReadMe("golang-coach", "Lessons")
			So(err, ShouldNotBeEmpty)
		})
	})

	Convey("GetLastCommitInfo", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		client := new(mocks.IClient)
		githubService := NewGithub(client, repositoryServices, backgroundContext)
		Convey("should should return last commit information", func() {
			repositoryCommit := new(RepositoryCommit)
			repositoryServices.On("ListCommits", backgroundContext, "golang-coach", "Lessons",
				(*CommitsListOptions)(nil)).Return([]*RepositoryCommit{repositoryCommit}, nil, nil)
			commitInfo, _ := githubService.GetLastCommitInfo("golang-coach", "Lessons")
			So(commitInfo, ShouldEqual, repositoryCommit)
		})

		Convey("Should return error when failed to retrieve  repository readme information", func() {
			repositoryServices.On("ListCommits", backgroundContext, "golang-coach", "Lessons",
				(*CommitsListOptions)(nil)).Return(nil, nil, errors.New("Error has been occurred"))
			_, err := githubService.GetLastCommitInfo("golang-coach", "Lessons")
			So(err, ShouldNotBeEmpty)
		})
	})

	Convey("GetRateLimitInfo", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		client := new(mocks.IClient)
		githubService := NewGithub(client, repositoryServices, backgroundContext)
		Convey("should should return rate limit information", func() {
			rateLimit := new(RateLimits)
			client.On("RateLimits", backgroundContext).Return(rateLimit, nil, nil)
			rateLimitInfo, _ := githubService.GetRateLimitInfo()
			So(rateLimitInfo, ShouldEqual, rateLimit)
		})

		Convey("Should return error when failed to retrieve  repository readme information", func() {
			client.On("RateLimits", backgroundContext).Return(nil, nil, errors.New("Error has been occurred"))
			_, err := githubService.GetRateLimitInfo()
			So(err, ShouldNotBeEmpty)
		})
	})

}
