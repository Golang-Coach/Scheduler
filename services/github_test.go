package services_test

import (
	"Scheduler/mocks"
	. "Scheduler/services"
	. "github.com/onsi/ginkgo"
	"context"
	. "github.com/google/go-github/github"
	. "github.com/onsi/gomega"
	"encoding/base64"
)

var _ = Describe("Github API ", func() {
	var (
		client             mocks.IClient
		repositoryServices *mocks.IRepositoryServices
		github             Github
		backgroundContext  context.Context
	)

	BeforeEach(func() {
		backgroundContext = context.Background()
		repositoryServices = new(mocks.IRepositoryServices)
		client = mocks.IClient{
			Repositories: repositoryServices,
		}
		github = NewGithub(&client, repositoryServices, backgroundContext)
	})

	It("should get repository information", func() {
		fullName := "ABC"
		starCount := 10
		repo := &Repository{
			FullName:        &fullName,
			Description:     &fullName,
			ForksCount:      &starCount,
			StargazersCount: &starCount,

		}
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)
		pack, _ := github.GetPackageRepoInfo("golang-coach", "Lessons")
		Ω(pack.ForksCount).Should(Equal(starCount))
	})

	It("should get GetReadMe information", func() {
		content := "ABC"
		encodedContent := base64.StdEncoding.EncodeToString([]byte(content))
		repositoryContent := &RepositoryContent{
			Content: &encodedContent,
		}
		repositoryServices.On("GetReadme", backgroundContext, "golang-coach", "Lessons", (*RepositoryContentGetOptions)(nil)).Return(repositoryContent, nil, nil)
		readme, _ := github.GetReadMe("golang-coach", "Lessons")
		Ω(readme).Should(Equal(encodedContent))
	})

	It("should get last commit information", func() {
		repositoryCommit := new(RepositoryCommit)
		repositoryServices.On("ListCommits", backgroundContext, "golang-coach", "Lessons",
			(*CommitsListOptions)(nil)).Return([]*RepositoryCommit{repositoryCommit}, nil, nil)
		commitInfo, _ := github.GetLastCommitInfo("golang-coach", "Lessons")
		Ω(commitInfo).Should(Equal(repositoryCommit))
	})

	It("should get rate limit information", func() {
		rateLimit := new(RateLimits)
		client.On("RateLimits", backgroundContext).Return(rateLimit, nil, nil)
		rateLimitInfo, _ := github.GetRateLimitInfo()
		Ω(rateLimitInfo).Should(Equal(rateLimit))
	})
})
