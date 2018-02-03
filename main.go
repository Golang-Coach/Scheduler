package main

import (
	"context"
	"fmt"
	"github.com/Golang-Coach/Scheduler/scheduler"
	"github.com/Golang-Coach/Scheduler/services"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2"
	"os"
	"time"
	"net"
	"crypto/tls"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() {
	backgroundContext := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("github_token")},
	)
	tokenClient := oauth2.NewClient(backgroundContext, tokenService)
	client := *github.NewClient(tokenClient)
	githubApi := services.NewGithub(&client, client.Repositories, backgroundContext)

	// TODO -- this is used to connect to MongoDB
	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{os.Getenv("database_hostname")}, // Get HOST + PORT
		Timeout:  5 * time.Second,
		Database: os.Getenv("database_name"),     // It can be anything
		Username: os.Getenv("database_username"), // Username
		Password: os.Getenv("database_password"),
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer session.Close()

	// SetSafe changes the session safety mode.
	// If the safe parameter is nil, the session is put in unsafe mode, and writes become fire-and-forget,
	// without error checking. The unsafe mode is faster since operations won't hold on waiting for a confirmation.
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode.
	session.SetSafe(&mgo.Safe{})

	// get collection
	collection := session.DB("golang-couch").C("repositories")

	dataStore := services.NewDataStore(collection)
	scheduler.Schedule(dataStore, githubApi)
}

func main() {
	environment := os.Getenv("environment")
	if environment != "production" {
		HandleRequest()
	} else {
		lambda.Start(HandleRequest)
	}

}
