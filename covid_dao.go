package dao

import (
	"log"

	. "myapp/models"
	mgo "gopkg.in/mgo.v2"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "covid_cases"
)

// Establish a connection to database
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie Covid) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}
