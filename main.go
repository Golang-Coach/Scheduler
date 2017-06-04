package main

import (
	"fmt"
	"context"
	"golang.org/x/oauth2"
	. "github.com/google/go-github/github"
	"Scheduler/services"
)

func main() {

	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "ff24cdf80f561cd058713b6c647b37edd0cdb0b6"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)
	client := *NewClient(tokenClient)
	githubApi := services.NewGithub(&client, client.Repositories, context)
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
