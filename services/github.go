package services

import (
	"github.com/google/go-github/github"
	"context"
	"github.com/Golang-Coach/Scheduler/models"
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
	GetPackageRepoInfo(owner string, repositoryName string) (*models.Package, error)
	GetLastCommitInfo(owner string, repositoryName string) (*github.RepositoryCommit, error)
	GetReadMe(owner string, repositoryName string) (string, error)
	GetRateLimitInfo(owner string, repositoryName string) (*github.RateLimits, error)
}

type IClient interface {
	RateLimits(ctx context.Context) (*github.RateLimits, *github.Response, error)
}

type Github struct {
	client  IClient
	repositoryServices IRepositoryServices
	context context.Context
}


func NewGithub(client IClient, repositoryServices IRepositoryServices, context context.Context) Github {
	return Github{
		client:  client,
		repositoryServices: repositoryServices,
		context: context,
	}
}

func (service *Github) GetPackageRepoInfo(owner string, repositoryName string) (*models.Package, error) {
	repo, _, err := service.repositoryServices.Get(service.context, owner, repositoryName)
	if err != nil {
		return nil, err
	}
	pack := &models.Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:   *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}
	return pack, nil
}

func (service *Github) GetLastCommitInfo(owner string, repositoryName string) (*github.RepositoryCommit, error) {
	commitInfo, _, err := service.repositoryServices.ListCommits(service.context, owner, repositoryName, nil)
	if err != nil {
		return nil, err
	}
	return commitInfo[0], nil
}

func (service *Github) GetReadMe(owner string, repositoryName string) (string, error) {
	readme, _, err := service.repositoryServices.GetReadme(service.context, owner, repositoryName, nil)
	if err != nil {
		return "", err
	}

	// get content
	return readme.GetContent()
}

func (service *Github) GetRateLimitInfo() (*github.RateLimits, error) {
	rateLimitInfo, _, err := service.client.RateLimits(service.context)
	return rateLimitInfo, err
}
