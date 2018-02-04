package db

import (
	"github.com/globalsign/mgo"
	"github.com/Golang-Coach/Scheduler/interfaces"
)

type DataStore struct {
	Session *mgo.Session
}

func (store DataStore) GetCollection() interfaces.ICollection {
	// get collection
	collection := store.getDatabase().C("repositories")
	return collection
}

func (store DataStore) getDatabase() *mgo.Database {
	session := store.Session.Clone()
	database := &mgo.Database{session, "golang-coach"} /// TODO get this from environment variable
	return database
}

func (store *DataStore) EnsureConnected() {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("Ping session")
			//store.session.Ping()
			//Your reconnect logic here.
		}
	}()

	//Ping panics if session is closed. (see mgo.Session.Panic())
	store.Session.Ping()
}
