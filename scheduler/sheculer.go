package scheduler

import (
	"github.com/Golang-Coach/Scheduler/interfaces"
	"github.com/Golang-Coach/Scheduler/models"
	"github.com/globalsign/mgo/bson"
	"time"
	"log"
)

type GithubResponse struct {
	RepositoryInfo *models.RepositoryInfo
	err            error
}

func Schedule(dataStore interfaces.IRepositoryStore, githubService interfaces.IGithub) {
	// get data from database
	repositories, err := dataStore.FindPackageWithinLimit(bson.M{}, 0, 500)
	if err != nil {
		log.Fatalln("Error", err)
		return
	}

	responses := updatePackage(*repositories, githubService)

	for _, response := range responses {
		if response.err != nil {
			log.Fatalln("Error", response.err)
		} else if response.RepositoryInfo != nil {
			dataStore.UpdatePackage(response.RepositoryInfo)
		}
	}
}

func updatePackage(repositories []models.RepositoryInfo, githubService interfaces.IGithub) []*GithubResponse {
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
			log.Fatalln("Timeout")
			return responses
		}
	}
	return responses
}
