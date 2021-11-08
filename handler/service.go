package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

type Actor struct {
	Id         uint
	ActorName  string
	ActorPhone int64
}
type Movie struct {
	Id        uint
	MovieName string
	MovieLang string
	MovieType string
	ActorID   uint
}
type Movierelease struct {
	Releaseyear int
	MovieId     uint
	Movie       *Movie
}

//this func displays the error message
func error(err interface{}) {
	if err != nil {
		fmt.Println(err)
	}
}

//displays all the actor details from db
func Getdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	var det []Actor
	err := db.Model(&det).Select() //Selects all details from table movie and stores it to Actor struct
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
	det := new(Actor)
	err := db.Model(det).Where("id=?", intkey).Select() //Selects details of particular id and stores it to var det of type struct
	error(err)
	b, err := json.Marshal(det) //decoding struct to json
	error(err)
	str := string(b) //converting  to string
	fmt.Fprintf(w, str)
}

//inserts details to the table Actor
func Postdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	newpost, _ := ioutil.ReadAll(r.Body) //reads body from request and returns byte val
	var newdet Actor
	json.Unmarshal(newpost, &newdet)     //encoding json to struct
	_, err := db.Model(&newdet).Insert() //inserts newdetails to db
	error(err)
}

//Deletes details of particular request in db
func Deletedetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r) //returns id of the current request
	val := vars["id"]
	intval, _ := strconv.Atoi(val) //converts string to int
	var det Actor
	_, err := db.Model(&det).Where("id=?", intval).Delete() //Deletes details of particular id  from database
	error(err)
}

//it updates column  of the db actor table
func Updatedetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val)         //converts string to int
	updateval, _ := ioutil.ReadAll(r.Body) //reads the body of request and returns value in byte
	var updatedetail Actor
	json.Unmarshal(updateval, &updatedetail)
	_, err := db.Model(&updatedetail).Where("id=?", intval).Update() //updates the coulumn values of particular row
	error(err)
}

//selects all the details of movie from db
func Getmoviedetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	var det []Movie
	err := db.Model(&det).Select() //Selects all details from table Movie and stores it to Movie struct
	error(err)
	b, err := json.Marshal(det) //decoding struct to json
	error(err)
	str := string(b) //converting  to string
	fmt.Fprintf(w, str)
}

//selects the movie names of particular movie
func Getamoviedetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r) //returns id of the current request
	key := vars["id"]
	intkey, _ := strconv.Atoi(key) //converts string to int
	var det []Movie
	err := db.Model(&det).Where("actor_id=?", intkey).Select() //Selects details of particular id and stores it to var det of type struct
	error(err)
	b, err := json.Marshal(det) //decoding struct to json
	error(err)
	str := string(b) //converting  to string
	fmt.Fprintf(w, str)
}

//Inserts the movie details in db
func Postcastdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	newpost, _ := ioutil.ReadAll(r.Body) //reads body from request and returns byte val
	var newdet Movie
	json.Unmarshal(newpost, &newdet)     //encoding json to struct
	_, err := db.Model(&newdet).Insert() //inserts newdetails to db
	error(err)
}

//it updates column  of the db movie table
func Updatecastdetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val) //converts string to int
	updateval, _ := ioutil.ReadAll(r.Body)
	var updatedetail Movie
	json.Unmarshal(updateval, &updatedetail)
	_, err := db.Model(&updatedetail).Where("id=?", intval).Update() //updates the coulumn values of particular row
	error(err)
}

/*func Getapidetails(w http.ResponseWriter, r *http.Request, db *pg.DB) {
	if db == nil {
		fmt.Fprint(w, "false")

	} else {
		fmt.Fprint(w, "true")
	}
}*/
func Readfile(db *pg.DB) {
	var detail []Actor
	var detail1 []Movie
	var detail2 []Movierelease
	//reading actor csv file
	readfile, err := os.Open("sample.csv") //open the file
	defer readfile.Close()
	if err != nil {
		fmt.Println("Failed to open file")
	}
	csvfile, err := csv.NewReader(readfile).ReadAll()
	temp := csvfile[1:]
	for _, details := range temp {
		id, _ := strconv.Atoi(details[0])       //converts string to int
		intphone, _ := strconv.Atoi(details[2]) //

		det := Actor{
			Id:         uint(id),
			ActorName:  details[1],
			ActorPhone: int64(intphone),
		}
		detail = append(detail, det)
	}
	//reading movie csv file
	readfile1, err := os.Open("movie.csv")
	defer readfile1.Close()
	if err != nil {
		fmt.Println("Failed to open file")
	}
	csvfile1, err := csv.NewReader(readfile1).ReadAll()
	temp1 := csvfile1[1:]
	for _, details := range temp1 {
		id, _ := strconv.Atoi(details[3])
		det := Movie{
			MovieName: details[0],
			MovieLang: details[1],
			MovieType: details[2],
			ActorID:   uint(id),
		}
		detail1 = append(detail1, det)
	}
	//reading release csv file
	readfile2, err := os.Open("release.csv")
	defer readfile2.Close()
	if err != nil {
		fmt.Println("Failed to open file")
	}
	csvfile2, err := csv.NewReader(readfile2).ReadAll()
	temp2 := csvfile2[1:]
	for _, details := range temp2 {
		year, _ := strconv.Atoi(details[0])
		id, _ := strconv.Atoi(details[1])
		det := Movierelease{
			Releaseyear: year,
			MovieId:     uint(id),
		}
		detail2 = append(detail2, det)
	}

	_, err = db.Model(&detail).Insert() //inserts into table Actor
	error(err)
	_, err = db.Model(&detail1).Insert() //inserts into table Movie
	error(err)
	_, err = db.Model(&detail2).Insert() //inserts into table Release
	error(err)

}
