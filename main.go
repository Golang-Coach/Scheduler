package main

import (
	"fmt"
	"os"
	"gopkg.in/mgo.v2"
	"time"
	//"net"
	//"crypto/tls"
	"github.com/Golang-Coach/Scheduler/services"
	"context"
	"github.com/google/go-github/github"


	//"log"
	//"github.com/Golang-Coach/Scheduler/models"
	//"golang.org/x/oauth2"
)

func main() {


	// TODO -- this is to create Github service
	backgroundContext := context.Background()
	//tokenService := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: "22ffe92b14c28bf8ec53e7f0102ed240c1e02633"},
	//)
	//tokenClient := oauth2.NewClient(backgroundContext, tokenService)
	//client := *github.NewClient(tokenClient)
	client := *github.NewClient(nil)
	githubApi := services.NewGithub(&client, client.Repositories, backgroundContext)
	pack, err := githubApi.GetLastCommitInfo("Golang-Coach", "Lessons")
	//packa := models.RepositoryInfo{
	//	UpdatedAt: pack.Commit.Committer.GetDate(),
	//}
	fmt.Println(pack.Commit.Committer.GetDate())
	//fmt.Println(packa)
	//fmt.Println(err);

	// TODO -- this is used to connect to MongoDB
	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"}, // Get HOST + PORT
		Timeout:  5 * time.Second,
		//Database: "golancoach",                                                                             // It can be anything
		//Username: "coach",                                                                             // Username
		//Password: "Pa55word", // PASSWORD
		//DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
		//	return tls.Dial("tcp", addr.String(), &tls.Config{})
		//},
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
	collection := session.DB("golang-couch").C("package")

	 //insert Document in collection
	//err = collection.Insert(&models.RepositoryInfo{
	//	RepoName:"react",
	//	Owner:"facebook",
	//	//FullName:"react",
	//	//Description:"A framework for building native apps with React.",
	//	//ForksCount: 11392,
	//	//StarsCount:48794,
	//	//LastUpdatedBy:"shergin",
	//
	//})
	//
	//if err != nil {
	//	log.Fatal("Problem inserting data: ", err)
	//	return
	//}
	dataStore := services.NewDataStore(collection)
	Schedule(dataStore, githubApi)
	//Schedule(dataStore, githubApi)
	//fmt.Println(collection)
	//fmt.Println(githubApi.GetRateLimitInfo())

	// Get Document from collection
	//result := models.RepositoryInfo{}
	//err = collection.Find(bson.M{"fullname": "react"}).One(&result)
	//if err != nil {
	//	log.Fatal("Error finding record: ", err)
	//	return
	//}
	//
	//fmt.Println("Description:", result.Description)
	//
	//// update document
	//updateQuery := bson.M{"_id": result.ID}
	//change := bson.M{"$set": bson.M{"fullname": "react-native"}}
	//err = collection.Update(updateQuery, change)
	//if err != nil {
	//	log.Fatal("Error updating record: ", err)
	//	return
	//}

	// delete document
	//err = collection.Remove(updateQuery)
	//if err != nil {
	//	log.Fatal("Error deleting record: ", err)
	//	return
	//}

	//c := cron.New()
	//c.AddFunc("@every 1s", func() { fmt.Println("Every hour thirty") })
	//c.Start()
	//
	//pack := models.RepositoryInfo{}
	//fmt.Printf("%+v", pack)
	//
	//// wait forever
	//<-make(chan int)

}
