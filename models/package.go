package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type RepositoryInfo struct{
	ID bson.ObjectId  `bson:"_id,omitempty"`
	RepoName string
	Owner string
	FullName string
	Description string
	StarsCount int
	ForksCount int
	UpdatedAt time.Time
	LastUpdatedBy string
	ReadMe string
	Tags []string
	Categories []string
}

