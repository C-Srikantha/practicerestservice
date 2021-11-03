package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

type Movie struct {
	Id        uint16
	MovieName string
	MovieLang string
	MovieType string
}
type Cast struct {
	Id           uint16
	ProducerName string
	DirectorName string
	ActorName    string
}

//this func displays the error message
func error(err interface{}) {
	if err != nil {
		fmt.Println(err)
	}
}

//displays all the movie details from db
func Getdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	var det []Movie
	err := db.Model(&det).Select() //Selects all details from table movie and stores it to Movie struct
	error(err)
	b, err := json.Marshal(det) //decoding struct to json
	error(err)
	str := string(b) //converting  to string
	fmt.Fprintf(w, str)
}

//displays details of particular id from db
func Getadetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r) //returns id of the current request
	key := vars["id"]
	intkey, _ := strconv.Atoi(key) //converts string to int
	det := new(Movie)
	err := db.Model(det).Where("id=?", intkey).Select() //Selects details of particular id and stores it to var det of type struct
	error(err)
	b, err := json.Marshal(det) //decoding struct to json
	error(err)
	str := string(b) //converting  to string
	fmt.Fprintf(w, str)
}

//inserts details to the table Movie
func Postdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	newpost, _ := ioutil.ReadAll(r.Body) //reads body from request and returns byte val
	var newdet Movie
	json.Unmarshal(newpost, &newdet)     //encoding json to struct
	_, err := db.Model(&newdet).Insert() //inserts newdetails to db
	error(err)
}

//Deletes details of particular request in db
func Deletedetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r) //returns id of the current request
	val := vars["id"]
	intval, _ := strconv.Atoi(val) //converts string to int
	var det Movie
	_, err := db.Model(&det).Where("id=?", intval).Delete() //Deletes details of particular id  from database
	error(err)
}

//it updates column  of the db movie table
func Updatedetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val)         //converts string to int
	updateval, _ := ioutil.ReadAll(r.Body) //reads the body of request and returns value in byte
	var updatedetail Movie
	json.Unmarshal(updateval, &updatedetail)
	_, err := db.Model(&updatedetail).Where("id=?", intval).Update() //updates the coulumn values of particular row
	error(err)
}

//selects all the details of cast from db
func Getcastdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	var det []Cast
	err := db.Model(&det).Select() //Selects all details from table Cast and stores it to Cast struct
	error(err)
	b, err := json.Marshal(det) //decoding struct to json
	error(err)
	str := string(b) //converting  to string
	fmt.Fprintf(w, str)
}

//selects the cast names of particular movie
func Getacastdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r) //returns id of the current request
	key := vars["id"]
	intkey, _ := strconv.Atoi(key) //converts string to int
	det := new(Cast)
	err := db.Model(det).Where("id=?", intkey).Select() //Selects details of particular id and stores it to var det of type struct
	error(err)
	b, err := json.Marshal(det) //decoding struct to json
	error(err)
	str := string(b) //converting  to string
	fmt.Fprintf(w, str)
}

//Inserts the cast details in db
func Postcastdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	newpost, _ := ioutil.ReadAll(r.Body) //reads body from request and returns byte val
	var newdet Cast
	json.Unmarshal(newpost, &newdet)     //encoding json to struct
	_, err := db.Model(&newdet).Insert() //inserts newdetails to db
	error(err)
}

//it updates column  of the db cast table
func Updatecastdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val) //converts string to int
	updateval, _ := ioutil.ReadAll(r.Body)
	var updatedetail Cast
	json.Unmarshal(updateval, &updatedetail)
	_, err := db.Model(&updatedetail).Where("id=?", intval).Update() //updates the coulumn values of particular row
	error(err)
}
