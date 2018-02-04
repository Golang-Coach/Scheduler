package interfaces

import "github.com/globalsign/mgo"

type ICollection interface {
	Insert(...interface{}) error
	Update(selector interface{}, update interface{}) error
	Find(query interface{}) *mgo.Query
}
