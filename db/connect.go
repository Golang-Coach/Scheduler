package db

import (
	"crypto/tls"
	"fmt"
	"github.com/globalsign/mgo"
	"net"
	"os"
	//"time"
	//"time"
)

func dialServer(addr *mgo.ServerAddr) (net.Conn, error) {
	return tls.Dial("tcp", addr.String(), &tls.Config{})
}

func Connect() *DataStore {
	// TODO -- this is used to connect to MongoDB
	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo, err := mgo.ParseURL(os.Getenv("connection-string"))
	dialInfo.DialServer = dialServer

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	// SetSafe changes the session safety mode.
	// If the safe parameter is nil, the session is put in unsafe mode, and writes become fire-and-forget,
	// without error checking. The unsafe mode is faster since operations won't hold on waiting for a confirmation.
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode.
	session.SetSafe(&mgo.Safe{})

	fmt.Println(session.DatabaseNames())

	dataStore := &DataStore{session}
	dataStore.EnsureConnected()

	return dataStore
}
