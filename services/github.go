package services

import (
	"github.com/google/go-github/github"
	"context"
	"github.com/Golang-Coach/Scheduler/models"
	"strings"
	"errors"
)

type IRepositoryContent interface {
	GetContent() (string, error)
}

type IRepositoryServices interface {
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
	ListCommits(ctx context.Context, owner, repo string, opt *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	GetReadme(ctx context.Context, owner, repo string, opt *github.RepositoryContentGetOptions) (*github.RepositoryContent, *github.Response, error)
}

type IGithub interface {
	GetRepositoryInfo(owner string, repositoryName string) (*models.RepositoryInfo, error)
	GetLastCommitInfo(owner string, repositoryName string) (*github.RepositoryCommit, error)
	GetReadMe(owner string, repositoryName string) (string, error)
	GetRateLimitInfo() (*github.RateLimits, error)
	GetUpdatedRepositoryInfo(repositoryInfo *models.RepositoryInfo) (*models.RepositoryInfo, error)
}

type IClient interface {
	RateLimits(ctx context.Context) (*github.RateLimits, *github.Response, error)
}

type Github struct {
	client             IClient
	repositoryServices IRepositoryServices
	context            context.Context
}

func NewGithub(client IClient, repositoryServices IRepositoryServices, context context.Context) Github {
	return Github{
		client:             client,
		repositoryServices: repositoryServices,
		context:            context,
	}
}

func (service Github) GetRepositoryInfo(owner string, repositoryName string) (*models.RepositoryInfo, error) {
	repo, _, err := service.repositoryServices.Get(service.context, owner, repositoryName)
	if err != nil {
		return nil, err
	}
	repositoryInfo := &models.RepositoryInfo{
		RepoName: *repo.Name,
		Owner:      strings.Split(*repo.FullName, "/")[0] ,
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}
	return repositoryInfo, nil
}

func (service Github) GetLastCommitInfo(owner string, repositoryName string) (*github.RepositoryCommit, error) {
	commitInfo, _, err := service.repositoryServices.ListCommits(service.context, owner, repositoryName, nil)
	if err != nil {
		return nil, err
	}
	return commitInfo[0], nil
}

func (service Github) GetReadMe(owner string, repositoryName string) (string, error) {
	readme, _, err := service.repositoryServices.GetReadme(service.context, owner, repositoryName, nil)
	if err != nil {
		return "", err
	}

	// get content
	return readme.GetContent()
}

func (service Github) GetRateLimitInfo() (*github.RateLimits, error) {
	rateLimitInfo, _, err := service.client.RateLimits(service.context)
	return rateLimitInfo, err
}

func (service Github) GetUpdatedRepositoryInfo(repositoryInfo *models.RepositoryInfo) (*models.RepositoryInfo, error) {
	// Call last update information from Github API
	if len(strings.TrimSpace(repositoryInfo.Owner)) == 0 || len(strings.TrimSpace(repositoryInfo.RepoName)) == 0 {
		return nil, errors.New("Repository Name is incorrect")
	}

	// check with existing caller api
	lastCommitInfo, err := service.GetLastCommitInfo(repositoryInfo.Owner, repositoryInfo.RepoName)
	if err != nil {
		return nil, err
	}

	if lastCommitInfo.Commit.Committer.GetDate().Equal(repositoryInfo.UpdatedAt) {
		return nil, nil
	}

	newRepositoryInfo, err := service.GetRepositoryInfo(repositoryInfo.Owner, repositoryInfo.RepoName)
	if err != nil {
		return nil, err
	}

	newRepositoryInfo.ID = repositoryInfo.ID
	newRepositoryInfo.UpdatedAt = lastCommitInfo.Commit.Committer.GetDate()
	newRepositoryInfo.LastUpdatedBy = lastCommitInfo.Commit.Committer.GetName()

	content, err := service.GetReadMe(repositoryInfo.Owner, repositoryInfo.RepoName)
	if err != nil {
		return nil, err
	}
	newRepositoryInfo.ReadMe = content

	// data is same ignore, else update data
	return newRepositoryInfo, nil
}
