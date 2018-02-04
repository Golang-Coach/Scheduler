package main

import (
	"context"
	"github.com/Golang-Coach/Scheduler/scheduler"
	"github.com/Golang-Coach/Scheduler/services"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/Golang-Coach/Scheduler/db"
)

func HandleRequest() {
	backgroundContext := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("github_token")},
	)
	tokenClient := oauth2.NewClient(backgroundContext, tokenService)
	client := *github.NewClient(tokenClient)
	githubApi := services.NewGithub(&client, client.Repositories, backgroundContext)

	dataStore := db.Connect()
	defer dataStore.Session.Close()

	repositoriesStore := services.NewRepositoryStore(dataStore)
	scheduler.Schedule(repositoriesStore, githubApi)
}

func main() {
	environment := os.Getenv("environment")
	if environment != "production" {
		HandleRequest()
	} else {
		lambda.Start(HandleRequest)
	}

}
