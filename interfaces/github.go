package interfaces

import (
	"github.com/google/go-github/github"
	"context"
	"github.com/Golang-Coach/Scheduler/models"
)

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
	GetUpdatedRepositoryInfo(repositoryInfo models.RepositoryInfo) (*models.RepositoryInfo, error)
}