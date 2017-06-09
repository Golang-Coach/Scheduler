package main

import (
	"fmt"
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"github.com/Golang-Coach/Scheduler/services"
)

func main() {

	backgroundContext := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "22ffe92b14c28bf8ec53e7f0102ed240c1e02633"},
	)
	tokenClient := oauth2.NewClient(backgroundContext, tokenService)
	client := *github.NewClient(tokenClient)
	githubApi := services.NewGithub(&client, client.Repositories, backgroundContext)
	pack, err := githubApi.GetPackageRepoInfo("Golang-coach", "Lessons")
	fmt.Println(pack)
	fmt.Println(err);
	//c := cron.New()
	//c.AddFunc("@every 1s", func() { fmt.Println("Every hour thirty") })
	//c.Start()
	//
	//pack := models.Package{}
	//fmt.Printf("%+v", pack)
	//
	//// wait forever
	//<-make(chan int)

}
