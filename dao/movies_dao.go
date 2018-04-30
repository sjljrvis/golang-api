package dao

import (
	. "github.com/sjljrvis/golang-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

// Connect  ...  Establish a connection to database ...
func (m *MoviesDAO) Connect() {
	log.Print(m.Server)
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//FindAll Find list of movies ...
func (m *MoviesDAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

//FindByID Find a movie by its id
func (m *MoviesDAO) FindByID(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update an existing movie
func (m *MoviesDAO) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
