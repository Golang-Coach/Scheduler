package scheduler

import (
	"github.com/Golang-Coach/Scheduler/models"
	"github.com/Golang-Coach/Scheduler/services"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)

type GithubResponse struct {
	RepositoryInfo *models.RepositoryInfo
	err            error
}

func Schedule(dataStore services.IDataStore, githubService services.IGithub) {
	// get data from database
	repositories, err := dataStore.FindPackageWithinLimit(bson.M{}, 0, 500)
	if err != nil {
		fmt.Println("Error")
		return
	}

	responses := updatePackage(*repositories, githubService)

	for _, response := range responses {
		if response.err != nil {
			fmt.Println(response.err)
		} else if response.RepositoryInfo != nil {
			dataStore.UpdatePackage(response.RepositoryInfo)
		}
	}
}

func updatePackage(repositories []models.RepositoryInfo, githubService services.IGithub) []*GithubResponse {
	ch := make(chan *GithubResponse)
	responses := []*GithubResponse{}

	for _, repository := range repositories {
		go func(repository models.RepositoryInfo) {
			repoInfo, err := githubService.GetUpdatedRepositoryInfo(repository)
			ch <- &GithubResponse{
				RepositoryInfo: repoInfo,
				err:            err,
			}
		}(repository)
	}

	for {
		select {
		case res := <-ch:
			responses = append(responses, res)
			if len(responses) == len(repositories) {
				return responses
			}
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			return responses;
		}
	}
	return responses
}
