package main

import (
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"Scheduler/models"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { fmt.Println("Every hour thirty") })
	c.Start()

	pack := models.Package{}
	fmt.Printf("%+v", pack)

	// wait forever
	<-make(chan int)

}
